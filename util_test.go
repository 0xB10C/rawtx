package rawtx

import (
	"testing"
)

func TestHexDecodeRawTxString(t *testing.T) {
	_, err := HexDecodeRawTxString("abcdefghi")
	if err == nil {
		t.Errorf("The input %s should not be a valid hex string.\n", "abcdefghi")
	}
}

func TestDeserializeRawTxBytes(t *testing.T) {
	_, err := DeserializeRawTxBytes([]byte{0xff, 0x00, 0x52})
	if err == nil {
		t.Errorf("The input %s should not be a valid transaction.\n", "0xff, 0x00, 0x52")
	}
}

func TestStringToTx(t *testing.T) {
	_, err := StringToTx("abcdefghi")
	if err == nil {
		t.Errorf("The input %s should not be a valid hex string.\n", "abcdefghi")
	}
}
