package rawtx

import (
	"testing"
)

func TestGetSizeWithoutWitness(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetSizeWithoutWitness()
		expected := testTx.VSize

		if result != expected {
			t.Errorf("Expected GetSizeWithoutWitness() to be %d, but got %d for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestGetSizeWithWitness(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetSizeWithWitness()
		expected := testTx.Size

		if result != expected {
			t.Errorf("Expected GetSizeWithWitness() to be %d, but got %d for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestGetNumInputs(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetNumInputs()
		expected := len(testTx.InputTypes)

		if result != expected {
			t.Errorf("Expected GetNumInputs() to be %d, but got %d for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestGetNumOutputs(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetNumOutputs()
		expected := len(testTx.OutputTypes)

		if result != expected {
			t.Errorf("Expected OutputTypes() to be %d, but got %d for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestGetLocktime(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetLocktime()
		expected := testTx.Locktime

		if result != expected {
			t.Errorf("Expected GetLocktime() to be %d, but got %d for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestGetOutputSum(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetOutputSum()
		expected := testTx.OutputSum

		if result != expected {
			t.Errorf("Expected GetOutputSum() to be %d, but got %d for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestIsCoinbase(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsCoinbase()
		expected := testTx.InputTypes[0] == InCOINBASE || testTx.InputTypes[0] == InCOINBASE_WITNESS
		if result != expected {
			t.Errorf("Expected IsCoinbase() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestIsSpendingSegWit(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingSegWit()
		expected := false
		for _, inputType := range testTx.InputTypes {
			if inputType == InP2SH_P2WPKH || inputType == InP2SH_P2WSH || inputType == InP2WPKH || inputType == InP2WSH || inputType == InCOINBASE_WITNESS {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected IsSpendingSegWit() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestIsExplicitlyRBFSignaling(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsExplicitlyRBFSignaling()
		expected := testTx.IsExplicitlySignalingRBF

		if result != expected {
			t.Errorf("Expected IsExplicitlyRBFSignaling() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestIsBIP69Compliant(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsBIP69Compliant()
		expected := testTx.IsBIP69Compliant

		if result != expected {
			t.Errorf("Expected IsBIP69Compliant() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestIsSpendingNativeSegWit(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingNativeSegWit()
		expected := false
		for _, inputType := range testTx.InputTypes {
			if inputType == InP2WPKH || inputType == InP2WSH {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected IsSpendingNativeSegWit() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestIsSpendingNestedSegWit(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingNestedSegWit()
		expected := false
		for _, inputType := range testTx.InputTypes {
			if inputType == InP2SH_P2WPKH || inputType == InP2SH_P2WSH {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected IsSpendingNestedSegWit() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasP2PKHOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2PKHOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutP2PKH {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasP2PKHOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasP2SHOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2SHOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutP2SH {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasP2SHOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasP2WPKHOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2WPKHOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutP2WPKH {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasP2WPKHOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasP2WSHOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2WSHOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutP2WSH {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasP2WSHOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasOPReturnOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasOPReturnOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutOPRETURN {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasOPReturnOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasP2MSOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2MSOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutP2MS {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasP2MSOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestSpendsMultisig(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingMultisig()
		expected := testTx.IsSpendingMultisig

		if result != expected {
			t.Errorf("Expected IsSpendingMultisig() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}

func TestHasP2PKOutput(t *testing.T) {
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2PKOutput()
		expected := false
		for _, outputType := range testTx.OutputTypes {
			if outputType == OutP2PK {
				expected = true
				break
			}
		}

		if result != expected {
			t.Errorf("Expected HasP2PKOutput() to be %t, but got %t for testTx: %+v", expected, result, testTx)
		}
	}
}
