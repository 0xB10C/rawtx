package rawtx

import (
	"testing"
)

func TestIsOPReturnOutput(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PKH, P2PKH
	txm[tx2] = []bool{false, false} // P2PKH, P2PKH
	txm[tx3] = []bool{false}        // P2PKH
	txm[tx4] = []bool{false, false} // P2PKH, P2PKH
	txm[tx5] = []bool{false, false} // P2PKH, P2PKH
	txm[tx6] = []bool{false}        // nonstandard
	txm[tx7] = []bool{false}        // P2SH

	txm[opReturnTx1] = []bool{false, false, true} // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []bool{false, true}        // P2SH, OP_RETURN
	txm[opReturnTx3] = []bool{false, true, false} // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []bool{false, true, false} // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsOPReturnOutput()
			if result != expexted[index] {
				t.Errorf("Expexted IsOPReturnOutput==%t at index %d, but got %t", expexted[index], index, result)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestGetOPReturnData(t *testing.T) {

	type response struct {
		Is       bool
		Length   int
		LastByte byte
	}

	notOPReturn := response{false, 0, 0}
	txm := make(map[string][]response)
	txm[tx1] = []response{notOPReturn, notOPReturn} // P2PKH, P2PKH
	txm[tx2] = []response{notOPReturn, notOPReturn} // P2PKH, P2PKH
	txm[tx3] = []response{notOPReturn}              // P2PKH
	txm[tx4] = []response{notOPReturn, notOPReturn} // P2PKH, P2PKH
	txm[tx5] = []response{notOPReturn, notOPReturn} // P2PKH, P2PKH
	txm[tx6] = []response{notOPReturn}              // nonstandard
	txm[tx7] = []response{notOPReturn}              // P2SH
	txm[tx8] = []response{notOPReturn, notOPReturn} // P2SH, P2WSH

	txm[opReturnTx1] = []response{notOPReturn, notOPReturn, response{true, 20, 0x20}} // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []response{notOPReturn, response{true, 32, 0x55}}              // P2SH, OP_RETURN
	txm[opReturnTx3] = []response{notOPReturn, response{true, 20, 0x20}, notOPReturn} // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []response{notOPReturn, response{true, 20, 0x50}, notOPReturn} // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			is, opcode := out.GetOPReturnData()
			if is != expexted[index].Is || len(opcode.PushedData) != expexted[index].Length || (len(opcode.PushedData) > 0 && opcode.PushedData[len(opcode.PushedData)-1] != expexted[index].LastByte) {
				t.Errorf("Expexted GetOPReturnData==%v at index %d, but got [%t, %d, %d]", expexted[index], index, is, len(opcode.PushedData), opcode.PushedData[len(opcode.PushedData)-1])
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestIsP2PKHOutput(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{true, true}   // P2PKH, P2PKH
	txm[tx2] = []bool{true, true}   // P2PKH, P2PKH
	txm[tx3] = []bool{true}         // P2PKH
	txm[tx4] = []bool{true, true}   // P2PKH, P2PKH
	txm[tx5] = []bool{true, true}   // P2PKH, P2PKH
	txm[tx6] = []bool{false}        // nonstandard
	txm[tx7] = []bool{false}        // P2SH
	txm[tx8] = []bool{false, false} // P2SH, P2WSH

	txm[opReturnTx1] = []bool{true, true, false}  // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []bool{false, false}       // P2SH, OP_RETURN
	txm[opReturnTx3] = []bool{false, false, true} // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []bool{true, false, true}  // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2PKHOutput()
			if result != expexted[index] {
				t.Errorf("Expexted IsP2PKHOutput==%t at index %d, but got %t", expexted[index], index, result)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestIsP2WPKHV0Output(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PKH, P2PKH
	txm[tx2] = []bool{false, false} // P2PKH, P2PKH
	txm[tx3] = []bool{false}        // P2PKH
	txm[tx4] = []bool{false, false} // P2PKH, P2PKH
	txm[tx5] = []bool{false, false} // P2PKH, P2PKH
	txm[tx6] = []bool{false}        // nonstandard
	txm[tx7] = []bool{false}        // P2SH
	txm[tx8] = []bool{false, false} // P2SH, P2WSH
	txm[tx9] = []bool{true, false}  // P2WPKH, P2PKH

	txm[opReturnTx1] = []bool{false, false, false} // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []bool{false, false}        // P2SH, OP_RETURN
	txm[opReturnTx3] = []bool{false, false, false} // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []bool{false, false, false} // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2WPKHV0Output()
			if result != expexted[index] {
				t.Errorf("Expexted IsP2WPKHOutput==%t at index %d, but got %t", expexted[index], index, result)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestIsP2SHOutput(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PKH, P2PKH
	txm[tx2] = []bool{false, false} // P2PKH, P2PKH
	txm[tx3] = []bool{false}        // P2PKH
	txm[tx4] = []bool{false, false} // P2PKH, P2PKH
	txm[tx5] = []bool{false, false} // P2PKH, P2PKH
	txm[tx6] = []bool{false}        // nonstandard
	txm[tx7] = []bool{true}         // P2SH
	txm[tx8] = []bool{true, false}  // P2SH, P2WSH

	txm[opReturnTx1] = []bool{false, false, false} // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []bool{true, false}         // P2SH, OP_RETURN
	txm[opReturnTx3] = []bool{true, false, false}  // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []bool{false, false, false} // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2SHOutput()
			if result != expexted[index] {
				t.Errorf("Expexted IsP2SHOutput==%t at index %d, but got %t", expexted[index], index, result)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestIsP2WSHV0Output(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PKH, P2PKH
	txm[tx2] = []bool{false, false} // P2PKH, P2PKH
	txm[tx3] = []bool{false}        // P2PKH
	txm[tx4] = []bool{false, false} // P2PKH, P2PKH
	txm[tx5] = []bool{false, false} // P2PKH, P2PKH
	txm[tx6] = []bool{false}        // nonstandard
	txm[tx7] = []bool{false}        // P2SH
	txm[tx8] = []bool{false, true}  // P2SH, P2WSH

	txm[opReturnTx1] = []bool{false, false, false} // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []bool{false, false}        // P2SH, OP_RETURN
	txm[opReturnTx3] = []bool{false, false, false} // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []bool{false, false, false} // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2WSHV0Output()
			if result != expexted[index] {
				t.Errorf("Expexted IsP2WSHOutput==%t at index %d, but got %t", expexted[index], index, result)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestIsP2MSOutput(t *testing.T) {
	type result struct {
		is bool
		m  int
		n  int
	}

	noP2MS := result{false, 0, 0}

	txm := make(map[string][]result)
	txm[tx1] = []result{noP2MS, noP2MS}                      // P2PKH, P2PKH
	txm[tx2] = []result{noP2MS, noP2MS}                      // P2PKH, P2PKH
	txm[tx3] = []result{noP2MS}                              // P2PKH
	txm[tx4] = []result{noP2MS, noP2MS}                      // P2PKH, P2PKH
	txm[tx5] = []result{noP2MS, noP2MS}                      // P2PKH, P2PKH
	txm[tx6] = []result{noP2MS}                              // nonstandard
	txm[tx7] = []result{noP2MS}                              // P2SH
	txm[tx8] = []result{noP2MS, noP2MS}                      // P2SH, P2WSH
	txm[tx8] = []result{noP2MS, noP2MS}                      // P2SH, P2WSH
	txm[tx13] = []result{result{true, 1, 1}, noP2MS, noP2MS} // P2MS 1-of-1, P2PKH, P2PKH
	txm[tx14] = []result{result{true, 2, 3}}                 // P2MS 2-of-3
	txm[tx15] = []result{result{true, 3, 3}, noP2MS, noP2MS} // P2MS 3-of-3, P2PKH, P2PKH

	txm[opReturnTx1] = []result{noP2MS, noP2MS, noP2MS} // P2PKH, P2PKH, OP_RETURN
	txm[opReturnTx2] = []result{noP2MS, noP2MS}         // P2SH, OP_RETURN
	txm[opReturnTx3] = []result{noP2MS, noP2MS, noP2MS} // P2SH, OP_RETURN, P2PKH
	txm[opReturnTx4] = []result{noP2MS, noP2MS, noP2MS} // P2PKH, OP_RETURN, P2PKH

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			is, m, n := out.IsP2MSOutput()
			if is != expexted[index].is || m != expexted[index].m || n != expexted[index].n {
				t.Errorf("Expexted IsP2MSOutput==[%t, %d-of-%d] at index %d, but got [%t, %d-of-%d]", expexted[index].is, expexted[index].m, expexted[index].n, index, is, m, n)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestIsP2PKOutput(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PKH, P2PKH
	txm[tx2] = []bool{false, false} // P2PKH, P2PKH
	txm[tx3] = []bool{false}        // P2PKH
	txm[tx4] = []bool{false, false} // P2PKH, P2PKH
	txm[tx5] = []bool{false, false} // P2PKH, P2PKH
	txm[tx6] = []bool{false}        // nonstandard
	txm[tx7] = []bool{false}        // P2SH
	txm[tx8] = []bool{false, false} // P2SH, P2WSH
	txm[tx23] = []bool{true, true}  // P2PK, P2PK

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.IsP2PKOutput()
			if result != expexted[index] {
				t.Errorf("Expexted IsP2PKOutput==%t at index %d, but got %t", expexted[index], index, result)
				ps := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input scriptPubKey: %s", ps.String())
			}
		}
	}
}

func TestOutputGetType(t *testing.T) {
	txm := make(map[string][]OutputType)
	txm[tx1] = []OutputType{OutP2PKH, OutP2PKH}
	txm[tx2] = []OutputType{OutP2PKH, OutP2PKH}
	txm[tx3] = []OutputType{OutP2PKH}
	txm[tx4] = []OutputType{OutP2PKH, OutP2PKH}
	txm[tx5] = []OutputType{OutP2PKH, OutP2PKH}
	txm[tx6] = []OutputType{OutUNKNOWN}
	txm[tx7] = []OutputType{OutP2SH}

	txm[tx8] = []OutputType{OutP2SH, OutP2WSH}
	txm[tx9] = []OutputType{OutP2WPKH, OutP2PKH}
	txm[tx10] = []OutputType{OutP2PKH, OutOPRETURN, OutOPRETURN, OutOPRETURN}
	txm[tx11] = []OutputType{OutP2SH, OutP2PKH}
	txm[tx12] = []OutputType{OutP2PKH}
	txm[tx13] = []OutputType{OutP2MS, OutP2PKH, OutP2PKH}
	txm[tx14] = []OutputType{OutP2MS}

	txm[tx15] = []OutputType{OutP2MS, OutP2PKH, OutP2PKH}
	txm[tx16] = []OutputType{OutUNKNOWN, OutP2PKH, OutP2PKH} // invalid P2MS because pubkey is wrong
	txm[tx17] = []OutputType{OutP2PKH}
	txm[tx18] = []OutputType{OutP2PKH}
	txm[tx21] = []OutputType{OutP2WPKH}
	txm[tx22] = []OutputType{OutP2PKH}
	txm[tx23] = []OutputType{OutP2PK, OutP2PK}

	for txString, expexted := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, out := range tx.Outputs {
			result := out.GetType()
			if result != expexted[index] {
				t.Errorf("Expexted GetType==%s at index %d, but got %s", expexted[index], index, result)
				pbs := out.ScriptPubKey.ParseWithPanic()
				t.Errorf("input script pub key: %s", pbs.String())
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
