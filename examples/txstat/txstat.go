package main

// This is a example to showcase the raw transaction stat functionality of rawtx
// It takes a rawtx as hex string and prints a JSON object containing stats about the transaction

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/0xb10c/rawtx"
	"github.com/btcsuite/btcd/wire"
)

func main() {
	raw := flag.String("raw", "", "raw transaction in hex")
	flag.Parse()

	if len(*raw) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	rawAsByteSlice, err := hex.DecodeString(*raw)
	if err != nil {
		fmt.Printf("Could not decode the raw transaction: %s\n", err)
		os.Exit(2)
	}

	// rawtx depends on wire.MsgTx for deserializing the raw transaction
	r := bytes.NewReader(rawAsByteSlice)
	wireTx := &wire.MsgTx{}
	err = wireTx.Deserialize(r)
	if err != nil {
		fmt.Printf("Could not deserialize the transaction: %s\n", err)
		os.Exit(3)
	}

	tx := rawtx.Tx{}
	tx.FromWireMsgTx(wireTx)

	txstats := tx.Stats()
	txstatJSON, err := json.MarshalIndent(txstats, "", "  ")
	if err != nil {
		fmt.Printf("Could not marshal txstats to JSON: %s\n", err)
		os.Exit(4)
	}

	fmt.Println(string(txstatJSON))
}
