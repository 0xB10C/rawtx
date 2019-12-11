// Package rawtx helps you answer questions about raw bitcoin transactions, their inputs, outputs and scripts.
// More information https://github.com/0xB10C/rawtx
package rawtx

import (
	"log"
	"math"

	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil/txsort"
)

// Tx represents a bitcoin transaction as a struct.
type Tx struct {
	Hash                  string
	Version               int32
	Inputs                []Input
	Outputs               []Output
	Locktime              uint32
	bip69sorted           bool
	serializeSizeStripped int
	serializeSize         int
}

// FromWireMsgTx populates a Tx struct with values from a wire.MsgTx.
func (tx *Tx) FromWireMsgTx(wireTx *wire.MsgTx) {
	tx.Version = wireTx.Version
	tx.Locktime = wireTx.LockTime
	tx.serializeSize = wireTx.SerializeSize()
	tx.serializeSizeStripped = wireTx.SerializeSizeStripped()
	tx.bip69sorted = txsort.IsSorted(wireTx)
	tx.Hash = wireTx.TxHash().String()

	for _, wireInput := range wireTx.TxIn {
		in := Input{}
		in.FromWireTxIn(wireInput)
		tx.Inputs = append(tx.Inputs, in)
	}

	for _, wireOutput := range wireTx.TxOut {
		out := Output{}
		out.FromWireTxOut(wireOutput)
		tx.Outputs = append(tx.Outputs, out)
	}
}

// GetNumInputs returns the number of inputs the transaction has
func (tx *Tx) GetNumInputs() int {
	return len(tx.Inputs)
}

// GetNumOutputs returns the number of outputs the transaction has
func (tx *Tx) GetNumOutputs() int {
	return len(tx.Outputs)
}

// GetOutputSum returns the sum of all output values of the transaction in satoshi
func (tx *Tx) GetOutputSum() (sumOutputValues int64) {
	for _, out := range tx.Outputs {
		sumOutputValues += out.Value
	}
	return
}

// GetLocktime returns the locktime of the transaction
func (tx *Tx) GetLocktime() uint32 {
	return tx.Locktime
}

// GetSizeWithoutWitness returns the transaction size **with** the witness stripped (vsize in vbyte)
func (tx *Tx) GetSizeWithoutWitness() int {
	vsizeFloat := (float64(tx.serializeSizeStripped*3) + float64(tx.serializeSize)) / 4
	return int(math.Ceil(vsizeFloat))
}

// GetSizeWithWitness returns the transaction size **without** the witness stripped (size in bytes)
func (tx *Tx) GetSizeWithWitness() int {
	return tx.serializeSize
}

// IsSpendingSegWit returns a boolean indicating if a transaction spends SegWit inputs
func (tx *Tx) IsSpendingSegWit() bool {
	for _, in := range tx.Inputs {
		if in.HasWitness() {
			return true
		}
	}
	return false
}

// IsSpendingNativeSegWit returns a boolean indicating if the transaction spends native SegWit
func (tx *Tx) IsSpendingNativeSegWit() bool {
	if tx.IsSpendingSegWit() {
		for _, in := range tx.Inputs {
			if in.SpendsNativeSegWit() {
				return true
			}
		}
	}
	return false
}

// IsSpendingNestedSegWit returns a boolean indicating if the transaction spends nested SegWit
func (tx *Tx) IsSpendingNestedSegWit() bool {
	if tx.IsSpendingSegWit() {
		for _, in := range tx.Inputs {
			if in.SpendsNestedSegWit() {
				return true
			}
		}
	}
	return false
}

// IsSpendingMultisig returns a boolean indicating if the transaction spends a multisig input
func (tx *Tx) IsSpendingMultisig() bool {
	for _, in := range tx.Inputs {
		if in.SpendsMultisig() {
			return true
		}
	}
	return false
}

// IsCoinbase returns a boolean indicating if a the transaction is a coinbase
// transaction.
func (tx *Tx) IsCoinbase() bool {
	if len(tx.Inputs) == 0 {
		log.Printf("Transaction %s has no input.", tx.Hash)
		return false
	}
	return tx.Inputs[0].IsCoinbase()
}

// IsExplicitlyRBFSignaling returns a boolean indicating if the transaction is explicitly signaling ReplaceByFee
// The transaction might still be implicitly able to be replaced by fee by e.g. a parent transaction signaling RBG
func (tx *Tx) IsExplicitlyRBFSignaling() bool {
	for _, in := range tx.Inputs {
		if in.Sequence < (uint32(0xffffffff) - 1) {
			return true
		}
	}
	return false
}

// IsBIP69Compliant returns a boolean indicating if the transaction is BIP 69 compliant
func (tx *Tx) IsBIP69Compliant() bool {
	return tx.bip69sorted
}

// HasP2PKHOutput returns a boolean indicating if the transaction has a P2PKH output
func (tx *Tx) HasP2PKHOutput() bool {
	for _, out := range tx.Outputs {
		if out.IsP2PKHOutput() {
			return true
		}
	}
	return false
}

// HasP2SHOutput returns a boolean indicating if the transaction has a P2SH output
func (tx *Tx) HasP2SHOutput() bool {
	for _, out := range tx.Outputs {
		if out.IsP2SHOutput() {
			return true
		}
	}
	return false
}

// HasP2WPKHOutput returns a boolean indicating if the transaction has a P2WPKH output
func (tx *Tx) HasP2WPKHOutput() bool {
	for _, out := range tx.Outputs {
		if out.IsP2WPKHV0Output() {
			return true
		}
	}
	return false
}

// HasP2WSHOutput returns a boolean indicating if the transaction has a P2WSH output
func (tx *Tx) HasP2WSHOutput() bool {
	for _, out := range tx.Outputs {
		if out.IsP2WSHV0Output() {
			return true
		}
	}
	return false
}

// HasP2MSOutput returns a boolean indicating if the transaction has a P2MS output
func (tx *Tx) HasP2MSOutput() bool {
	for _, out := range tx.Outputs {
		has, _, _ := out.IsP2MSOutput()
		if has {
			return true
		}
	}
	return false
}

// HasP2PKOutput returns a boolean indicating if the transaction has a P2PK output
func (tx *Tx) HasP2PKOutput() bool {
	for _, out := range tx.Outputs {

		if out.IsP2PKOutput() {
			return true
		}
	}
	return false
}

// HasOPReturnOutput returns a boolean indicating if the transaction has a OP_RETURN output
func (tx *Tx) HasOPReturnOutput() bool {
	for _, out := range tx.Outputs {
		if out.IsOPReturnOutput() {
			return true
		}
	}
	return false
}
