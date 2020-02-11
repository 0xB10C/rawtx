package rawtx

import (
	"testing"
)

func TestIsOPReturnOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsOPReturnOutput()
			expected := testTx.OutputTypes[index] == OutOPRETURN
			if result != expected {
				t.Errorf("Expected IsOPReturnOutput() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestGetOPReturnData(t *testing.T) {

	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			if out.GetType() == OutOPRETURN {
				is, opcode := out.GetOPReturnData()
				expected := testTx.OpReturnData[index]
				if is != expected.is ||
					len(opcode.PushedData) != expected.length ||
					opcode.PushedData[len(opcode.PushedData)-1] != expected.lastByte {
					t.Errorf("Expected GetOPReturnData() to be %+v, but got [%t, %d, %d] for testTx %+v", expected, is, len(opcode.PushedData), opcode.PushedData[len(opcode.PushedData)-1], testTx)
				}
			}
		}
	}
}

func TestIsP2PKHOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2PKHOutput()
			expected := testTx.OutputTypes[index] == OutP2PKH
			if result != expected {
				t.Errorf("Expected IsP2PKHOutput() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestIsP2WPKHV0Output(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2WPKHV0Output()
			expected := testTx.OutputTypes[index] == OutP2WPKH
			if result != expected {
				t.Errorf("Expected IsP2WPKHV0Output() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestIsP2SHOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2SHOutput()
			expected := testTx.OutputTypes[index] == OutP2SH
			if result != expected {
				t.Errorf("Expected IsP2SHOutput() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestIsP2WSHV0Output(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2WSHV0Output()
			expected := testTx.OutputTypes[index] == OutP2WSH
			if result != expected {
				t.Errorf("Expected IsP2WSHV0Output() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestIsP2MSOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			if out.GetType() == OutP2MS {
				is, m, n := out.IsP2MSOutput()
				expected := testTx.P2MSType[index]
				if is != expected.is || m != expected.m || n != expected.n {
					t.Errorf("Expected IsP2MSOutput==[%t, %d-of-%d] at index %d, but got [%t, %d-of-%d]", expected.is, expected.m, expected.n, index, is, m, n)
				}
			}
		}
	}
}

func TestIsP2PKOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2PKOutput()
			expected := testTx.OutputTypes[index] == OutP2PK
			if result != expected {
				t.Errorf("Expected IsP2PKOutput() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestOutputGetType(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.GetType()
			expected := testTx.OutputTypes[index]
			if result != expected {
				t.Errorf("Expected GetType() to be %+v at index %d, but got %+v for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestOutputTypeString(t *testing.T) {
	if OutP2PK.String() != "P2PK" {
		t.Error("Expected P2PK")
	}
	if OutP2PKH.String() != "P2PKH" {
		t.Error("Expected P2PKH")
	}
	if OutP2WPKH.String() != "P2WPKH" {
		t.Error("Expected P2WPKH")
	}
	if OutP2MS.String() != "P2MS" {
		t.Error("Expected P2MS")
	}
	if OutP2SH.String() != "P2SH" {
		t.Error("Expected P2SH")
	}
	if OutP2WSH.String() != "P2WSH" {
		t.Error("Expected P2WSH")
	}
	if OutOPRETURN.String() != "OPRETURN" {
		t.Error("Expected COINBASE")
	}
	if OutUNKNOWN.String() != "UNKNOWN" {
		t.Error("Expected UNKNOWN")
	}
}
