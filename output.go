package rawtx

import (
	"github.com/btcsuite/btcd/wire"
)

// OutputType defines the input type
type OutputType int

// Possible types a output can be
const (
	OutP2PK OutputType = iota + 1
	OutP2PKH
	OutP2WPKH
	OutP2MS
	OutP2SH
	OutP2WSH
	OutOPRETURN
	OutUNKNOWN
)

var outputTypeStringMap = map[OutputType]string{
	OutP2PK:     "P2PK",
	OutP2PKH:    "P2PKH",
	OutP2WPKH:   "P2WPKH",
	OutP2MS:     "P2MS",
	OutP2SH:     "P2SH",
	OutP2WSH:    "P2WSH",
	OutOPRETURN: "OPRETURN",
	OutUNKNOWN:  "UNKNOWN",
}

// Output represents a bitcoin transaction output as a struct.
type Output struct {
	Value        int64
	ScriptPubKey BitcoinScript
	outputType   OutputType
}

// FromWireTxOut populates an Output struct with values from a wire.TxOut.
func (out *Output) FromWireTxOut(txOut *wire.TxOut) {
	out.ScriptPubKey = txOut.PkScript
	out.Value = txOut.Value
	out.outputType = out.GetType()
}

func (ot OutputType) String() string {
	return outputTypeStringMap[ot]
}

// GetType retruns the output type as a OutputType
func (out *Output) GetType() OutputType {
	if out.outputType != 0 {
		return out.outputType
	} else if out.IsP2PKHOutput() {
		return OutP2PKH
	} else if out.IsP2SHOutput() {
		return OutP2SH
	} else if out.IsP2WPKHV0Output() {
		return OutP2WPKH
	} else if out.IsP2WSHV0Output() {
		return OutP2WSH
	} else if out.IsOPReturnOutput() {
		return OutOPRETURN
	} else if is, _, _ := out.IsP2MSOutput(); is {
		return OutP2MS
	} else if out.IsP2PKOutput() {
		return OutP2PK
	}
	return OutUNKNOWN
}

// IsOPReturnOutput returns if an Output is an OP_RETURN output
// An OP_RETURN scriptPubKey looks like:
//  OP_RETURN <SomeDataPush> <OP_RETURN data>
func (out *Output) IsOPReturnOutput() (is bool) {
	if len(out.ScriptPubKey) > 0 {
		pbs := out.ScriptPubKey.ParseWithPanic()
		if pbs[0].OpCode == OpRETURN {
			return true
		}
	}
	return false
}

// GetOPReturnData returns a ParsedOpCode struct,
// which includes the Op Code and the data pushed by OP_RETURN.
// An OP_RETURN scriptPubKey looks like:
//  OP_RETURN <SomeDataPush> <OP_RETURN data>
func (out *Output) GetOPReturnData() (bool, ParsedOpCode) {
	if out.IsOPReturnOutput() {
		pbs := out.ScriptPubKey.ParseWithPanic()
		return true, pbs[1]
	}
	return false, ParsedOpCode{}
}

// IsP2PKHOutput returns a boolean indicating if a output is a P2PKH output
// A P2PKH scriptPubKey looks like:
//  OP_DUP OP_HASH160 OP_DATA_20(20 byte pubKeyHash) OP_EQUALVERIFY OP_CHECKSIG
//  OP_DUP OP_HASH160 OP_DATA_20(                  ) OP_EQUALVERIFY OP_CHECKSIG
func (out *Output) IsP2PKHOutput() bool {
	pbs := out.ScriptPubKey.ParseWithPanic()
	if len(pbs) == 5 {
		if pbs[0].OpCode == OpDUP && // OP_DUP
			pbs[1].OpCode == OpHASH160 && // OP_HASH160
			pbs[2].OpCode == OpDATA20 && // OP_DATA_20 // FIXME: could also be inefficent with OpPUSHDATA1 / 2 / 4  ?
			pbs[3].OpCode == OpEQUALVERIFY && // OP_EQUALVERIFY
			pbs[4].OpCode == OpCHECKSIG { // OP_CHECKSIG
			return true
		}
	}
	return false
}

// IsP2SHOutput returns a boolean indicating if a output is a P2SH output
// A P2SH scriptPubKey looks like:
//  OP_HASH160 OP_DATA_20(20 byte hash) OP_EQUAL
func (out *Output) IsP2SHOutput() bool {
	pbs := out.ScriptPubKey.ParseWithPanic()
	if len(pbs) == 3 {
		if pbs[0].OpCode == OpHASH160 &&
			pbs[1].OpCode == OpDATA20 && // FIXME: could also be inefficent with OpPUSHDATA1 / 2 / 4  ?
			pbs[2].OpCode == OpEQUAL {
			return true
		}
	}
	return false
}

// IsP2WPKHV0Output returns a boolean indicating if a output is a P2WPKH output with witness program 0
// A P2WPKH V0 output looks like:
//  OP_0 OP_DATA_20(20 byte hash) (where the leading OP_0 indicates witness program 0)
func (out *Output) IsP2WPKHV0Output() bool {
	pbs := out.ScriptPubKey.ParseWithPanic()
	if len(pbs) == 2 {
		if pbs[0].OpCode == Op0 && // witness program 0
			pbs[1].OpCode == OpDATA20 { // OP_DATA_20 // FIXME: could also be inefficent with OpPUSHDATA1 / 2 / 4  ?
			return true
		}
	}
	return false
}

// IsP2WSHV0Output returns a boolean indicating if a output is a P2WSH output with witness program 0
// A P2WSH V0 output looks like:
//  OP_0 (as witness program 0) OP_DATA_32(32 byte hash)
func (out *Output) IsP2WSHV0Output() bool {
	pbs := out.ScriptPubKey.ParseWithPanic()
	if len(pbs) == 2 {
		if pbs[0].OpCode == Op0 && // witness program 0
			pbs[1].OpCode == OpDATA32 { // OP_DATA_32 // FIXME: could also be inefficent with OpPUSHDATA1 / 2 / 4  ?
			return true
		}
	}
	return false
}

// IsP2MSOutput returns a boolean indicating if a output is a P2MS output
// A P2MS 1-of-2 scriptPubKey looks like:
//  OP_1 PubKey PubKey OP_2 OP_CHECKMULTISIG
func (out *Output) IsP2MSOutput() (is bool, m int, n int) {
	return out.ScriptPubKey.IsMultisigScript()
}

// IsP2PKOutput returns a boolean indicating if a output is a P2PK output
// A P2PK output looks like:
//  PubKey OP_CHECKSIG
func (out *Output) IsP2PKOutput() bool {
	pbs := out.ScriptPubKey.ParseWithPanic()
	if len(pbs) == 2 {
		if pbs[0].IsPubKey() && pbs[1].OpCode == OpCHECKSIG {
			return true
		}
	}
	return false
}
