package rawtx

import (
	"math"
	"math/big"
)

// DERMarker represents the compound maker used in the DER encoding
const DERMarker = 0x30

// DERValueMarker represents the value (integer) maker used in the DER encoding
const DERValueMarker = 0x02

// ECDSASignature hold the raw byte fields of a deserialized ECDSA signature
// However, these fields might e.g. still be padded with zero values.
type ECDSASignature struct {
	r []byte
	s []byte
}

func getHalfCurveOrderSecp256k1() *big.Int {

	hco, ok := new(big.Int).SetString("7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0", 16)
	if !ok {
		panic("the hardcoded half curve order is invalid")
	}

	return hco
}

// DeserializeECDSASignature tries to deserialize a ECDSA signature. The caller can
// choose strict DER encoding is required. If the function returns ok=true then
// the ECDSASignature object is filled with the r and S values. The caller is expected
// to pass a signature with the SigHash flag removed.
// https://github.com/bitcoin/bips/blob/master/bip-0066.mediawiki#der-encoding-reference
func DeserializeECDSASignature(sigBytes []byte, requireStrictDER bool) (sig ECDSASignature, ok bool) {

	// check for a minimum length of 8 bytes
	if len(sigBytes) < 8 {
		return sig, false
	}

	// check for a maximum length of 72 bytes
	if len(sigBytes) > 73 {
		return sig, false
	}

	// check for DER maker
	if sigBytes[0] != DERMarker {
		return sig, false
	}

	// extract the length of the signature
	sigLength := int(sigBytes[1])

	if requireStrictDER {
		// check that the length covers the entire byte array
		if sigLength != len(sigBytes)-2 {
			return sig, false
		}
	}

	// check the DERValueMarker for the r element
	if sigBytes[2] != DERValueMarker {
		return sig, false
	}

	// extract the length of the r element
	rLength := int(sigBytes[3])

	// check that the <length of S> is still inside the signature bounds
	if 5+rLength >= len(sigBytes) {
		return sig, false
	}

	// extract the length of the S element
	sLength := int(sigBytes[5+rLength])

	if requireStrictDER {
		// check that the length of the signature matches the sum of the length
		// of the elements.
		if rLength+sLength+6 != len(sigBytes) {
			return sig, false
		}
	}

	// check that the length of the r element is longer than 0
	if rLength <= 0 {
		return sig, false
	}

	if requireStrictDER {
		// check that the r element is positive
		if sigBytes[4]&0x80 == 0x80 {
			return sig, false
		}
	}

	if requireStrictDER {
		// Null bytes at the start of R are not allowed, unless R would
		// otherwise be interpreted as a negative number.
		if (rLength > 1) && (sigBytes[4] == 0x00) && !(sigBytes[5]&0x80 == 0x80) {
			return sig, false
		}
	}

	// check the DERValueMarker for the S element
	if sigBytes[rLength+4] != DERValueMarker {
		return sig, false
	}

	// check that the length of the s element is longer than 0
	if sLength <= 0 {
		return sig, false
	}

	if requireStrictDER {
		// check that the S element is positive
		if sigBytes[rLength+6]&0x80 == 0x80 {
			return sig, false
		}
	}

	if requireStrictDER {
		// Null bytes at the start of S are not allowed, unless S would otherwise be
		// interpreted as a negative number.
		if (sLength > 1) && (sigBytes[rLength+6] == 0x00) && !(sigBytes[rLength+7]&0x80 == 0x80) {
			return sig, false
		}
	}

	r := sigBytes[4 : 4+rLength]
	s := sigBytes[4+rLength+1+1 : int(math.Min(float64(4+rLength+1+1+sLength), float64(len(sigBytes))))]
	sig = ECDSASignature{r: r, s: s}

	return sig, true
}

// HasLowS checks if a ECDSASignature has a low S value as defined in
// https://github.com/bitcoin/bips/blob/master/bip-0146.mediawiki#low_s
func (sig *ECDSASignature) HasLowS() bool {
	s := new(big.Int).SetBytes(sig.s)
	if s.Cmp(getHalfCurveOrderSecp256k1()) < 1 {
		return true
	}
	return false
}

// HasLowR checks if a ECDSASignature has a low r value as implemented in
// https://github.com/bitcoin/bitcoin/pull/13666
func (sig *ECDSASignature) HasLowR() bool {
	if len(sig.r) == 0 {
		return true
	}

	// r is normally 32 byte.
	// If its less than 32 byte then then r can't be high.
	if len(sig.r) < 32 {
		return true
	}

	// r might be longer than 32 byte since the DER field could include padding
	// if the byte at [pos] is > 0x80 then the r is high
	pos := len(sig.r) - 32
	if sig.r[pos] >= 0x80 {
		return false
	}
	return true
}
