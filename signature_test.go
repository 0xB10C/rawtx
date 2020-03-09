package rawtx

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
)

type TestDERSignature struct {
	ok           bool
	reqStrictDER bool
	r            string
	s            string
	sigBytes     string
}

func getTestDERSignatures() []TestDERSignature {
	return []TestDERSignature{
		TestDERSignature{
			// BIP66 example 1 (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "d7a0417c3f6d1a15094d1cf2a3378ca0503eb8a57630953a9e2987e21ddd0a65",
			s:            "7a6266d686c99090920249991d3d42065b6d43eb70187b219c0db82e4f94d1a2",
			sigBytes:     "30440220d7a0417c3f6d1a15094d1cf2a3378ca0503eb8a57630953a9e2987e21ddd0a6502207a6266d686c99090920249991d3d42065b6d43eb70187b219c0db82e4f94d1a2",
		},
		TestDERSignature{
			// BIP66 example 1 (not strict DER)
			ok:           false, // erroring because the sig is not reqStrictDER
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "30440220d7a0417c3f6d1a15094d1cf2a3378ca0503eb8a57630953a9e2987e21ddd0a6502207a6266d686c99090920249991d3d42065b6d43eb70187b219c0db82e4f94d1a2",
		},
		TestDERSignature{
			// BIP66 example 2 (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "8e43c0b91f7c1e5bc58e41c8185f8a6086e111b0090187968a86f2822462d3c9",
			s:            "0a58f4076b1133b18ff1dc83ee51676e44c60cc608d9534e0df5ace0424fc0be",
			sigBytes:     "304402208e43c0b91f7c1e5bc58e41c8185f8a6086e111b0090187968a86f2822462d3c902200a58f4076b1133b18ff1dc83ee51676e44c60cc608d9534e0df5ace0424fc0be",
		},
		TestDERSignature{
			// BIP66 example 2 (not strict DER)
			ok:           false, // erroring because the sig is not reqStrictDER
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "304402208e43c0b91f7c1e5bc58e41c8185f8a6086e111b0090187968a86f2822462d3c902200a58f4076b1133b18ff1dc83ee51676e44c60cc608d9534e0df5ace0424fc0be",
		},
		TestDERSignature{
			// BIP66 example 7 (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "cae00b1444babfbf6071b0ba8707f6bd373da3df494d6e74119b0430c5db8105",
			s:            "5d5231b8c5939c8ff0c82242656d6e06edb073d42af336c99fe8837c36ea39d5",
			sigBytes:     "30440220cae00b1444babfbf6071b0ba8707f6bd373da3df494d6e74119b0430c5db810502205d5231b8c5939c8ff0c82242656d6e06edb073d42af336c99fe8837c36ea39d5",
		},
		TestDERSignature{
			// BIP66 example 7 (not strict DER)
			ok:           false,
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "30440220cae00b1444babfbf6071b0ba8707f6bd373da3df494d6e74119b0430c5db810502205d5231b8c5939c8ff0c82242656d6e06edb073d42af336c99fe8837c36ea39d5",
		},
		TestDERSignature{
			// BIP66 example 10 (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "afa76a8f60622f813b05711f051c6c3407e32d1b1b70b0576c1f01b54e4c05c7",
			s:            "0d58e9df044fd1845cabfbeef6e624ba0401daf7d7e084736f9ff601c3783bf5",
			sigBytes:     "30440220afa76a8f60622f813b05711f051c6c3407e32d1b1b70b0576c1f01b54e4c05c702200d58e9df044fd1845cabfbeef6e624ba0401daf7d7e084736f9ff601c3783bf501",
		},
		TestDERSignature{
			// BIP66 example 10 (not strict DER)
			ok:           false,
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "30440220afa76a8f60622f813b05711f051c6c3407e32d1b1b70b0576c1f01b54e4c05c702200d58e9df044fd1845cabfbeef6e624ba0401daf7d7e084736f9ff601c3783bf501",
		},
		TestDERSignature{
			// only 0x00 as signature
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "00",
		},
		TestDERSignature{
			// invalid DER value maker for r (0x03 instead of 0x02)
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "304403201b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260022045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// cut off signature after the r element
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "304402201b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260",
		},
		TestDERSignature{
			// r length of 0
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "30240200022045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// S length of 0
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "302402201b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e92600200",
		},
		TestDERSignature{
			// wrong signature length
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "304302001b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260022045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// no r element zero padding on (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "000096de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260",
			s:            "45db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
			sigBytes:     "30440220000096de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260022045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// no r element zero padding on (not strict DER)
			ok:           false,
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "30440220000096de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260022045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// invalid DER value maker for S (0x03 instead of 0x02)
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "304402201b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260032045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// negative S element (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "4d62295117b6e7b645e4c867092591d0501aac9f4f6a68e27e3b07658849161b",
			s:            "8b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44",
			sigBytes:     "304402204d62295117b6e7b645e4c867092591d0501aac9f4f6a68e27e3b07658849161b02208b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44",
		},
		TestDERSignature{
			// negative S element (not strict DER)
			ok:           false, // erroring because the sig is not reqStrictDER
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "304402204d62295117b6e7b645e4c867092591d0501aac9f4f6a68e27e3b07658849161b02208b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44",
		},
		TestDERSignature{
			// too long signature
			ok:       false,
			r:        "",
			s:        "",
			sigBytes: "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
		TestDERSignature{
			// negative r value (not strict DER)
			ok:           true,
			reqStrictDER: false,
			r:            "f00a77260d34ec2f0c59621dc710f58169d0ca06df1a88cd4b1f1b97bd46991b",
			s:            "1ee220c7e04f26aed03f94aa97fb09ca5627163bf4ba07e6979972ec737db226",
			sigBytes:     "30440220f00a77260d34ec2f0c59621dc710f58169d0ca06df1a88cd4b1f1b97bd46991b02201ee220c7e04f26aed03f94aa97fb09ca5627163bf4ba07e6979972ec737db226",
		},
		TestDERSignature{
			// negative r value (not strict DER)
			ok:           false, // erroring because the sig is not reqStrictDER
			reqStrictDER: true,
			r:            "",
			s:            "",
			sigBytes:     "30440220f00a77260d34ec2f0c59621dc710f58169d0ca06df1a88cd4b1f1b97bd46991b02201ee220c7e04f26aed03f94aa97fb09ca5627163bf4ba07e6979972ec737db226",
		},
		TestDERSignature{
			// test that r begining with 00 is allowed if it would otherwise be interpreted as negative number
			reqStrictDER: true,
			ok:           true,
			r:            "00dcbf285834e8d6ebec4b981fa77aaf88af7d8172a03c62d645c653caf0df6212",
			s:            "14f66da2f03a6c5a228c2f256988aa81ef1006a53cff6d6d1830c4a0bd51d02d",
			sigBytes:     "3045022100dcbf285834e8d6ebec4b981fa77aaf88af7d8172a03c62d645c653caf0df6212022014f66da2f03a6c5a228c2f256988aa81ef1006a53cff6d6d1830c4a0bd51d02d",
		},
		TestDERSignature{
			// test a random valid signature
			reqStrictDER: true,
			ok:           true,
			r:            "4d62295117b6e7b645e4c867092591d0501aac9f4f6a68e27e3b07658849161b",
			s:            "4b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44",
			sigBytes:     "304402204d62295117b6e7b645e4c867092591d0501aac9f4f6a68e27e3b07658849161b02204b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44",
		},
		TestDERSignature{
			// test a random valid signature
			reqStrictDER: true,
			ok:           true,
			r:            "3c02bd6f63e479c7c6d1dc7d943b584ba203c6f150379c789f84f8a5f33c9512",
			s:            "1ccf01bbeb2c5f68b178ab96a1a564af6609f73387b9355a62655448b6a27a43",
			sigBytes:     "304402203c02bd6f63e479c7c6d1dc7d943b584ba203c6f150379c789f84f8a5f33c951202201ccf01bbeb2c5f68b178ab96a1a564af6609f73387b9355a62655448b6a27a43",
		},
		TestDERSignature{
			// test a random valid signature
			reqStrictDER: true,
			ok:           true,
			r:            "1b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260",
			s:            "45db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
			sigBytes:     "304402201b3996de096c395d3185abe4e5b7a6f932475502a61e41979d33b2661d9e9260022045db2c937f1e13013d87f7ad7f4dbf8285455203ee37bfc164ee54091b7603a3",
		},
		TestDERSignature{
			// Satoshi's first sig in Block 170
			reqStrictDER: true,
			ok:           true,
			r:            "4e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd41",
			s:            "181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d09",
			sigBytes:     "304402204e45e16932b8af514961a1d3a1a25fdf3f4f7732e9d624c6c61548ab5fb8cd410220181522ec8eca07de4860a4acdd12909d831cc56cbbac4622082221a8768d1d09",
		},
		TestDERSignature{
			// Input 0 of ad32e67d438eb46c80f026e1788247e07ad70a51833e7adcbff4812c1873a791
			reqStrictDER: true,
			ok:           true,
			r:            "00e0e99a54ec1fe04c75693252968affaef0036209dbf172d0b6e65887601de320",
			s:            "008d12612846ea1424af09bd0caff3979c6f9d76da138e3f7b78b4222b544ac4b4",
			sigBytes:     "3046022100e0e99a54ec1fe04c75693252968affaef0036209dbf172d0b6e65887601de3200221008d12612846ea1424af09bd0caff3979c6f9d76da138e3f7b78b4222b544ac4b4",
		},
		TestDERSignature{
			// Input 3 of 015ee1a8a7f1e2a4a0939b119f6d96e3db3fa759489c2fad6db50cb11428590d
			reqStrictDER: false,
			ok:           true,
			r:            "75a7269ad5506a755ca2ccb586fbf4bcf4177546f0ba58aa65051a8efe840b43",
			s:            "00d99c68e841e9a495ada20d0c7f870b3388d724d3be5b0ada5768d75f62179fc6",
			sigBytes:     "3045022075a7269ad5506a755ca2ccb586fbf4bcf4177546f0ba58aa65051a8efe840b43022100d99c68e841e9a495ada20d0c7f870b3388d724d3be5b0ada5768d75f62179fc600",
		},
		TestDERSignature{
			// Input 3 of 015ee1a8a7f1e2a4a0939b119f6d96e3db3fa759489c2fad6db50cb11428590d
			reqStrictDER: true,
			ok:           false,
			r:            "",
			s:            "",
			sigBytes:     "3045022075a7269ad5506a755ca2ccb586fbf4bcf4177546f0ba58aa65051a8efe840b43022100d99c68e841e9a495ada20d0c7f870b3388d724d3be5b0ada5768d75f62179fc600",
		},
		TestDERSignature{
			// Input 0 of 39baeb3b2579dac22cec858be3a4d70d8d229206127b43fa4133ed63fb7b1b40
			reqStrictDER: false,
			ok:           true,
			r:            "0064ddabf1af28c21103cf61cf19dbef814aff2eba0440c5e5e20a605d16d780",
			s:            "00f45c4bc6a4ab317dc3a600129fc6a87a0df6329dbc71c5fcca9effdb30f18579",
			sigBytes:     "304502200064ddabf1af28c21103cf61cf19dbef814aff2eba0440c5e5e20a605d16d780022100f45c4bc6a4ab317dc3a600129fc6a87a0df6329dbc71c5fcca9effdb30f18579",
		},
		TestDERSignature{
			// Input 0 of f4597ab5b6d45ba3a04486f3edf1a27f9f2cc3ab23300eb16d6b7067b8cf47dd
			reqStrictDER: true,
			ok:           true,
			r:            "00eb232172f28bc933f8bd0b5c40c83f98d01792ed45c4832a887d4e95bff3322a",
			s:            "00f4d7ee1d3b8f71995f197ffcdd4e5b2327a735c5edca12724502051f33cb18c0",
			sigBytes:     "3046022100eb232172f28bc933f8bd0b5c40c83f98d01792ed45c4832a887d4e95bff3322a022100f4d7ee1d3b8f71995f197ffcdd4e5b2327a735c5edca12724502051f33cb18c0",
		},
		TestDERSignature{
			// Input 0 of 6ca68cc6ae0a6dc7bfca12e794a9cd734ea7fbbefe181147bb520c39d6c21146
			reqStrictDER: true,
			ok:           true,
			r:            "00eced8afd6a57b4f4c8374113a5bb15d4b098022b864bb383186d8774875769e9",
			s:            "00ecee7d57989e230281aa1ccbc472e791143d181800f8c16f692c0a77183da670",
			sigBytes:     "3046022100eced8afd6a57b4f4c8374113a5bb15d4b098022b864bb383186d8774875769e9022100ecee7d57989e230281aa1ccbc472e791143d181800f8c16f692c0a77183da670",
		},
		TestDERSignature{
			// Input 1 of efa73eee0de1c8325973db5c168d03311e1a0a930ec47248e6c534a950251134
			reqStrictDER: true,
			ok:           true,
			r:            "00f8776cb0fd28a52b8d8a5252c976af17d45ce1c9869f01883293c85e3b4ad4ea",
			s:            "00b7a3c6ce5a38b51ab0ebe2464c1dd1143b67c85893f1ef583b7287d2469389e6",
			sigBytes:     "3046022100f8776cb0fd28a52b8d8a5252c976af17d45ce1c9869f01883293c85e3b4ad4ea022100b7a3c6ce5a38b51ab0ebe2464c1dd1143b67c85893f1ef583b7287d2469389e6",
		},
		TestDERSignature{
			// Input 0 of fe7030c4ea8cb2904b641441ed725874cb9157b61cf4a93c9161d855d10ad7c8
			reqStrictDER: true,
			ok:           true,
			r:            "0096b04409da48353be6cceb273835c9c402c4a5afb120f4720e730f7661c85061",
			s:            "00c6c240f8e3d6696a4b74783e153447a551f3bf629a280c262d02b0f46b678372",
			sigBytes:     "304602210096b04409da48353be6cceb273835c9c402c4a5afb120f4720e730f7661c85061022100c6c240f8e3d6696a4b74783e153447a551f3bf629a280c262d02b0f46b678372",
		},
		TestDERSignature{
			// Input 0 of fe7030c4ea8cb2904b641441ed725874cb9157b61cf4a93c9161d855d10ad7c8
			reqStrictDER: true,
			ok:           true,
			r:            "00c2053ce55d823212a45aad1e8c63369f7a37483c80187b6cfe3e75f82f0325db",
			s:            "00e0389192e17666088d9fa04aa86e357d2406b0fbb9cdc4d3bfef49740ce44003",
			sigBytes:     "3046022100c2053ce55d823212a45aad1e8c63369f7a37483c80187b6cfe3e75f82f0325db022100e0389192e17666088d9fa04aa86e357d2406b0fbb9cdc4d3bfef49740ce44003",
		},
	}
}

func (testSig *TestDERSignature) decodeHexStrings(t *testing.T) (r []byte, s []byte, sigBytes []byte) {
	r, err := hex.DecodeString(testSig.r)
	if err != nil {
		t.Error(err)
	}

	s, err = hex.DecodeString(testSig.s)
	if err != nil {
		t.Error(err)
	}

	sigBytes, err = hex.DecodeString(testSig.sigBytes)
	if err != nil {
		t.Error(err)
	}

	return
}

func TestDeserializeAsDERECDSA(t *testing.T) {
	for _, testSig := range getTestDERSignatures() {
		r, s, sigBytes := testSig.decodeHexStrings(t)

		sig, ok := DeserializeECDSASignature(sigBytes, testSig.reqStrictDER)

		if testSig.ok != ok {
			t.Errorf("Expected signature to be ok=%t but got ok=%t for Signature=%s", testSig.ok, ok, testSig.sigBytes)
		}

		if !bytes.Equal(r, sig.r) {
			t.Errorf("Expected r to be r=%v but got r=%v for Signature=%s", r, sig.r, testSig.sigBytes)
		}

		if !bytes.Equal(s, sig.s) {
			t.Errorf("Expected s to be s=%v but got s=%v for Signature=%s", s, sig.s, testSig.sigBytes)
		}
	}
}

func TestHasLowS(t *testing.T) {
	halfCurveOrderSecp256k1 := getHalfCurveOrderSecp256k1()

	sigEqualS := ECDSASignature{
		r: nil,
		s: halfCurveOrderSecp256k1.Bytes(),
	}
	if !sigEqualS.HasLowS() {
		t.Error("Expected sigEqualS to count as HasLowS()=true")
	}

	lower := big.NewInt(0)
	lower.Sub(halfCurveOrderSecp256k1, big.NewInt(1))
	sigLowerS := ECDSASignature{
		r: nil,
		s: lower.Bytes(),
	}
	if !sigLowerS.HasLowS() {
		t.Error("Expected sigLowerS to be as HasLowS()=true")
	}

	higher := big.NewInt(0)
	higher.Add(halfCurveOrderSecp256k1, big.NewInt(1))
	sigHigherS := ECDSASignature{
		r: nil,
		s: higher.Bytes(),
	}
	if sigHigherS.HasLowS() {
		t.Error("Expected sigHigherS to be as HasLowS()=false")
	}

	sHigh, _ := hex.DecodeString("8b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44")
	sigHighS := ECDSASignature{
		r: nil,
		s: sHigh,
	}
	if sigHighS.HasLowS() {
		t.Error("Expected sigHigherS to be as HasLowS()=false")
	}

}

func TestHasLowR(t *testing.T) {
	sigEqualR := ECDSASignature{
		r: []byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		s: nil,
	}
	if sigEqualR.HasLowR() {
		t.Error("Expected sigEqualR to count as HasLowR()=false")
	}

	sigLowerR := ECDSASignature{
		r: []byte{0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		s: nil,
	}
	if !sigLowerR.HasLowR() {
		t.Error("Expected sigLowerR to count as HasLowR()=true")
	}

	sigHigherR := ECDSASignature{
		r: []byte{0x81, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		s: nil,
	}
	if sigHigherR.HasLowR() {
		t.Error("Expected sigHigherR to count as HasLowR()=false")
	}

	sigHigherRWithPadding := ECDSASignature{
		r: []byte{0x00, 0x81, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		s: nil,
	}
	if sigHigherRWithPadding.HasLowR() {
		t.Error("Expected sigHigherRWithPadding to count as HasLowR()=false")
	}

	sigLowerRWithPadding := ECDSASignature{
		r: []byte{0x00, 0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		s: nil,
	}
	if !sigLowerRWithPadding.HasLowR() {
		t.Error("Expected sigLowerRWithPadding to count as HasLowR()=true")
	}

	sigNil := ECDSASignature{
		r: nil,
		s: nil,
	}
	if !sigNil.HasLowR() {
		t.Error("Expected sigNil to count as HasLowR()=true")
	}

	rHigh, _ := hex.DecodeString("8b87598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44")
	sigHighR := ECDSASignature{
		r: rHigh,
		s: nil,
	}
	if sigHighR.HasLowR() {
		t.Error("Expected sigHighR to be as HasLowR()=false")
	}

	rLow, _ := hex.DecodeString("7187598c7ce7e24617b7ce861779ef9666e3992180c18569ee4dddf05ddc1e44")
	sigLowR := ECDSASignature{
		r: rLow,
		s: nil,
	}
	if !sigLowR.HasLowR() {
		t.Error("Expected sigLowR to be as HasLowR()=false")
	}
}
