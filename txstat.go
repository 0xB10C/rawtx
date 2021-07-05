package rawtx

// TxStats contains stats about a transaction.
type TxStats struct {
	TxID                     []byte
	TxIDString               string
	Version                  int32
	Payments                 uint32
	OutAmount                int64
	VSize                    int
	Size                     int
	IsCoinbase               bool
	IsSpendingSegWit         bool
	IsSpendingNativeSegWit   bool
	IsSpendingNestedSegWit   bool
	IsBIP69Compliant         bool
	IsExplicitlyRBFSignaling bool
	Locktime                 *LocktimeStats
	InStats                  []*InputStats
	OutStats                 []*OutputStats
}

// Stats returns a *TxStats for the transaction
func (tx *Tx) Stats() *TxStats {
	txstats := &TxStats{}
	txstats.TxID = tx.Hash
	txstats.TxIDString = tx.HashString
	txstats.Version = tx.Version
	txstats.IsCoinbase = tx.IsCoinbase()
	txstats.VSize = tx.GetSizeWithoutWitness()
	txstats.Size = tx.GetSizeWithWitness()
	txstats.IsSpendingNativeSegWit = tx.IsSpendingNativeSegWit()
	txstats.IsSpendingNestedSegWit = tx.IsSpendingNestedSegWit()
	txstats.IsSpendingSegWit = tx.IsSpendingSegWit()
	txstats.IsBIP69Compliant = tx.IsBIP69Compliant()
	txstats.IsExplicitlyRBFSignaling = tx.IsExplicitlyRBFSignaling()
	txstats.Locktime = tx.LocktimeStats()

	txstats.InStats = make([]*InputStats, 0)
	for _, input := range tx.Inputs {
		txstats.InStats = append(txstats.InStats, input.InputStats())
	}

	txstats.OutStats = make([]*OutputStats, 0)
	for _, output := range tx.Outputs {
		txstats.OutStats = append(txstats.OutStats, output.OutputStats())
		txstats.OutAmount += output.Value
	}

	// The payments metric is defined as:
	// number of outputs minus one change output if more than one output
	txstats.Payments = uint32(len(txstats.OutStats))
	if len(txstats.OutStats) > 1 {
		txstats.Payments--
	}

	return txstats
}

// LocktimeStats contains stats about the transaction locktime
type LocktimeStats struct {
	Locktime      uint32
	IsEnforced    bool
	IsBlockHeight bool
	IsTimestamp   bool
}

// LocktimeStats returns a populated *LocktimeStats struct for the transaction
func (tx *Tx) LocktimeStats() *LocktimeStats {
	locktimeStats := &LocktimeStats{}
	locktimeStats.Locktime = tx.Locktime
	locktimeStats.IsBlockHeight = (tx.Locktime > 0 && tx.Locktime < 500000000)
	locktimeStats.IsTimestamp = (tx.Locktime >= 500000000)
	for _, input := range tx.Inputs {
		if input.Sequence < 0xffffffff {
			locktimeStats.IsEnforced = true
		}
	}
	return locktimeStats
}

// InputStats contains stats about a transaction input
type InputStats struct {
	Type                   InputType
	TypeString             string
	Sequence               uint32
	IsSpendingSegWit       bool
	IsSpendingNativeSegWit bool
	IsSpendingNestedSegWit bool
	IsLNUniliteralClosing  bool
	IsSpendingMultisig     bool
	MultiSigM              int
	MultiSigN              int
	SigStats               []*SignatureStats
	PubKeyStats            []*PubKeyStats
	OpCodes                []OpCode
}

// InputStats returns a populated *InputStats struct for the input
func (input *Input) InputStats() *InputStats {
	inputStats := &InputStats{}
	inputStats.Type = input.GetType()
	inputStats.TypeString = input.GetType().String()
	inputStats.Sequence = input.Sequence
	inputStats.IsSpendingNativeSegWit = input.SpendsNativeSegWit()
	inputStats.IsSpendingNestedSegWit = input.SpendsNestedSegWit()
	inputStats.IsSpendingSegWit = inputStats.IsSpendingNativeSegWit || inputStats.IsSpendingNestedSegWit
	inputStats.IsLNUniliteralClosing = input.IsLNUniliteralClosing()
	inputStats.IsSpendingMultisig = input.SpendsMultisig()
	if inputStats.IsSpendingMultisig {
		switch inputStats.Type {
		case InP2SH_P2WSH:
			_, inputStats.MultiSigM, inputStats.MultiSigN = input.GetNestedP2WSHRedeemScript().IsMultisigScript()
		case InP2SH:
			_, inputStats.MultiSigM, inputStats.MultiSigN = input.GetP2SHRedeemScript().IsMultisigScript()
		case InP2WSH:
			_, inputStats.MultiSigM, inputStats.MultiSigN = input.GetP2WSHRedeemScript().IsMultisigScript()
		}
	}

	inputStats.SigStats = make([]*SignatureStats, 0)
	inputStats.PubKeyStats = make([]*PubKeyStats, 0)
	inputStats.OpCodes = make([]OpCode, 0)
	if len(input.ScriptSig) > 0 {
		parsedScriptSig := input.ScriptSig.Parse()
		for _, opCode := range parsedScriptSig {
			if opCode.IsSignature() {
				ss := opCode.SignatureStats()
				inputStats.SigStats = append(inputStats.SigStats, ss)
			} else if opCode.IsECDSAPubKey() {
				pks := opCode.PubKeyStats()
				inputStats.PubKeyStats = append(inputStats.PubKeyStats, pks)
			} else if opCode.OpCode.IsDataPushOpCode() {
				parsedOpCodeScript := BitcoinScript(opCode.PushedData).Parse()
				for _, inScriptOpcode := range parsedOpCodeScript {
					if inScriptOpcode.IsSignature() {
						ss := inScriptOpcode.SignatureStats()
						inputStats.SigStats = append(inputStats.SigStats, ss)
					} else if inScriptOpcode.IsECDSAPubKey() {
						pks := inScriptOpcode.PubKeyStats()
						inputStats.PubKeyStats = append(inputStats.PubKeyStats, pks)
					}
				}
			}
			inputStats.OpCodes = append(inputStats.OpCodes, opCode.OpCode)
		}
	}

	if len(input.Witness) > 0 {
		for _, witnessElement := range input.Witness {
			if witnessElement.IsSignature() {
				ss := witnessElement.SignatureStats()
				inputStats.SigStats = append(inputStats.SigStats, ss)
			} else if witnessElement.IsECDSAPubKey() {
				pks := witnessElement.PubKeyStats()
				inputStats.PubKeyStats = append(inputStats.PubKeyStats, pks)
			} else if witnessElement.OpCode.IsDataPushOpCode() {
				parsedWitnessScript := BitcoinScript(witnessElement.PushedData).Parse()
				for _, witnessOpCode := range parsedWitnessScript {
					if witnessOpCode.IsSignature() {
						ss := witnessOpCode.SignatureStats()
						inputStats.SigStats = append(inputStats.SigStats, ss)
					} else if witnessOpCode.IsECDSAPubKey() {
						pks := witnessOpCode.PubKeyStats()
						inputStats.PubKeyStats = append(inputStats.PubKeyStats, pks)
					}
					inputStats.OpCodes = append(inputStats.OpCodes, witnessOpCode.OpCode)
				}
			}
			inputStats.OpCodes = append(inputStats.OpCodes, witnessElement.OpCode)
		}
	}
	return inputStats
}

// SignatureStats contains stats about a signature
type SignatureStats struct {
	Length      int
	IsECDSA     bool
	IsStrictDER bool
	HasLowR     bool
	HasLowS     bool
	SigHash     byte
}

// SignatureStats returns a populated *SignatureStats struct for the signature
// The caller must make sure that the *ParsedOpCode is a signature.
func (sigOpCode *ParsedOpCode) SignatureStats() *SignatureStats {
	sigStats := &SignatureStats{}
	sigStats.Length = len(sigOpCode.PushedData)
	sigStats.IsECDSA = sigOpCode.IsECDSASignature(false)
	if sigStats.IsECDSA {
		sigStats.IsStrictDER = sigOpCode.IsECDSASignature(true)
		ecdsaSig, ok := DeserializeECDSASignature(sigOpCode.PushedData[:sigStats.Length-1], false)
		if !ok {
			panic("Signature should evaluate to ok here")
		}
		sigStats.HasLowR = ecdsaSig.HasLowR()
		sigStats.HasLowS = ecdsaSig.HasLowS()
	}
	sigStats.SigHash = sigOpCode.GetSigHash()
	return sigStats
}

// PubKeyStats contains stats about a PubKey
type PubKeyStats struct {
	IsCompressed bool
}

// PubKeyStats returns a populated *PubKeyStats struct for the PubKey
// The caller must make sure that the *ParsedOpCode is a pubkey.
func (sigOpCode *ParsedOpCode) PubKeyStats() *PubKeyStats {
	pkStats := &PubKeyStats{}
	pkStats.IsCompressed = sigOpCode.IsCompressedECDSAPubKey()
	return pkStats
}

// OutputStats contains stats about an output
type OutputStats struct {
	Type         OutputType
	TypeString   string
	Amount       int64
	OpReturnData []byte
	PubKeyStats  []*PubKeyStats // P2MS outputs have pubkeys
	OpCodes      []OpCode
}

// OutputStats returns a populated *OutputStats struct for an output
func (out *Output) OutputStats() *OutputStats {
	outStats := &OutputStats{}
	outStats.Type = out.GetType()
	outStats.TypeString = out.GetType().String()
	outStats.Amount = out.Value
	if outStats.Type == OutOPRETURN {
		_, opCode := out.GetOPReturnData()
		outStats.OpReturnData = opCode.PushedData
	}

	outStats.PubKeyStats = make([]*PubKeyStats, 0)
	outStats.OpCodes = make([]OpCode, 0)
	parsedScriptPubKey := out.ScriptPubKey.Parse()
	for _, opCode := range parsedScriptPubKey {
		if opCode.IsECDSAPubKey() {
			pks := opCode.PubKeyStats()
			outStats.PubKeyStats = append(outStats.PubKeyStats, pks)
		}
		outStats.OpCodes = append(outStats.OpCodes, opCode.OpCode)
	}
	return outStats
}
