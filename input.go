package rawtx

import (
	"bytes"
	"github.com/btcsuite/btcd/wire"
)

// InputType defines the input type
type InputType int

// Possible types a input can be
const (
	InP2PK InputType = iota + 1
	InP2PKH
	InP2SH_P2WPKH
	InP2WPKH
	InP2MS
	InP2SH
	InP2SH_P2WSH
	InP2WSH
	InCOINBASE
	InUNKNOWN
)

var inputTypeStringMap = map[InputType]string{
	InP2PK:        "P2PK",
	InP2PKH:       "P2PKH",
	InP2SH_P2WPKH: "P2SH_P2WPKH",
	InP2WPKH:      "P2WPKH",
	InP2MS:        "P2MS",
	InP2SH:        "P2SH",
	InP2SH_P2WSH:  "P2SH_P2WSH",
	InP2WSH:       "P2WSH",
	InCOINBASE:    "COINBASE",
	InUNKNOWN:     "UNKNOWN",
}

func (it InputType) String() string {
	return inputTypeStringMap[it]
}

// Outpoint represents a bitcoin transaction input's previous outpoint as a struct.
type Outpoint struct {
	PrevTxHash  [32]byte
	OutputIndex uint32
}

// FromWireOutpoint populates an Outpoint struct with values from a wire.OutPoint.
func (outpoint *Outpoint) FromWireOutpoint(wireOutpoint *wire.OutPoint) {
	copy(outpoint.PrevTxHash[:], wireOutpoint.Hash.CloneBytes())
	outpoint.OutputIndex = wireOutpoint.Index
}

// Input represents a bitcoin transaction input as a struct.
type Input struct {
	Outpoint  Outpoint
	ScriptSig BitcoinScript
	Sequence  uint32
	Witness   []BitcoinScript
	inputType InputType
}

// FromWireTxIn populates an Input struct with values from a wire.TxIn.
func (in *Input) FromWireTxIn(txIn *wire.TxIn) {
	in.Sequence = txIn.Sequence
	in.ScriptSig = txIn.SignatureScript
	in.Outpoint.FromWireOutpoint(&txIn.PreviousOutPoint)
	for _, witness := range txIn.Witness {
		in.Witness = append(in.Witness, witness)
	}
	in.inputType = in.GetType()
}

// GetType retruns the input type as a InputType
func (in *Input) GetType() InputType {
	if in.inputType != 0 {
		// return the cached input type
		return in.inputType
	}

	if in.IsCoinbase() {
		return InCOINBASE
	} else if in.SpendsP2PKH() {
		return InP2PKH
	} else if in.SpendsP2SH() {
		return InP2SH
	} else if in.SpendsNestedP2WPKH() {
		return InP2SH_P2WPKH
	} else if in.SpendsP2WPKH() {
		return InP2WPKH
	} else if in.SpendsNestedP2WSH() {
		return InP2SH_P2WSH
	} else if in.SpendsP2WSH() {
		return InP2WSH
	} else if in.SpendsP2MS() {
		return InP2MS
	} else if in.SpendsP2PK() {
		return InP2PK
	}
	return InUNKNOWN
}

// HasWitness returns a boolean indicating if an input has a witness
func (in *Input) HasWitness() bool {
	return len(in.Witness) > 0
}

// SpendsNativeSegWit checks if the input spend is a native SegWit input.
// A native SegWit input has a witness but an empty scriptSig.
func (in *Input) SpendsNativeSegWit() bool {
	pbs := in.ScriptSig.ParseWithPanic()
	if len(pbs) == 0 && in.HasWitness() {
		return in.SpendsP2WPKH() || in.SpendsP2WSH()
	}
	return false
}

// IsCoinbase checks if an input is a coinbase input by checking the previous-
// output-index to be equal to 0xffffffff and then checking the previous-tx-hash
// to be all zero.
func (in *Input) IsCoinbase() bool {
	// first do the inexpensive check if equal to 0xffffffff
	if in.Outpoint.OutputIndex == 0xffffffff {
		// only then check the more expensive equal for byte arrays
		if bytes.Equal(in.Outpoint.PrevTxHash[:], make([]byte, 32)) {
			return true
		}
	}
	return false
}

// SpendsP2WSH checks if an input spends a P2WSH input.
// Since all native SegWit inputs that aren't P2WPKH are probably
// P2WSH this function just returns the complement for all native SegWit
// transactions.
func (in *Input) SpendsP2WSH() bool {
	if in.HasWitness() && len(in.ScriptSig) == 0 {
		return !in.SpendsP2WPKH()
	}
	return false
}

// SpendsP2WPKH checks if an input spends a P2WPKH input.
// A P2WPKH input has a empty scriptSig, but contains exactly two items in the witness:
// [signature, pubkey]
func (in *Input) SpendsP2WPKH() bool {
	if in.HasWitness() && len(in.ScriptSig) == 0 {
		if len(in.Witness) == 2 {
			if len(in.Witness[0]) >= int(OpDATA1) && len(in.Witness[0]) <= int(OpDATA75) {
				signature := ParsedOpCode{OpCode: OpCode(len(in.Witness[0])), PushedData: in.Witness[0]}
				if signature.IsSignature() && len(in.Witness[1]) >= int(OpDATA1) && len(in.Witness[0]) <= int(OpDATA75) {
					pubKey := ParsedOpCode{OpCode: OpCode(len(in.Witness[1])), PushedData: in.Witness[1]}
					return pubKey.IsPubKey()
				}
			}
		}
	}
	return false
}

// SpendsP2MS checks if an input spends a P2MS input.
// A P2MS has no witness items and a scriptSig as follows:
// OP_0 <signature> [<signature>] [<signature>] (where [] means optional)
func (in *Input) SpendsP2MS() bool {
	if !in.HasWitness() && len(in.ScriptSig) > 0 {
		pbs := in.ScriptSig.ParseWithPanic()
		if len(pbs) >= 2 && pbs[0].OpCode == Op0 {
			switch len(pbs) - 1 {
			case 1: // has one signature for a 1-of-(1/2/3) multisig
				if pbs[1].IsSignature() {
					return true
				}
			case 2: // has two signatures for a 2-of-(2/3) multisig
				if pbs[1].IsSignature() && pbs[2].IsSignature() {
					return true
				}
			case 3: // has three signatures for a 3-of-3 multisig
				if pbs[1].IsSignature() && pbs[2].IsSignature() && pbs[3].IsSignature() {
					return true
				}
			default:
				return false
			}
		}
	}
	return false
}

// SpendsNestedSegWit checks if the input spend is a nested SegWit input.
// A nested SegWit input has an **not** empty scriptSig and a witness.
func (in *Input) SpendsNestedSegWit() bool {
	if in.SpendsNestedP2WPKH() || in.SpendsNestedP2WSH() {
		return true
	}
	return false
}

// SpendsNestedP2WPKH checks if the input spend is a nested P2WPKH input.
// A nested P2WPKH input has a witness and the scriptSig looks like
// OP_DATA_22(OP_0 OP_DATA_20(20 byte hash))
func (in *Input) SpendsNestedP2WPKH() bool {
	pbs := in.ScriptSig.ParseWithPanic()
	if in.HasWitness() && len(pbs) == 1 && pbs[0].OpCode == OpDATA22 {
		var inner BitcoinScript = pbs[0].PushedData
		innerPbs := inner.ParseWithPanic()
		if len(innerPbs) == 2 &&
			innerPbs[0].OpCode == Op0 &&
			innerPbs[1].OpCode == OpDATA20 {
			return true
		}
	}
	return false
}

// SpendsNestedP2WSH checks if the input spend is a nested P2WSH input.
// A nested P2WSH input has a witness and the scriptSig looks like
// OP_DATA_34(OP_0 OP_DATA_32(32 byte hash))
func (in *Input) SpendsNestedP2WSH() bool {
	pbs := in.ScriptSig.ParseWithPanic()
	if in.HasWitness() && len(pbs) == 1 && pbs[0].OpCode == OpDATA34 {
		var inner BitcoinScript = pbs[0].PushedData
		innerPbs := inner.ParseWithPanic()
		if len(innerPbs) == 2 &&
			innerPbs[0].OpCode == Op0 &&
			innerPbs[1].OpCode == OpDATA32 {
			return true
		}
	}
	return false
}

// SpendsP2PKH checks if the input spend is a P2PKH input.
// <signature> <pubkey>
func (in *Input) SpendsP2PKH() (spendsP2PKH bool) {
	pbs := in.ScriptSig.ParseWithPanic()
	if !in.HasWitness() && len(pbs) == 2 {
		return pbs[0].IsSignature() && pbs[1].IsPubKey()
	}
	return false
}

// SpendsP2PKHWithIsCompressed checks if the input spend is a P2PKH input.
// Additionally it returns a boolean indicating if the revealed pubkey is compressed.
// <signature> <pubkey>
func (in *Input) SpendsP2PKHWithIsCompressed() (spendsP2PKH bool, isCompressedPubKey bool) {
	pbs := in.ScriptSig.ParseWithPanic()
	if !in.HasWitness() && len(pbs) == 2 {
		isPubKey, isCompressedPubKey := pbs[1].IsPubKeyWithIsCompressed()
		return pbs[0].IsSignature() && isPubKey, isCompressedPubKey
	}
	return false, false
}

// SpendsP2PK checks if the input spend is a P2PK input.
// A P2PK input only contains the signature in the scriptSig.
// <signature>
func (in *Input) SpendsP2PK() bool {
	pbs := in.ScriptSig.ParseWithPanic()
	if !in.HasWitness() && len(pbs) == 1 {
		return pbs[0].IsSignature()
	}
	return false
}

// SpendsP2SH checks if the input spend is a P2SH input.
// A P2SH input has a redeemscript push at the end of the scriptSig,
// which is neither a signature or a pubkey.
func (in *Input) SpendsP2SH() (spendsP2SH bool) {
	pbs := in.ScriptSig.ParseWithPanic()
	if !in.HasWitness() && len(pbs) > 0 {
		redeemScript := pbs[len(pbs)-1]
		return !redeemScript.IsSignature() && !redeemScript.IsPubKey() // is not a signature and is not a pubkey
	}
	return false
}

// GetP2SHRedeemScript checks if the input spend is a P2SH input and then returns the redeemScript.
// The returned redeemScript is empty if the input does not spend a P2SH input.
func (in *Input) GetP2SHRedeemScript() (redeemScript BitcoinScript) {
	if in.SpendsP2SH() {
		pbs := in.ScriptSig.ParseWithPanic()
		if !in.HasWitness() && len(pbs) > 0 {
			return pbs[len(pbs)-1].PushedData
		}
	}
	return
}

// GetP2WSHRedeemScript checks if the input spend is a P2WSH input and then returns the redeemScript.
// The returned redeemScript is empty if the input does not spend a P2WSH input.
func (in *Input) GetP2WSHRedeemScript() (redeemScript BitcoinScript) {
	if in.SpendsP2WSH() {
		if in.HasWitness() && len(in.Witness) >= 1 {
			return in.Witness[len(in.Witness)-1]
		}
	}
	return
}

// GetNestedP2WSHRedeemScript checks if the input spend is a Nested-P2WSH input and then returns the redeemScript.
// The returned redeemScript is empty if the input does not spend a Nested-P2WSH input.
func (in *Input) GetNestedP2WSHRedeemScript() (redeemScript BitcoinScript) {
	if in.SpendsNestedP2WSH() {
		if in.HasWitness() && len(in.Witness) >= 1 {
			return in.Witness[len(in.Witness)-1]
		}
	}
	return
}

// SpendsMultisig checks if the input spend is a multisig input.
// Checked are P2MS, P2SH, P2SH-P2WSH and P2WSH inputs.
func (in *Input) SpendsMultisig() bool {
	if in.SpendsNestedP2WSH() {
		redeemScript := in.GetNestedP2WSHRedeemScript()
		if ok, _, _ := redeemScript.IsMultisigScript(); ok {
			return true
		}
	}

	if in.SpendsP2SH() {
		redeemScript := in.GetP2SHRedeemScript()
		if ok, _, _ := redeemScript.IsMultisigScript(); ok {
			return true
		}
	}

	if in.SpendsP2WSH() {
		redeemScript := in.GetP2WSHRedeemScript()
		if ok, _, _ := redeemScript.IsMultisigScript(); ok {
			return true
		}
	}

	if in.SpendsP2MS() {
		return true
	}

	return false
}

// IsLNUniliteralClosing checks if the input spend is a lightning network unilateral close.
//   OP_IF
// 	  pubKey
//   OP_ELSE
// 	  OP_DATA_X (i.e. CSV time) OP_CHECKSEQUENCEVERIFY OP_DROP pubKey
//   OP_ENDIF
//   OP_CHECKSIG
func (in *Input) IsLNUniliteralClosing() bool {
	redeemScript := in.GetP2WSHRedeemScript()
	pbs := redeemScript.ParseWithPanic()
	if len(pbs) == 9 {
		if pbs[0].OpCode == OpIF && pbs[1].IsPubKey() &&
			pbs[2].OpCode == OpELSE && pbs[4].OpCode == OpCHECKSEQUENCEVERIFY && pbs[5].OpCode == OpDROP && pbs[6].IsPubKey() &&
			pbs[7].OpCode == OpENDIF &&
			pbs[8].OpCode == OpCHECKSIG {
			return true
		}
	}
	return false
}
