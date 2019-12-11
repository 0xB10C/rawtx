package rawtx

import (
	"testing"
)

func TestParseEmpty(t *testing.T) {
	var empty BitcoinScript = []byte{}

	if len(empty.Parse()) != 0 {
		t.Errorf("The empty BitcoinScript should have a empty ParsedBitcoinScript.")
	}

}

func TestReadPushLength(t *testing.T) {
	// using byte arrays here, since slices can't be used in a Go maps
	// the byte arrays are 0xff padded to be all 11 byte long
	push10byte := [11]byte{10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	push75byte := [11]byte{75, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	pushdata1push10byte := [11]byte{byte(OpPUSHDATA1), 10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	pushdata1push75byte := [11]byte{byte(OpPUSHDATA1), 75, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	pushdata2push10byte := [11]byte{byte(OpPUSHDATA2), 10, 00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	pushdata2push75byte := [11]byte{byte(OpPUSHDATA2), 75, 00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	pushdata4push10byte := [11]byte{byte(OpPUSHDATA4), 10, 00, 00, 00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	pushdata4push75byte := [11]byte{byte(OpPUSHDATA4), 75, 00, 00, 00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	scriptm := make(map[[11]byte][2]int)
	scriptm[push10byte] = [2]int{10, 1}
	scriptm[push75byte] = [2]int{75, 1}
	scriptm[pushdata1push10byte] = [2]int{10, 2}
	scriptm[pushdata1push75byte] = [2]int{75, 2}
	scriptm[pushdata2push10byte] = [2]int{10, 3}
	scriptm[pushdata2push75byte] = [2]int{75, 3}
	scriptm[pushdata4push10byte] = [2]int{10, 5}
	scriptm[pushdata4push75byte] = [2]int{75, 5}

	for script, expected := range scriptm {
		dataPushLength, encodingLength := BitcoinScript(script[:]).getDataPushLength()
		if dataPushLength != expected[0] || encodingLength != expected[1] {
			t.Errorf("Expected getDataPushLength==%d, offset %d, but got %d, %d", expected[0], expected[1], dataPushLength, encodingLength)
			t.Errorf("tx: %v", script)
		}
	}

	// test zero script length case
	dataPushLength, encodingLength := BitcoinScript([]byte{}).getDataPushLength()
	if dataPushLength != 0 || encodingLength != 0 {
		t.Errorf("Expected getDataPushLength==%d, offset %d, but got %d, %d", 0, 0, dataPushLength, encodingLength)
	}

}

func TestGetDataPushLengthPanic(t *testing.T) {
	// test invalid length for PUSHDATA_1, _2 and _4
	failPushData1 := [1]byte{76}
	failPushData2 := [1]byte{77}
	failPushData4 := [1]byte{78}
	failPushDataArr := [][1]byte{failPushData1, failPushData2, failPushData4}

	for _, fail := range failPushDataArr {
		dataPushLength, encodingLength := BitcoinScript(fail[:]).getDataPushLength()
		if dataPushLength != 0 && encodingLength != 1 {
			t.Errorf("getDataPushLength should have returned a dataPushLength of 0 (did return %d) and a encodingLenght of 1 (did return %d)", dataPushLength, encodingLength)
		}
	}
}

func TestIsMultisigScript(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false}                              // P2PK, P2WPKH
	txm[tx2] = []bool{false}                                     // P2SH-P2WPKH
	txm[tx3] = []bool{false, false}                              // P2PK, P2WSH
	txm[tx4] = []bool{false, false}                              // P2WSH, P2WSH
	txm[tx5] = []bool{true}                                      // P2SH-P2WSH 6-of-6
	txm[tx6] = []bool{false}                                     // P2WSH
	txm[tx7] = []bool{false}                                     // P2PKH
	txm[tx8] = []bool{true}                                      // P2WSH
	txm[tx9] = []bool{false}                                     // P2WPKH
	txm[tx10] = []bool{false}                                    // P2PK
	txm[tx11] = []bool{true, true, true, true, true, true, true} // P2SH,P2SH,P2SH,P2SH,P2SH,P2SH,P2SH
	txm[tx12] = []bool{false}                                    // P2PKH
	txm[tx13] = []bool{false}                                    // P2PKH
	txm[tx14] = []bool{false}                                    // P2PKH
	txm[tx16] = []bool{false}                                    // P2MS
	txm[tx17] = []bool{false}                                    // P2MS
	txm[tx18] = []bool{false}                                    // P2MS
	txm[tx20] = []bool{true}
	txm[tx22] = []bool{true} //  P2WSH 2-of-3

	for txString, expected := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			var redeemScript BitcoinScript

			redeemScript = in.GetNestedP2WSHRedeemScript()

			if len(redeemScript) == 0 {
				redeemScript = in.GetP2WSHRedeemScript()
			}

			if len(redeemScript) == 0 {
				redeemScript = in.GetP2SHRedeemScript()
			}

			result, m, n := redeemScript.IsMultisigScript()
			if result != expected[index] {
				t.Errorf("expected IsMultisigScript={%t, %d, %d}, got %t", expected[index], m, n, result)
				t.Errorf("redeemScript %s", redeemScript.Parse())
			}
		}
	}

	redeemScript := BitcoinScript([]byte{byte(Op1), byte(Op0), byte(Op0), byte(Op0), byte(Op3), byte(OpCHECKMULTISIG)})
	result, m, n := redeemScript.IsMultisigScript()
	if result != false || m != 0 || n != 0 {
		t.Errorf("expected %s not to be a valid Multisig script", redeemScript.Parse())
	}

	redeemScript = BitcoinScript([]byte{byte(Op1), byte(OpCHECKMULTISIG)})
	result, m, n = redeemScript.IsMultisigScript()
	if result != false || m != 0 || n != 0 {
		t.Errorf("expected %s not to be a valid Multisig script", redeemScript.Parse())
	}
}

func TestGetSigHash(t *testing.T) {
	txm := make(map[string][]bool)
	txm[tx1] = []bool{false, false} // P2PK, P2WPKH
	txm[tx7] = []bool{false}        // P2PKH
	txm[tx10] = []bool{false}       // P2PK
	txm[tx12] = []bool{false}       // P2PKH

	for txString := range txm {
		tx, err := StringToTx(txString)
		if err != nil {
			t.Error(err.Error())
		}

		for _, in := range tx.Inputs {
			for _, p := range in.ScriptSig.Parse() {
				if p.GetSigHash() != 0x1 {
					if p.IsSignature() {
						t.Errorf("Expected GetSigHash=%#x, but got %#x", 0x1, p.GetSigHash())
					} else if p.GetSigHash() != 0x00 {
						t.Errorf("Expected GetSigHash=%#x because the input is no signature, but got %#x", 0x00, p.GetSigHash())
					}
				}
			}
		}
	}
}

func TestPocString(t *testing.T) {
	op0 := ParsedOpCode{OpCode: Op0}
	if op0.String() != "OP_0" {
		t.Errorf("Expected op0.String()=OP_0, got %s", op0.String())
	}

	opData1 := ParsedOpCode{OpCode: OpDATA1, PushedData: []byte{0xfa}}
	if opData1.String() != "OP_DATA_1(fa)" {
		t.Errorf("Expected opData1.String()=OP_DATA_1(fa), got %s", opData1.String())
	}

	opPushData1 := ParsedOpCode{OpCode: OpPUSHDATA1, PushedData: []byte{0xde, 0xad, 0xbe, 0xef, 0x53}}
	if opPushData1.String() != "OP_PUSHDATA1(5, deadbeef53)" {
		t.Errorf("Expected opPushData1.String()=OP_PUSHDATA1(5, deadbeef53), got %s", opPushData1.String())
	}
}

func TestPBSString(t *testing.T) {
	op0 := ParsedOpCode{OpCode: Op0}
	opData1 := ParsedOpCode{OpCode: OpDATA1, PushedData: []byte{0xfa}}
	opPushData1 := ParsedOpCode{OpCode: OpPUSHDATA1, PushedData: []byte{0xde, 0xad, 0xbe, 0xef, 0x53}}

	pbs := ParsedBitcoinScript{op0, opData1, opPushData1}

	if pbs.String() != "OP_0 OP_DATA_1(fa) OP_PUSHDATA1(5, deadbeef53)" {
		t.Errorf("Expected pbs.String()=OP_0 OP_DATA_1(fa) OP_PUSHDATA1(5, deadbeef53), got %s", pbs.String())
	}
}
