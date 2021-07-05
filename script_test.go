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
	testTxns := GetTestTransactions()
	for _, testTx := range testTxns {
		tx, err := StringToTx(testTx.RawTx)
		if err != nil {
			t.Error(err.Error())
		}

		for index, in := range tx.Inputs {
			var redeemScript BitcoinScript

			switch in.GetType() {
			case InP2SH_P2WSH:
				redeemScript = in.GetNestedP2WSHRedeemScript()
			case InP2WSH:
				redeemScript = in.GetP2WSHRedeemScript()
			case InP2SH:
				redeemScript = in.GetP2SHRedeemScript()
			}

			is, m, n := redeemScript.IsMultisigScript()
			expected := testTx.MultisigType[index]
			if is != expected.is || m != expected.m || n != expected.n {
				t.Errorf("Expected IsMultisigScript={%t, %d, %d}, but got {%t, %d, %d} for testTx: %+v", expected.is, expected.m, expected.n, is, m, n, testTx)
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

func TestStrictlyDEREncodedECDSASig(t *testing.T) {
	ECDSASigStrictlyDER := ParsedOpCode{OpCode: OpDATA71, PushedData: []byte{0x30, 0x44, 0x02, 0x20, 0x3c, 0x02, 0xbd, 0x6f, 0x63, 0xe4, 0x79, 0xc7, 0xc6, 0xd1, 0xdc, 0x7d, 0x94, 0x3b, 0x58, 0x4b, 0xa2, 0x03, 0xc6, 0xf1, 0x50, 0x37, 0x9c, 0x78, 0x9f, 0x84, 0xf8, 0xa5, 0xf3, 0x3c, 0x95, 0x12, 0x02, 0x20, 0x1c, 0xcf, 0x01, 0xbb, 0xeb, 0x2c, 0x5f, 0x68, 0xb1, 0x78, 0xab, 0x96, 0xa1, 0xa5, 0x64, 0xaf, 0x66, 0x09, 0xf7, 0x33, 0x87, 0xb9, 0x35, 0x5a, 0x62, 0x65, 0x54, 0x48, 0xb6, 0xa2, 0x7a, 0x43, 0x01}}

	if ECDSASigStrictlyDER.IsCompressedECDSAPubKey() {
		t.Errorf("A DER encoded signature should not be recognized as compressed pubkey")
	}

	if ECDSASigStrictlyDER.IsUncompressedECDSAPubKey() {
		t.Errorf("A DER encoded signature should not be recognized as uncompressed pubkey")
	}

	if !ECDSASigStrictlyDER.IsSignature() {
		t.Errorf("A DER encoded signature should be recognized as a signature")
	}

	if !ECDSASigStrictlyDER.IsECDSASignature(true) {
		t.Errorf("A DER encoded signature should be recognized as ECDSA Signature")
	}

	if !ECDSASigStrictlyDER.IsECDSASignatureInStrictDER() {
		t.Errorf("A DER encoded signature should be recognized as a strictly DER encoded ECDSA signature")
	}

	if ECDSASigStrictlyDER.GetSigHash() != 0x01 {
		t.Errorf("The SigHash of the signature should be 0x01")
	}
}

func TestGarbageSignature(t *testing.T) {
	tooBigOpCodeSig := ParsedOpCode{OpCode: OpDATA75, PushedData: []byte{0x10}}
	tooSmallOpCodeSig := ParsedOpCode{OpCode: OpDATA6, PushedData: []byte{0x10}}
	emptyPushDataSig := ParsedOpCode{OpCode: OpDATA33, PushedData: []byte{}}

	if tooBigOpCodeSig.IsCompressedECDSAPubKey() {
		t.Errorf("A garbage signature should not be recognized as compressed pubkey")
	}
	if tooBigOpCodeSig.IsUncompressedECDSAPubKey() {
		t.Errorf("A garbage signature should not be recognized as uncompressed pubkey")
	}
	if tooBigOpCodeSig.IsSignature() {
		t.Errorf("A garbage signature should be recognized as a signature")
	}
	if tooBigOpCodeSig.IsECDSASignature(false) {
		t.Errorf("A garbage signature should be recognized as ECDSA Signature")
	}
	if tooBigOpCodeSig.IsECDSASignatureInStrictDER() {
		t.Errorf("A garbage signature should be recognized as a strictly DER encoded ECDSA signature")
	}
	if tooBigOpCodeSig.GetSigHash() != 0x00 {
		t.Errorf("The SigHash of the garbage signature should be invalid (0x00)")
	}

	if tooSmallOpCodeSig.IsCompressedECDSAPubKey() {
		t.Errorf("A garbage signature should not be recognized as compressed pubkey")
	}
	if tooSmallOpCodeSig.IsUncompressedECDSAPubKey() {
		t.Errorf("A garbage signature should not be recognized as uncompressed pubkey")
	}
	if tooSmallOpCodeSig.IsSignature() {
		t.Errorf("A garbage signature should be recognized as a signature")
	}
	if tooSmallOpCodeSig.IsECDSASignature(false) {
		t.Errorf("A garbage signature should be recognized as ECDSA Signature")
	}
	if tooSmallOpCodeSig.IsECDSASignatureInStrictDER() {
		t.Errorf("A garbage signature should be recognized as a strictly DER encoded ECDSA signature")
	}
	if tooSmallOpCodeSig.GetSigHash() != 0x00 {
		t.Errorf("The SigHash of the garbage signature should be invalid (0x00)")
	}

	if emptyPushDataSig.IsCompressedECDSAPubKey() {
		t.Errorf("A garbage signature should not be recognized as compressed pubkey")
	}
	if emptyPushDataSig.IsUncompressedECDSAPubKey() {
		t.Errorf("A garbage signature should not be recognized as uncompressed pubkey")
	}
	if emptyPushDataSig.IsSignature() {
		t.Errorf("A garbage signature should be recognized as a signature")
	}
	if emptyPushDataSig.IsECDSASignature(false) {
		t.Errorf("A garbage signature should be recognized as ECDSA Signature")
	}
	if emptyPushDataSig.IsECDSASignatureInStrictDER() {
		t.Errorf("A garbage signature should be recognized as a strictly DER encoded ECDSA signature")
	}
	if emptyPushDataSig.GetSigHash() != 0x00 {
		t.Errorf("The SigHash of the garbage signature should be invalid (0x00)")
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
