# rawtx

[![](https://godoc.org/github.com/0xB10C/rawtx?status.svg)](https://godoc.org/github.com/0xB10C/rawtx)

A Golang module that helps you answer questions about raw bitcoin transactions, their inputs, outputs and scripts.

### Transactions

Transactions can be deserialized [from a hex string][50] or [from][51] a [wire.MsgTx][52]. 
Based on that following questions can be answered.

- [x] [How many inputs?][60]
- [x] [How many outputs?][61]
- [x] [Whats the sum of all ouputs?][70]
- [x] [Spends SegWit?][62]
- [x] [Spends nested SegWit?][63]
- [x] [Spends native SegWit?][64]
- [x] [Spends multisig?][65]
- [x] [What is the locktime?][66]
- [x] [Is a coinbase transaction?][67]
- [x] [Is BIP69 compliant?][68] _Wait, what is [BIP-69](https://github.com/bitcoin/bips/blob/master/bip-0069.mediawiki)?_
- [x] [Signals explicit RBF?][69]
- [x] [Whats the size in bytes?][71]
- [x] [Whats the vsize in vbytes?][72]
- [x] [and more...][more]

[50]: https://www.godoc.org/github.com/0xb10c/rawtx/#StringToTx
[51]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.FromWireMsgTx
[52]: https://godoc.org/github.com/btcsuite/btcd/wire#MsgTx

[60]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.GetNumInputs
[61]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.GetNumOutputs
[62]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.IsSpendingSegWit
[63]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.IsSpendingNestedSegWit
[64]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.IsSpendingNativeSegWit
[65]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsMultisig
[66]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.GetLocktime
[67]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.IsCoinbase
[68]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.IsBIP69Compliant
[69]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.IsExplicitlyRBFSignaling
[70]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.GetOutputSum
[71]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.GetSizeWithWitness
[72]: https://www.godoc.org/github.com/0xb10c/rawtx/#Tx.GetSizeWithoutWitness

### Input and Output type 

- [x] [What type has this input?][24]
- [x] [What type has this output?][25]

Following is answered:

|             |  spends?  |  pays to? |
|-------------|:-------:|:-------:|
| P2PK        | [✓][01] | [✓][02] |
| P2PKH       | [✓][03] | [✓][04] |
| P2SH-P2WPKH | [✓][05] | can't tell <br> from  the rawtx |
| P2WPKH_V0   | [✓][07] | [✓][08] |
| P2MS        | [✓][09] | [✓][10] |
| P2SH        | [✓][11] | [✓][12] |
| P2SH-P2WSH  | [✓][13] | can't tell <br> from the rawtx |
| P2WSH_V0    | [✓][15] | [✓][16] |
| OP_RETURN   | can't spend an <br> `OP_RETURN` output | [✓][18] |

Additionally you might ask:

- What is the reedem script of this [P2SH][20], [P2SH-P2WSH][21] or [P2WSH_V0][22] input?
- What is the [OP_RETURN data pushed][23] by this output?
- and is this [input unilitteraly closing a lightning channel?][26]

[01]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsP2PK
[02]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsP2PKOutput
[03]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsP2PKH
[04]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsP2PKHOutput
[05]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsNestedP2WPKH
[07]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsP2PKH
[08]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsP2WPKHV0Output
[09]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsP2MS
[10]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsP2MSOutput
[11]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsP2SH
[12]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsP2SHOutput
[13]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsNestedP2WSH
[15]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.SpendsP2WSH
[16]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsP2WSHV0Output
[18]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.IsOPReturnOutput

[20]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.GetP2SHRedeemScript
[21]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.GetNestedP2WSHRedeemScript
[22]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.GetP2WSHRedeemScript
[23]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.GetOPReturnData
[24]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.GetType
[25]: https://www.godoc.org/github.com/0xb10c/rawtx/#Output.GetType
[26]: https://www.godoc.org/github.com/0xb10c/rawtx/#Input.IsLNUniliteralClosing

## Bitcoin Script

There is a [BitcoinScript][30] parser for byte slices build in.
A [ParsedBitcoinScript][31] is a slice of [ParsedOpCodes][32].
It has a `String()` method which displays the OP codes in a bitcoin developer readable format.

```go away
bs1 := BitcoinScript{0x6e, 0x87, 0x91, 0x69, 0xa7, 0x7c, 0xa7, 0x87}
pbs1 := bs1.ParseWithPanic()
fmt.Println(pbs1.String())
// -> OP_2DUP OP_EQUAL OP_NOT OP_VERIFY OP_SHA1 OP_SWAP OP_SHA1 OP_EQUAL
```

```go awayy
bs2 = BitcoinScript{byte(OpRETURN), byte(OpDATA2), byte(OpCHECKLOCKTIMEVERIFY), byte(OpDATA12)}
pbs2 = bs2.ParseWithPanic()
fmt.Println(pbs2.String())
// -> OP_RETURN OP_DATA_2(b10c)
```

The actual [OpCode][33] behind the ParsedOpCode can, but doesn't have to push data. You can check if a ParsedOpCode is 
- [x] a [signature?][34] (and [what's the SigHash?][37])
- [x] a [compressed public key?][35]
- [x] an [uncompressed public key?][36]
- [x] or either a [compressed or uncompressed public key?][38]
- [x] or compare it any other OpCode.  

[30]: https://www.godoc.org/github.com/0xb10c/rawtx/#BitcoinScript
[31]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedBitcoinScript
[32]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedOpCode
[33]: https://www.godoc.org/github.com/0xb10c/rawtx/#OpCode
[34]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedOpCode.IsSignature
[35]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedOpCode.IsCompressedPubKey
[36]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedOpCode.IsUncompressedPubKey
[37]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedOpCode.GetSigHash
[38]: https://www.godoc.org/github.com/0xb10c/rawtx/#ParsedOpCode.IsPubKey


[more]: https://www.godoc.org/github.com/0xb10c/rawtx/#pkg-index

## Running tests

Either normal test suit with coverage report in percent.

```bash
$ make test
...
PASS
coverage: 100.0% of statements
ok      github.com/0xB10C/rawtx 0.028s
```

Or with detailed coverage HTML report opened in your browser.

```bash
$ make cover
...
PASS
coverage: 100.0% of statements
ok      github.com/0xB10C/rawtx 0.028s
go tool cover -html=count.out
```

## Discalimer

Do not use in consensus related code!

## License

rawTx is licensed under a BSD 3-Clause License.