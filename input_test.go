package rawtx

import (
	"testing"
)

func TestInputSpendsNativeSegWit(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, true} // P2PK, P2WPKH
	txm[tx2] = []bool{false}       // P2SH-P2WPKH
	txm[tx3] = []bool{false, true} // P2PK, P2WSH
	txm[tx4] = []bool{true, true}  // P2WSH, P2WSH
	txm[tx5] = []bool{false}       // P2SH-P2WSH
	txm[tx6] = []bool{true}        // P2WSH
	txm[tx7] = []bool{false}       // P2PKH
	txm[tx8] = []bool{true}        // P2WSH
	txm[tx9] = []bool{true}        // P2WPKH

	txm[opReturnTx1] = []bool{false}        // P2PKH
	txm[opReturnTx2] = []bool{false}        // P2SH-P2PKH
	txm[opReturnTx3] = []bool{false, false} // P2PKH, P2PKH
	txm[opReturnTx4] = []bool{false, false} // P2PKH, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNativeSegWit()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsNativeSegWit==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsNestedSegWit(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PK, P2WPKH
	txm[tx2] = []bool{true}         // P2SH-P2WPKH
	txm[tx3] = []bool{false, false} // P2PK, P2WSH
	txm[tx4] = []bool{false, false} // P2WSH, P2WSH
	txm[tx5] = []bool{true}         // P2SH-P2WSH
	txm[tx6] = []bool{false}        // P2WSH
	txm[tx7] = []bool{false}        // P2PKH
	txm[tx8] = []bool{false}        // P2WSH
	txm[tx9] = []bool{false}        // P2WPKH

	txm[opReturnTx1] = []bool{false}        // P2PKH
	txm[opReturnTx2] = []bool{true}         // P2SH-P2PKH
	txm[opReturnTx3] = []bool{false, false} // P2PKH, P2PKH
	txm[opReturnTx4] = []bool{false, false} // P2PKH, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNestedSegWit()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsNestedSegWit==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsNestedP2WPKH(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PK, P2WPKH
	txm[tx2] = []bool{true}         // P2SH-P2WPKH
	txm[tx3] = []bool{false, false} // P2PK, P2WSH
	txm[tx4] = []bool{false, false} // P2WSH, P2WSH
	txm[tx5] = []bool{false}        // P2SH-P2WSH
	txm[tx6] = []bool{false}        // P2WSH
	txm[tx7] = []bool{false}        // P2PKH
	txm[tx8] = []bool{false}        // P2WSH

	txm[opReturnTx1] = []bool{false}        // P2PKH
	txm[opReturnTx2] = []bool{true}         // P2SH-P2WPKH
	txm[opReturnTx3] = []bool{false, false} // P2PKH, P2PKH
	txm[opReturnTx4] = []bool{false, false} // P2PKH, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNestedP2WPKH()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsNestedP2WPKH==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsNestedP2WSH(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PK, P2WPKH
	txm[tx2] = []bool{false}        // P2SH-P2WPKH
	txm[tx3] = []bool{false, false} // P2PK, P2WSH
	txm[tx4] = []bool{false, false} // P2WSH, P2WSH
	txm[tx5] = []bool{true}         // P2SH-P2WSH
	txm[tx6] = []bool{false}        // P2WSH
	txm[tx7] = []bool{false}        // P2PKH
	txm[tx8] = []bool{false}        // P2WSH
	txm[tx9] = []bool{false}        // P2WPKH

	txm[opReturnTx1] = []bool{false}        // P2PKH
	txm[opReturnTx2] = []bool{false}        // P2SH-P2WPKH
	txm[opReturnTx3] = []bool{false, false} // P2PKH, P2PKH
	txm[opReturnTx4] = []bool{false, false} // P2PKH, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsNestedP2WSH()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsNestedP2WSH==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsP2PKH(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PK, P2WPKH
	txm[tx2] = []bool{false}        // P2SH-P2WPKH
	txm[tx3] = []bool{false, false} // P2PK, P2WSH
	txm[tx4] = []bool{false, false} // P2WSH, P2WSH
	txm[tx5] = []bool{false}        // P2SH-P2WSH
	txm[tx6] = []bool{false}        // P2WSH
	txm[tx7] = []bool{true}         // P2PKH
	txm[tx8] = []bool{false}        // P2WSH
	txm[tx9] = []bool{false}        // P2WPKH
	txm[tx12] = []bool{true}        // P2PKH (compressed)

	txm[opReturnTx1] = []bool{true}       // P2PKH
	txm[opReturnTx2] = []bool{false}      // P2SH-P2WPKH
	txm[opReturnTx3] = []bool{true, true} // P2PKH, P2PKH
	txm[opReturnTx4] = []bool{true, true} // P2PKH, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2PKH()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsP2PKH==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsP2PKHWithIsCompressed(t *testing.T) {
	txm := make(map[string][][2]bool)
	txm[tx1] = [][2]bool{{false, false}, {false, false}}     // P2PK, P2WPKH
	txm[tx12] = [][2]bool{{true, false}}                     // P2PKH (uncompressed)
	txm[opReturnTx1] = [][2]bool{{true, true}}               // P2PKH (compressed)
	txm[opReturnTx3] = [][2]bool{{true, true}, {true, true}} // P2PKH (compressed), P2PKH (compressed)
	txm[opReturnTx4] = [][2]bool{{true, true}, {true, true}} // P2PKH (compressed), P2PKH (compressed)

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			spends, compressed := in.SpendsP2PKHWithIsCompressed()
			if spends != expexted[index][0] || compressed != expexted[index][1] {
				t.Errorf("Expexted SpendsP2PKHWithIsCompressed=={%v} at index %d, but got {%t,%t}", expexted[index], index, spends, compressed)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsP2PK(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{true, false}                                      // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                            // P2SH-P2WPKH
	txm[tx3] = []bool{true, false}                                      // P2PK, P2WSH
	txm[tx4] = []bool{false, false}                                     // P2WSH, P2WSH
	txm[tx5] = []bool{false}                                            // P2SH-P2WSH
	txm[tx6] = []bool{false}                                            // P2WSH
	txm[tx7] = []bool{false}                                            // P2PKH
	txm[tx8] = []bool{false}                                            // P2WSH
	txm[tx9] = []bool{false}                                            // P2WPKH
	txm[tx10] = []bool{false}                                           // P2PK
	txm[tx11] = []bool{false, false, false, false, false, false, false} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                           // P2PKH

	txm[opReturnTx1] = []bool{false}        // P2PKH
	txm[opReturnTx2] = []bool{false}        // P2SH-P2WPKH
	txm[opReturnTx3] = []bool{false, false} // P2PKH, P2PKH
	txm[opReturnTx4] = []bool{false, false} // P2PKH, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2PK()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsP2PK==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsP2SH(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false}                              // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                     // P2SH-P2WPKH
	txm[tx3] = []bool{false, false}                              // P2PK, P2WSH
	txm[tx4] = []bool{false, false}                              // P2WSH, P2WSH
	txm[tx5] = []bool{false}                                     // P2SH-P2WSH
	txm[tx6] = []bool{false}                                     // P2WSH
	txm[tx7] = []bool{false}                                     // P2PKH
	txm[tx8] = []bool{false}                                     // P2WSH
	txm[tx9] = []bool{false}                                     // P2WPKH
	txm[tx10] = []bool{false}                                    // P2PK
	txm[tx11] = []bool{true, true, true, true, true, true, true} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                    // P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2SH()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsP2SH==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsP2WSH(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false}                                     // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                            // P2SH-P2WPKH
	txm[tx3] = []bool{false, true}                                      // P2PK, P2WSH
	txm[tx4] = []bool{true, true}                                       // P2WSH, P2WSH
	txm[tx5] = []bool{false}                                            // P2SH-P2WSH
	txm[tx6] = []bool{true}                                             // P2WSH
	txm[tx7] = []bool{false}                                            // P2PKH
	txm[tx8] = []bool{true}                                             // P2WSH
	txm[tx9] = []bool{false}                                            // P2WPKH
	txm[tx10] = []bool{false}                                           // P2PK
	txm[tx11] = []bool{false, false, false, false, false, false, false} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                           // P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2WSH()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsP2WSH==%t at index %d, but got %t", expexted[index], index, result)
			}
		}
	}
}

func TestSpendsP2WPKH(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, true}                                      // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                            // P2SH-P2WPKH
	txm[tx3] = []bool{false, false}                                     // P2PK, P2WSH
	txm[tx4] = []bool{false, false}                                     // P2WSH, P2WSH
	txm[tx5] = []bool{false}                                            // P2SH-P2WSH
	txm[tx6] = []bool{false}                                            // P2WSH
	txm[tx7] = []bool{false}                                            // P2PKH
	txm[tx8] = []bool{false}                                            // P2WSH
	txm[tx9] = []bool{true}                                             // P2WPKH
	txm[tx10] = []bool{false}                                           // P2PK
	txm[tx11] = []bool{false, false, false, false, false, false, false} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                           // P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2WPKH()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsP2WPKH==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestSpendsP2MS(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false}                                     // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                            // P2SH-P2WPKH
	txm[tx3] = []bool{false, false}                                     // P2PK, P2WSH
	txm[tx4] = []bool{false, false}                                     // P2WSH, P2WSH
	txm[tx5] = []bool{false}                                            // P2SH-P2WSH
	txm[tx6] = []bool{false}                                            // P2WSH
	txm[tx7] = []bool{false}                                            // P2PKH
	txm[tx8] = []bool{false}                                            // P2WSH
	txm[tx9] = []bool{false}                                            // P2WPKH
	txm[tx10] = []bool{false}                                           // P2PK
	txm[tx11] = []bool{false, false, false, false, false, false, false} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                           // P2PKH
	txm[tx13] = []bool{false}                                           // P2PKH
	txm[tx14] = []bool{false}                                           // P2PKH
	txm[tx16] = []bool{true}                                            // P2MS
	txm[tx17] = []bool{true}                                            // P2MS
	txm[tx18] = []bool{true}                                            // P2MS

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.SpendsP2MS()
			if result != expexted[index] {
				t.Errorf("Expexted SpendsP2MS==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestIsLNUniliteralClosing(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false}                                     // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                            // P2SH-P2WPKH
	txm[tx3] = []bool{false, false}                                     // P2PK, P2WSH
	txm[tx4] = []bool{false, false}                                     // P2WSH, P2WSH
	txm[tx5] = []bool{false}                                            // P2SH-P2WSH
	txm[tx6] = []bool{false}                                            // P2WSH
	txm[tx7] = []bool{false}                                            // P2PKH
	txm[tx8] = []bool{false}                                            // P2WSH
	txm[tx9] = []bool{false}                                            // P2WPKH
	txm[tx10] = []bool{false}                                           // P2PK
	txm[tx11] = []bool{false, false, false, false, false, false, false} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                           // P2PKH
	txm[tx13] = []bool{false}                                           // P2PKH
	txm[tx14] = []bool{false}                                           // P2PKH
	txm[tx16] = []bool{false}                                           // P2MS
	txm[tx17] = []bool{false}                                           // P2MS
	txm[tx18] = []bool{false}                                           // P2MS
	txm[tx21] = []bool{true}                                            // P2WSH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.IsLNUniliteralClosing()
			if result != expexted[index] {
				t.Errorf("Expexted IsLNChannelClosing==%t at index %d, but got %t", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
			}
		}
	}
}

func TestInputGetType(t *testing.T) {
	txm := make(map[string][]InputType)
	txm[tx1] = []InputType{InP2PK, InP2WPKH}
	txm[tx2] = []InputType{InP2SH_P2WPKH}
	txm[tx3] = []InputType{InP2PK, InP2WSH}
	txm[tx4] = []InputType{InP2WSH, InP2WSH}
	txm[tx5] = []InputType{InP2SH_P2WSH}
	txm[tx6] = []InputType{InP2WSH}
	txm[tx7] = []InputType{InP2PKH}
	txm[tx8] = []InputType{InP2WSH}
	txm[tx9] = []InputType{InP2WPKH}
	txm[tx10] = []InputType{InUNKNOWN}
	txm[tx11] = []InputType{InP2SH, InP2SH, InP2SH, InP2SH, InP2SH, InP2SH, InP2SH}
	txm[tx12] = []InputType{InP2PKH}
	txm[tx13] = []InputType{InP2PKH}
	txm[tx14] = []InputType{InP2PKH}
	txm[tx16] = []InputType{InP2MS}
	txm[tx17] = []InputType{InP2MS}
	txm[tx18] = []InputType{InP2MS}
	txm[tx21] = []InputType{InP2WSH}
	txm[tx22] = []InputType{InP2WSH}
	txm[tx23] = []InputType{InP2PK}

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			result := in.GetType()
			if result != expexted[index] {
				t.Errorf("Expexted GetType==%s at index %d, but got %s", expexted[index], index, result)
				pbs := in.ScriptSig.ParseWithPanic()
				t.Errorf("input script sig: %s", pbs.String())
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
	if InP2SH_P2WPKH.String() != "P2SH-P2WPKH" {
		t.Error("Expected P2SH-P2WPKH")
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
	if InP2SH_P2WSH.String() != "P2SH-P2WSH" {
		t.Error("Expected P2SH-P2WSH")
	}
	if InP2WSH.String() != "P2WSH" {
		t.Error("Expected P2WSH")
	}
	if InCOINBASE.String() != "COINBASE" {
		t.Error("Expected COINBASE")
	}
	if InUNKNOWN.String() != "UNKNOWN" {
		t.Error("Expected UNKNOWN")
	}
}
