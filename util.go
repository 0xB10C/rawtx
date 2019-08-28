package rawtx

import (
	"encoding/hex"
	"strings"

	"github.com/btcsuite/btcd/wire"
)

/* utility functions */

// StringToTx returns a wire.MsgTx for a raw transaction hex string
func StringToTx(rawTx string) (Tx, error) {
	hexDecodedTx, err := HexDecodeRawTxString(rawTx)
	if err == nil {
		tx, err := DeserializeRawTxBytes(hexDecodedTx)
		if err == nil {
			return tx, nil
		}
	}
	return Tx{}, err
}

// DeserializeRawTxBytes returns a wire.MsgTx for a hex decoded rawTx as byte slice.
// If the rawTx is can't be deserialized an error is returned.
func DeserializeRawTxBytes(rawTx []byte) (tx Tx, err error) {
	wireTx := wire.NewMsgTx(1)
	r := strings.NewReader(string(rawTx)) // FIXME: could this be made more efficenient?
	err = wireTx.Deserialize(r)
	if err != nil {
		return
	}

	tx.FromWireMsgTx(wireTx)
	return
}

// HexDecodeRawTxString hex decodes a rawTx string and returns it as byte slice.
func HexDecodeRawTxString(rawTx string) (hexDecodedTx []byte, err error) {
	hexDecodedTx, err = hex.DecodeString(rawTx)
	if err != nil {
		return
	}
	return
}
