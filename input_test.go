package rawtx

import (
	"testing"
)

func TestInputSpendsNativeSegWit(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNativeSegWit()
			expected := testTx.InputTypes[index] == InP2WPKH || testTx.InputTypes[index] == InP2WSH || testTx.InputTypes[index] == InP2TRKP || testTx.InputTypes[index] == InP2TRSP
			if result != expected {
				t.Errorf("Expected SpendsNativeSegWit() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsNestedSegWit(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNestedSegWit()
			expected := testTx.InputTypes[index] == InP2SH_P2WPKH || testTx.InputTypes[index] == InP2SH_P2WSH
			if result != expected {
				t.Errorf("Expected SpendsNestedSegWit() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsNestedP2WPKH(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNestedP2WPKH()
			expected := testTx.InputTypes[index] == InP2SH_P2WPKH
			if result != expected {
				t.Errorf("Expected SpendsNestedP2WPKH() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsNestedP2WSH(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNestedP2WSH()
			expected := testTx.InputTypes[index] == InP2SH_P2WSH
			if result != expected {
				t.Errorf("Expected SpendsNestedP2WSH() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsP2PKH(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2PKH()
			expected := testTx.InputTypes[index] == InP2PKH
			if result != expected {
				t.Errorf("Expected SpendsP2PKH() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsP2PKHWithIsCompressed(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			if in.GetType() == InP2PKH {
				_, compressed := in.SpendsP2PKHWithIsCompressed()
				expected := testTx.CompressedPubKey[index]
				if expected != compressed {
					t.Errorf("Expected SpendsP2PKHWithIsCompressed() to be %t at index %d, but got %t for testTx: %+v", expected, index, compressed, testTx)
				}
			}
		}
	}
}

func TestSpendsP2PK(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2PK()
			expected := testTx.InputTypes[index] == InP2PK
			if result != expected {
				t.Errorf("Expected SpendsP2PK() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsP2SH(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2SH()
			expected := testTx.InputTypes[index] == InP2SH
			if result != expected {
				t.Errorf("Expected SpendsP2SH() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsP2WSH(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2WSH()
			expected := testTx.InputTypes[index] == InP2WSH
			if result != expected {
				t.Errorf("Expected SpendsP2WSH() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsP2WPKH(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2WPKH()
			expected := testTx.InputTypes[index] == InP2WPKH
			if result != expected {
				t.Errorf("Expected SpendsP2WPKH() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestSpendsP2MS(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2MS()
			expected := testTx.InputTypes[index] == InP2MS
			if result != expected {
				t.Errorf("Expected SpendsP2MS() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestIsLNUniliteralClosing(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.IsLNUniliteralClosing()
			expected := testTx.IsLNUniliteralClosing
			if result != expected {
				t.Errorf("Expected IsLNUniliteralClosing() to be %t at index %d, but got %t for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestInputGetType(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.GetType()
			expected := testTx.InputTypes[index]
			if result != expected {
				t.Errorf("Expected GetType() to be %+v at index %d, but got %+v for testTx: %+v", expected, index, result, testTx)
			}
		}
	}
}

func TestInputTypeString(t *testing.T) {
	if InP2PK.String() != "P2PK" {
		t.Error("Expected P2PK")
	}
	if InP2PKH.String() != "P2PKH" {
		t.Error("Expected P2PKH")
	}
	if InP2SH_P2WPKH.String() != "P2SH_P2WPKH" {
		t.Error("Expected P2SH_P2WPKH")
	}
	if InP2WPKH.String() != "P2WPKH" {
		t.Error("Expected P2WPKH")
	}
	if InP2MS.String() != "P2MS" {
		t.Error("Expected P2MS")
	}
	if InP2SH.String() != "P2SH" {
		t.Error("Expected P2SH")
	}
	if InP2SH_P2WSH.String() != "P2SH_P2WSH" {
		t.Error("Expected P2SH_P2WSH")
	}
	if InP2WSH.String() != "P2WSH" {
		t.Error("Expected P2WSH")
	}
	if InCOINBASE.String() != "COINBASE" {
		t.Error("Expected COINBASE")
	}
	if InCOINBASE_WITNESS.String() != "COINBASE_WITNESS" {
		t.Error("Expected COINBASE_WITNESS")
	}
	if InUNKNOWN.String() != "UNKNOWN" {
		t.Error("Expected UNKNOWN")
	}
}
