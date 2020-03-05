# Example txstats

The file `txstat.go` contains a simple example usage of the rawtx `tx.Stats()` functionality.
The `txstat` binary takes a raw transaction in hex as argument with `-raw` and prints a JSON object containing stats about the transaction.

```terminal
$ go run txstat.go -raw 0100000001c997a5e56e104102fa209c6a852dd90660a20b2d9c352423edce25857fcd3704000000004847304402204e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd410220181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d0901ffffffff0200ca9a3b00000000434104ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84cac00286bee0000000043410411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5cb2e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3ac00000000
{
  "TxID": "Fp4eg+kwhTORvG819gXGdUz+rVfPg4djnTtAlsVPGPQ=",
  "TxIDString": "f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16",
  "Version": 1,
  "Payments": 1,
  "OutAmount": 5000000000,
  "VSize": 275,
  "Size": 275,
  "IsCoinbase": false,
  "IsSpendingSegWit": false,
  "IsSpendingNativeSegWit": false,
  "IsSpendingNestedSegWit": false,
  "IsBIP69Compliant": true,
  "IsExplicitlyRBFSignaling": false,
  "Locktime": {
    "Locktime": 0,
    "IsEnforced": false,
    "IsBlockHeight": true,
    "IsTimestamp": false
  },
  "InStats": [
    {
      "Type": 1,
      "TypeString": "P2PK",
      "Sequence": 4294967295,
      "IsSpendingSegWit": false,
      "IsSpendingNativeSegWit": false,
      "IsSpendingNestedSegWit": false,
      "IsLNUniliteralClosing": false,
      "IsSpendingMultisig": false,
      "MultiSigM": 0,
      "MultiSigN": 0,
      "SigStats": [
        {
          "Length": 71,
          "IsECDSA": true,
          "IsStrictDER": true,
          "HasLowR": true,
          "HasLowS": true,
          "SigHash": 1
        }
      ],
      "PubKeyStats": [],
      "OpCodes": "Rw=="
    }
  ],
  "OutStats": [
    {
      "Type": 1,
      "TypeString": "P2PK",
      "Amount": 1000000000,
      "OpReturnData": null,
      "PubKeyStats": [
        {
          "IsCompressed": false
        }
      ],
      "OpCodes": "Qaw="
    },
    {
      "Type": 1,
      "TypeString": "P2PK",
      "Amount": 4000000000,
      "OpReturnData": null,
      "PubKeyStats": [
        {
          "IsCompressed": false
        }
      ],
      "OpCodes": "Qaw="
    }
  ]
}
```
