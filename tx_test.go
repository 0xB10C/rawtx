package rawtx

import (
	"testing"
)

func TestGetSizeWithoutWitness(t *testing.T) {
	txm := make(map[string]int)
	txm[tx1] = 261
	txm[tx2] = 170
	txm[tx3] = 253
	txm[tx4] = 218
	txm[tx5] = 316
	txm[tx6] = 107
	txm[tx7] = 189

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetSizeWithoutWitness()
		if result != expexted {
			t.Errorf("Expexted GetSizeWithoutWitness==%d, but got %d", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestGetSizeWithWitness(t *testing.T) {
	txm := make(map[string]int)
	txm[tx1] = 343
	txm[tx2] = 251
	txm[tx3] = 418
	txm[tx4] = 389
	txm[tx5] = 800
	txm[tx6] = 245
	txm[tx7] = 189

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetSizeWithWitness()
		if result != expexted {
			t.Errorf("Expexted GetSizeWithWitness==%d, but got %d", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestGetNumInputs(t *testing.T) {
	txm := make(map[string]int)
	txm[tx1] = 2
	txm[tx2] = 1
	txm[tx3] = 2
	txm[tx4] = 2
	txm[tx5] = 1
	txm[tx6] = 1
	txm[tx7] = 1

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetNumInputs()
		if result != expexted {
			t.Errorf("Expexted GetNumInputs==%d, but got %d", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestGetNumOutputs(t *testing.T) {
	txm := make(map[string]int)
	txm[tx1] = 2
	txm[tx2] = 2
	txm[tx3] = 1
	txm[tx4] = 2
	txm[tx5] = 2
	txm[tx6] = 1
	txm[tx7] = 1

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetNumOutputs()
		if result != expexted {
			t.Errorf("Expexted GetNumOutputs==%d, but got %d", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestGetLocktime(t *testing.T) {
	txm := make(map[string]uint32)
	txm[tx1] = 17
	txm[tx2] = 1170
	txm[tx3] = 0
	txm[tx4] = 0
	txm[tx5] = 0
	txm[tx6] = 0
	txm[tx7] = 0

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetLocktime()
		if result != expexted {
			t.Errorf("Expexted GetLocktime==%d, but got %d", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestGetOutputSum(t *testing.T) {
	txm := make(map[string]int64)
	txm[tx1] = 335790000
	txm[tx2] = 999996600
	txm[tx3] = 5000000000
	txm[tx4] = 20000000
	txm[tx5] = 987000000
	txm[tx6] = 1
	txm[tx7] = 16516810

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.GetOutputSum()
		if result != expexted {
			t.Errorf("Expexted GetOutputSum==%d, but got %d", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestIsCoinbase(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = false
	txm[tx9] = false
	txm[tx10] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsCoinbase()
		if result != expexted {
			t.Errorf("Expexted IsCoinbase==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestIsSpendingSegWit(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = true
	txm[tx2] = true
	txm[tx3] = true
	txm[tx4] = true
	txm[tx5] = true
	txm[tx6] = true
	txm[tx7] = false
	txm[tx9] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingSegWit()
		if result != expexted {
			t.Errorf("Expexted GetOutputSum==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestIsExplicitlyRBFSignaling(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = true
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = true
	txm[tx9] = false
	txm[tx10] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsExplicitlyRBFSignaling()
		if result != expexted {
			t.Errorf("Expexted IsExplicitlyRBFSignaling==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestIsBIP69Compliant(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = true
	txm[tx3] = true
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = true
	txm[tx7] = true
	txm[tx8] = true
	txm[tx9] = false
	txm[tx10] = false

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsBIP69Compliant()
		if result != expexted {
			t.Errorf("Expexted IsBIP69Compliant==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestIsSpendingNativeSegWit(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = true
	txm[tx2] = false
	txm[tx3] = true
	txm[tx4] = true
	txm[tx5] = false
	txm[tx6] = true
	txm[tx7] = false
	txm[tx8] = true
	txm[tx9] = true
	txm[tx10] = false

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingNativeSegWit()
		if result != expexted {
			t.Errorf("Expexted SpendsNativeSegWit==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestIsSpendingNestedSegWit(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = true
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = true
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = false
	txm[tx9] = false

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingNestedSegWit()
		if result != expexted {
			t.Errorf("Expexted SpendsNestedSegWit==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestHasP2PKHOutput(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = true
	txm[tx2] = true
	txm[tx3] = true
	txm[tx4] = true
	txm[tx5] = true
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = false
	txm[tx9] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2PKHOutput()
		if result != expexted {
			t.Errorf("Expexted HasP2PKHOutput==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestHasP2SHOutput(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = true
	txm[tx8] = true
	txm[tx9] = false

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2SHOutput()
		if result != expexted {
			t.Errorf("Expexted HasP2SHOutput==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestHasP2WPKHOutput(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = false
	txm[tx9] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2WPKHOutput()
		if result != expexted {
			t.Errorf("Expexted HasP2WPKHOutput==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestHasP2WSHOutput(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = true
	txm[tx9] = false

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2WSHOutput()
		if result != expexted {
			t.Errorf("Expexted HasP2WSHOutput==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestHasOPReturnOutput(t *testing.T) {
	opReturnMap := make(map[string]bool)

	// have a OP_RETURN Output
	opReturnMap[opReturnTx1] = true
	opReturnMap[opReturnTx2] = true
	opReturnMap[opReturnTx3] = true
	opReturnMap[opReturnTx4] = true

	// no OP_RETURN Output
	opReturnMap[tx1] = false
	opReturnMap[tx2] = false
	opReturnMap[tx3] = false
	opReturnMap[tx4] = false
	opReturnMap[tx5] = false
	opReturnMap[tx6] = false
	opReturnMap[tx7] = false
	opReturnMap[tx8] = false
	opReturnMap[tx9] = false

	for txString, hasOpReturnOutput := range opReturnMap {

		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		testingHasOpRetrunOut := tx.HasOPReturnOutput()
		if testingHasOpRetrunOut != hasOpReturnOutput {
			t.Errorf("hasOpReturnOutput(txOuts) = %t; want %t", testingHasOpRetrunOut, hasOpReturnOutput)
		}
	}
}

func TestHasP2MSOutput(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = false
	txm[tx9] = false
	txm[tx10] = false
	txm[tx11] = false
	txm[tx12] = false
	txm[tx13] = true
	txm[tx14] = true
	txm[tx15] = true
	txm[tx16] = false
	txm[tx17] = false

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2MSOutput()
		if result != expexted {
			t.Errorf("Expexted HasP2MSOutput==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestSpendsMultisig(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = true
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = true
	txm[tx9] = false
	txm[tx10] = false
	txm[tx11] = true
	txm[tx12] = false
	txm[tx13] = false
	txm[tx14] = false
	txm[tx15] = false
	txm[tx16] = true
	txm[tx17] = true
	txm[tx18] = true
	txm[tx19] = true
	txm[tx20] = true
	txm[tx21] = false
	txm[tx22] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.IsSpendingMultisig()
		if result != expexted {
			t.Errorf("Expexted SpendsMultisig==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}

func TestHasP2PKOutput(t *testing.T) {
	txm := make(map[string]bool)
	txm[tx1] = false
	txm[tx2] = false
	txm[tx3] = false
	txm[tx4] = false
	txm[tx5] = false
	txm[tx6] = false
	txm[tx7] = false
	txm[tx8] = false
	txm[tx9] = false
	txm[tx10] = false
	txm[tx11] = false
	txm[tx12] = false
	txm[tx13] = false
	txm[tx14] = false
	txm[tx15] = false
	txm[tx16] = false
	txm[tx17] = false
	txm[tx18] = false
	txm[tx19] = false
	txm[tx20] = false
	txm[tx21] = false
	txm[tx22] = false
	txm[tx23] = true

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		result := tx.HasP2PKOutput()
		if result != expexted {
			t.Errorf("Expexted HasP2PKOutput==%t, but got %t", expexted, result)
			t.Errorf("tx: %s", txString)
		}
	}
}
