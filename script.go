package rawtx

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

// BitcoinScript represents a bitcoin script as a byte slice
// A script can for example be the scriptPubKey or the scriptSig
type BitcoinScript []byte

// ParsedOpCode respresents a OP Code as a part of a parsed BitcoinScript.
// TODO: rename ito ParsedBitcoinScriptElement or something similar?
type ParsedOpCode struct {
	OpCode     OpCode
	PushedData []byte
}

func (poc *ParsedOpCode) String() string {
	// handle OP_DATA_X as OP_DATA_X(<data in hex>)
	if poc.OpCode >= OpDATA1 && poc.OpCode <= OpDATA75 {
		return OpCodeStringMap[poc.OpCode] + "(" + hex.EncodeToString(poc.PushedData) + ")"
	}

	// handle OP_PUSHDATAX as OP_PUSHDATAX(X bytes, <data in hex>)
	if poc.OpCode == OpPUSHDATA1 || poc.OpCode == OpPUSHDATA2 || poc.OpCode == OpPUSHDATA4 {
		return OpCodeStringMap[poc.OpCode] + "(" + strconv.Itoa(len(poc.PushedData)) + ", " + hex.EncodeToString(poc.PushedData) + ")"
	}

	return OpCodeStringMap[poc.OpCode]
}

// GetDataPushOpCodeForLength returns the optimal DataPush OpCode for the given data length
// OpINVALIDOPCODE is returned for zero, negative or bigger than 0xffffffff inputs.
func GetDataPushOpCodeForLength(length int) OpCode {
	if length < 1 {
		return OpINVALIDOPCODE
	}

	if length >= 1 && length <= 75 {
		return OpCode(length)
	}

	if length > 75 && length <= 0xff {
		return OpPUSHDATA1
	} else if length > 0xff && length <= 0xffff {
		return OpPUSHDATA2
	} else if length > 0xffff && length <= 0xffffffff {
		return OpPUSHDATA4
	}

	return OpINVALIDOPCODE
}

// ParsedBitcoinScript is a parsed BitcoinScript
type ParsedBitcoinScript []ParsedOpCode

func (pbs ParsedBitcoinScript) String() (s string) {
	for _, poc := range pbs {
		s += poc.String() + " "
	}
	return strings.TrimSuffix(s, " ")
}

// Parse parses the BitcoinScript and returns a ParsedBitcoinScript. It expects
// a script with correctly formatted data pushes, or it might return
// ParsedOpCodes with shorter than expected data for pushes-past-script-end.
func (s BitcoinScript) Parse() (parsed ParsedBitcoinScript) {
	if len(s) == 0 {
		return parsed
	}

	parsed = make([]ParsedOpCode, 0)
	for len(s) > 0 {
		p, remainder := s.parseNextOpCode()
		parsed = append(parsed, p)
		s = remainder
	}
	return
}

// parseNextOpCode parses the next OpCode in the BitcoinScript. It returns the
// parsed OpCode and the remaining BitcoinScript.
func (s BitcoinScript) parseNextOpCode() (front ParsedOpCode, remainder BitcoinScript) {
	opCode := OpCode(s[0])

	if opCode.IsDataPushOpCode() {
		dataPushLength, encodingLength := s.getDataPushLength()

		var opCodeLength int = dataPushLength + encodingLength

		// BitcoinScripts, for example in coinbase inputs, can push past the script
		// length.
		if opCodeLength > len(s) {
			opCodeLength = len(s)
		}

		return ParsedOpCode{OpCode: OpCode(opCode), PushedData: s[encodingLength:opCodeLength]}, s[opCodeLength:]
	}

	return ParsedOpCode{OpCode: OpCode(opCode)}, s[1:]
}

// getDataPushLength expects the next OpCode in the BitcoinScript to be an
// OpCode pushing data to the stack.
// The length of the pushed data (`dataPushLength`) and the number of bytes used
// to encode the data push `encodingLength` (including a byte for the OpCode)
// are returned.
//
// If the BitcoinScript is not long enough to encode a valid data push OpCode
// then both the dataPushLength and the encodingLength are 0.
func (s BitcoinScript) getDataPushLength() (dataPushLength int, encodingLength int) {
	if len(s) > 0 {
		opCode := OpCode(s[0])

		if opCode >= OpDATA1 && opCode <= OpDATA75 {
			dataPushLength = int(opCode)
			encodingLength = 1 // 1 byte OpCode
			return

		} else if opCode == OpPUSHDATA1 {
			if len(s) >= 2 {
				dataPushLength = int(s[1]) // takes one byte
				encodingLength = 1 + 1     // 1 byte OpCode + 1 byte to encode the data push length
				return
			}
		} else if opCode == OpPUSHDATA2 {
			if len(s) >= 3 {
				dataPushLength = int(binary.LittleEndian.Uint16(s[1:3])) // takes 2 bytes
				encodingLength = 1 + 2                                   // 1 byte OpCode + 2 byte to encode the data push length
				return
			}
		} else if opCode == OpPUSHDATA4 {
			if len(s) >= 5 {
				dataPushLength = int(binary.LittleEndian.Uint16(s[1:5])) // takes 4 bytes
				encodingLength = 1 + 4                                   // 1 byte OpCode + 4 byte to encode the data push length
				return
			}
		}
		// fallback if the remaining bytes in the BitcoinScript are less than the
		// data push OpCode would need
		dataPushLength = 0
		encodingLength = len(s)
		return
	}
	return 0, 0
}

// IsECDSASignature checks a ParsedOpCode if it could represent a ECDSA signature.
func (poc *ParsedOpCode) IsECDSASignature(requireStrictDER bool) bool {
	if len(poc.PushedData) < 1 {
		return false
	}

	if poc.OpCode <= OpDATA7 || poc.OpCode >= OpDATA75 {
		return false
	}

	sigBytes := poc.PushedData[:len(poc.PushedData)-1]
	_, ok := DeserializeECDSASignature(sigBytes, requireStrictDER)

	return ok
}

// IsECDSASignatureInStrictDER checks a ParsedOpCode if it represents a DER-encoded
// signature.
func (poc *ParsedOpCode) IsECDSASignatureInStrictDER() bool {
	if len(poc.PushedData) < 1 {
		return false
	}

	if poc.OpCode <= OpDATA7 || poc.OpCode >= OpDATA75 {
		return false
	}

	sigBytes := poc.PushedData[:len(poc.PushedData)-1]
	_, ok := DeserializeECDSASignature(sigBytes, true)

	return ok
}

// IsSignature checks a ParsedOpCode if it could represent a signature.
func (poc *ParsedOpCode) IsSignature() bool {
	if len(poc.PushedData) < 1 {
		return false
	}

	// All Sigs valid under IsECDSASignature are valid under IsECDSASignatureInStrictDER as well
	if poc.IsECDSASignature(false) {
		return true
	}

	// TODO: Add Schnorr signatures here

	return false
}

// GetSigHash returns the SIGHASH of a signature. If the passed
// ParsedOpCode does not push a Signature the returned SIGHASH is 0.
func (poc ParsedOpCode) GetSigHash() (sighash byte) {
	if poc.IsSignature() {
		return poc.PushedData[len(poc.PushedData)-1]
	}
	return 0x00
}

// IsPubKeyWithIsCompressed checks a two byte slices if they could represent a pubkey and if that pubkey is compressed.
func (poc ParsedOpCode) IsPubKeyWithIsCompressed() (isPubKey bool, isCompressed bool) {
	compressed := poc.IsCompressedPubKey()
	uncompressed := poc.IsUncompressedPubKey()
	return (compressed || uncompressed), compressed
}

// IsPubKey checks a two byte slices if they could represent a public key.
func (poc ParsedOpCode) IsPubKey() bool {
	if poc.IsCompressedPubKey() { // check compressed pubKeys first because they are more common
		return true
	} else if poc.IsUncompressedPubKey() {
		return true
	}
	return false
}

// IsCompressedPubKey checks a two byte slices if they could represent a
// compressed public key. <pubkey length> <|pubkey|>
func (poc ParsedOpCode) IsCompressedPubKey() bool {
	if poc.OpCode == OpDATA33 && len(poc.PushedData) == 33 {
		return (poc.PushedData[0] == 0x02 || poc.PushedData[0] == 0x03)
	}
	return false
}

// IsUncompressedPubKey checks a two byte slices if they could represent a
// uncompressed public key. <pubkey length> <|pubkey|>
func (poc ParsedOpCode) IsUncompressedPubKey() bool {
	if poc.OpCode == OpDATA65 && len(poc.PushedData) == 65 {
		return poc.PushedData[0] == 0x04
	}
	return false
}

// IsMultisigScript checks if a passed BitcoinScript is multisig.
// Supported are P2SH, P2SH-P2WSH and P2WSH redeem scripts.
func (s BitcoinScript) IsMultisigScript() (isMultisig bool, numRequiredSigs int, numPossiblePubKeys int) {
	if len(s) == 0 {
		return false, 0, 0
	}

	parsed := s.Parse()
	pLength := len(parsed)
	pLast := parsed[pLength-1] // last item in `parsed`

	// At least four OP codes are needed for a 1-of-1 multisig
	// <OP_1> <|pubkey|> <OP_1> <OP_CHECKMULTISIG>
	if pLength >= 4 && pLast.OpCode == OpCHECKMULTISIG {
		p2ndLast := parsed[pLength-2] // 2nd last item in `parsed`
		if p2ndLast.OpCode >= Op1 && p2ndLast.OpCode <= Op16 {
			numPossiblePubKeys = int(p2ndLast.OpCode - 80) // OP_1 starts with a of value 81 (0x51)

			// check that the next `numPossiblePubKeys` are pubKeys
			for i := 1; i < numPossiblePubKeys; i++ {
				if !parsed[pLength-2-i].IsPubKey() {
					return false, 0, 0
				}
			}

			pFirst := parsed[pLength-2-numPossiblePubKeys-1] // presumably the first item in `parsed`
			if pFirst.OpCode >= Op1 && pFirst.OpCode <= Op16 {
				numRequiredSigs = int(pFirst.OpCode - 80) // OP_1 starts with a of value 81 (0x51)
				return true, numRequiredSigs, numPossiblePubKeys
			}
		}
	}

	return false, 0, 0
}
