package rawtx

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// BitcoinScript represents a bitcoin script as a byte slice
// A script can for example be the scriptPubKey or the scriptSig
type BitcoinScript []byte

// ParsedOpCode respresents a OP Code as a part of a parsed BitcoinScript.
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

// ParsedBitcoinScript is a parsed BitcoinScript
type ParsedBitcoinScript []ParsedOpCode

func (pbs ParsedBitcoinScript) String() (s string) {
	for _, poc := range pbs {
		s += poc.String() + " "
	}
	return strings.TrimSuffix(s, " ")
}

// ParseWithPanic parses a BitcoinScript and returns contained opcodes
// and data pushes seperated as a ParsedBitcoinScript.
// Parse panics if a BitcoinScript can not be parsed!
func (s BitcoinScript) ParseWithPanic() (parsed ParsedBitcoinScript) {
	if len(s) == 0 {
		return parsed
	}

	parsed = make([]ParsedOpCode, 0)
	for len(s) > 0 {
		p, remainder, err := s.PopFront()
		if err != nil {
			panic(fmt.Errorf("could not parse the BitcoinScript: %s", err))
		}
		parsed = append(parsed, p)
		s = remainder
	}
	return
}

// PopFront returns the first OpCode and it's data push as a ParsedOpCode.
// Additionally the remaining BitcoinScript is returned.
func (s BitcoinScript) PopFront() (front ParsedOpCode, remainder BitcoinScript, err error) {
	opCode := OpCode(s[0])

	if (opCode >= OpDATA1 && opCode <= OpDATA75) || opCode == OpPUSHDATA1 || opCode == OpPUSHDATA2 || opCode == OpPUSHDATA4 {
		// the OP code indicates a data push
		length, offset, err := readPushLength(s)
		if err != nil {
			return front, remainder, fmt.Errorf("could not PopFront: %s", err)
		}

		if length+offset > len(s) {
			return front, remainder, fmt.Errorf("invalid data push in script after %s", OpCodeStringMap[opCode])
		}
		return ParsedOpCode{OpCode: OpCode(opCode), PushedData: s[offset : length+offset]}, s[length+offset:], nil
	}

	// the OP code does not push data
	return ParsedOpCode{OpCode: OpCode(opCode)}, s[1:], nil
}

// readPushLength returns the length of the pushed bytes at the beginning of the script
// and an offset at which the length encoding ends and the pushed data begins
func readPushLength(s []byte) (length int, offset int, err error) {
	if len(s) > 0 {
		opCode := OpCode(s[0])
		if opCode >= OpDATA1 && opCode <= OpDATA75 {
			return int(opCode), 1, nil
		} else if opCode == OpPUSHDATA1 {
			if len(s) >= 5 {
				return int(s[1]), 1 + 1, nil // 1 byte OPcode + 1 byte length
			}
			return 0, 0, fmt.Errorf("could not read data push length for %s", OpCodeStringMap[OpPUSHDATA1])
		} else if opCode == OpPUSHDATA2 {
			if len(s) >= 3 {
				return int(binary.LittleEndian.Uint16(s[1:4])), 1 + 2, nil // 1 byte OPcode + 2 byte length
			}
			return 0, 0, fmt.Errorf("could not read data push length for %s", OpCodeStringMap[OpPUSHDATA2])
		} else if opCode == OpPUSHDATA4 {
			if len(s) >= 5 {
				return int(binary.LittleEndian.Uint16(s[1:6])), 1 + 4, nil // 1 byte OPcode + 4 byte length
			}
			return 0, 0, fmt.Errorf("could not read data push length for %s", OpCodeStringMap[OpPUSHDATA4])
		}
	}
	return 0, 0, fmt.Errorf("s does not contain any data")
}

// IsSignature checks a two byte slices if they could represent a
// DER encoded signature. <sig length> <|signature|> <sig hash>
// While the signature looks like:
// <DER marker> <Sig length> <R marker> <R length> <|R value|> <S marker> <S length> <|S value|> <sig hash>
func (poc ParsedOpCode) IsSignature() bool {
	const derMarker = 0x30
	const derValueMarker = 0x02
	signature := poc.PushedData
	if poc.OpCode >= OpDATA1 && poc.OpCode <= OpDATA75 { // OPCode is between OP_DATA_1 and OP_DATA_75
		if len(signature) > 6 && // is at least seven bytes
			signature[0] == derMarker && // starts with DER signature marker
			signature[2] == derValueMarker { // and the third byte is a valid R value marker
			signatureLength := int(signature[1])
			rLength := int(signature[3])
			if signatureLength > rLength && len(signature) == signatureLength+1+1+1 { // + 1 for DER marker, +1 for signature length, +1 for the SigHash flag
				sMarker := signature[rLength+4]
				if sMarker == derValueMarker {
					return true
				}
			}
		}
	}
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
	return poc.OpCode == OpDATA33 && (poc.PushedData[0] == 0x02 || poc.PushedData[0] == 0x03)
}

// IsUncompressedPubKey checks a two byte slices if they could represent a
// uncompressed public key. <pubkey length> <|pubkey|>
func (poc ParsedOpCode) IsUncompressedPubKey() bool {
	return poc.OpCode == OpDATA65 && poc.PushedData[0] == 0x04
}

// IsMultisigScript checks if a passed BitcoinScript is multisig.
// Supported are P2SH, P2SH-P2WSH and P2WSH redeem scripts.
func (s BitcoinScript) IsMultisigScript() (isMultisig bool, numRequiredSigs int, numPossiblePubKeys int) {
	if len(s) == 0 {
		return false, 0, 0
	}

	parsed := s.ParseWithPanic()
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
