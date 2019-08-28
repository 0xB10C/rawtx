package rawtx

// OpCode representes a bitcoin operation code
type OpCode byte

// Bitcoin OP Codes.
// Mainly copied and adopted from https://github.com/btcsuite/btcd/blob/master/txscript/opcode.go
const (
	Op0                   OpCode = 0x00
	OpDATA1               OpCode = 0x01
	OpDATA2               OpCode = 0x02
	OpDATA3               OpCode = 0x03
	OpDATA4               OpCode = 0x04
	OpDATA5               OpCode = 0x05
	OpDATA6               OpCode = 0x06
	OpDATA7               OpCode = 0x07
	OpDATA8               OpCode = 0x08
	OpDATA9               OpCode = 0x09
	OpDATA10              OpCode = 0x0a
	OpDATA11              OpCode = 0x0b
	OpDATA12              OpCode = 0x0c
	OpDATA13              OpCode = 0x0d
	OpDATA14              OpCode = 0x0e
	OpDATA15              OpCode = 0x0f
	OpDATA16              OpCode = 0x10
	OpDATA17              OpCode = 0x11
	OpDATA18              OpCode = 0x12
	OpDATA19              OpCode = 0x13
	OpDATA20              OpCode = 0x14
	OpDATA21              OpCode = 0x15
	OpDATA22              OpCode = 0x16
	OpDATA23              OpCode = 0x17
	OpDATA24              OpCode = 0x18
	OpDATA25              OpCode = 0x19
	OpDATA26              OpCode = 0x1a
	OpDATA27              OpCode = 0x1b
	OpDATA28              OpCode = 0x1c
	OpDATA29              OpCode = 0x1d
	OpDATA30              OpCode = 0x1e
	OpDATA31              OpCode = 0x1f
	OpDATA32              OpCode = 0x20
	OpDATA33              OpCode = 0x21
	OpDATA34              OpCode = 0x22
	OpDATA35              OpCode = 0x23
	OpDATA36              OpCode = 0x24
	OpDATA37              OpCode = 0x25
	OpDATA38              OpCode = 0x26
	OpDATA39              OpCode = 0x27
	OpDATA40              OpCode = 0x28
	OpDATA41              OpCode = 0x29
	OpDATA42              OpCode = 0x2a
	OpDATA43              OpCode = 0x2b
	OpDATA44              OpCode = 0x2c
	OpDATA45              OpCode = 0x2d
	OpDATA46              OpCode = 0x2e
	OpDATA47              OpCode = 0x2f
	OpDATA48              OpCode = 0x30
	OpDATA49              OpCode = 0x31
	OpDATA50              OpCode = 0x32
	OpDATA51              OpCode = 0x33
	OpDATA52              OpCode = 0x34
	OpDATA53              OpCode = 0x35
	OpDATA54              OpCode = 0x36
	OpDATA55              OpCode = 0x37
	OpDATA56              OpCode = 0x38
	OpDATA57              OpCode = 0x39
	OpDATA58              OpCode = 0x3a
	OpDATA59              OpCode = 0x3b
	OpDATA60              OpCode = 0x3c
	OpDATA61              OpCode = 0x3d
	OpDATA62              OpCode = 0x3e
	OpDATA63              OpCode = 0x3f
	OpDATA64              OpCode = 0x40
	OpDATA65              OpCode = 0x41
	OpDATA66              OpCode = 0x42
	OpDATA67              OpCode = 0x43
	OpDATA68              OpCode = 0x44
	OpDATA69              OpCode = 0x45
	OpDATA70              OpCode = 0x46
	OpDATA71              OpCode = 0x47
	OpDATA72              OpCode = 0x48
	OpDATA73              OpCode = 0x49
	OpDATA74              OpCode = 0x4a
	OpDATA75              OpCode = 0x4b
	OpPUSHDATA1           OpCode = 0x4c
	OpPUSHDATA2           OpCode = 0x4d
	OpPUSHDATA4           OpCode = 0x4e
	Op1NEGATE             OpCode = 0x4f
	OpRESERVED            OpCode = 0x50
	Op1                   OpCode = 0x51
	OpTRUE                OpCode = 0x51
	Op2                   OpCode = 0x52
	Op3                   OpCode = 0x53
	Op4                   OpCode = 0x54
	Op5                   OpCode = 0x55
	Op6                   OpCode = 0x56
	Op7                   OpCode = 0x57
	Op8                   OpCode = 0x58
	Op9                   OpCode = 0x59
	Op10                  OpCode = 0x5a
	Op11                  OpCode = 0x5b
	Op12                  OpCode = 0x5c
	Op13                  OpCode = 0x5d
	Op14                  OpCode = 0x5e
	Op15                  OpCode = 0x5f
	Op16                  OpCode = 0x60
	OpNOP                 OpCode = 0x61
	OpVER                 OpCode = 0x62
	OpIF                  OpCode = 0x63
	OpNOTIF               OpCode = 0x64
	OpVERIF               OpCode = 0x65
	OpVERNOTIF            OpCode = 0x66
	OpELSE                OpCode = 0x67
	OpENDIF               OpCode = 0x68
	OpVERIFY              OpCode = 0x69
	OpRETURN              OpCode = 0x6a
	OpTOALTSTACK          OpCode = 0x6b
	OpFROMALTSTACK        OpCode = 0x6c
	Op2DROP               OpCode = 0x6d
	Op2DUP                OpCode = 0x6e
	Op3DUP                OpCode = 0x6f
	Op2OVER               OpCode = 0x70
	Op2ROT                OpCode = 0x71
	Op2SWAP               OpCode = 0x72
	OpIFDUP               OpCode = 0x73
	OpDEPTH               OpCode = 0x74
	OpDROP                OpCode = 0x75
	OpDUP                 OpCode = 0x76
	OpNIP                 OpCode = 0x77
	OpOVER                OpCode = 0x78
	OpPICK                OpCode = 0x79
	OpROLL                OpCode = 0x7a
	OpROT                 OpCode = 0x7b
	OpSWAP                OpCode = 0x7c
	OpTUCK                OpCode = 0x7d
	OpCAT                 OpCode = 0x7e
	OpSUBSTR              OpCode = 0x7f
	OpLEFT                OpCode = 0x80
	OpRIGHT               OpCode = 0x81
	OpSIZE                OpCode = 0x82
	OpINVERT              OpCode = 0x83
	OpAND                 OpCode = 0x84
	OpOR                  OpCode = 0x85
	OpXOR                 OpCode = 0x86
	OpEQUAL               OpCode = 0x87
	OpEQUALVERIFY         OpCode = 0x88
	OpRESERVED1           OpCode = 0x89
	OpRESERVED2           OpCode = 0x8a
	Op1ADD                OpCode = 0x8b
	Op1SUB                OpCode = 0x8c
	Op2MUL                OpCode = 0x8d
	Op2DIV                OpCode = 0x8e
	OpNEGATE              OpCode = 0x8f
	OpABS                 OpCode = 0x90
	OpNOT                 OpCode = 0x91
	Op0NOTEQUAL           OpCode = 0x92
	OpADD                 OpCode = 0x93
	OpSUB                 OpCode = 0x94
	OpMUL                 OpCode = 0x95
	OpDIV                 OpCode = 0x96
	OpMOD                 OpCode = 0x97
	OpLSHIFT              OpCode = 0x98
	OpRSHIFT              OpCode = 0x99
	OpBOOLAND             OpCode = 0x9a
	OpBOOLOR              OpCode = 0x9b
	OpNUMEQUAL            OpCode = 0x9c
	OpNUMEQUALVERIFY      OpCode = 0x9d
	OpNUMNOTEQUAL         OpCode = 0x9e
	OpLESSTHAN            OpCode = 0x9f
	OpGREATERTHAN         OpCode = 0xa0
	OpLESSTHANOREQUAL     OpCode = 0xa1
	OpGREATERTHANOREQUAL  OpCode = 0xa2
	OpMIN                 OpCode = 0xa3
	OpMAX                 OpCode = 0xa4
	OpWITHIN              OpCode = 0xa5
	OpRIPEMD160           OpCode = 0xa6
	OpSHA1                OpCode = 0xa7
	OpSHA256              OpCode = 0xa8
	OpHASH160             OpCode = 0xa9
	OpHASH256             OpCode = 0xaa
	OpCODESEPARATOR       OpCode = 0xab
	OpCHECKSIG            OpCode = 0xac
	OpCHECKSIGVERIFY      OpCode = 0xad
	OpCHECKMULTISIG       OpCode = 0xae
	OpCHECKMULTISIGVERIFY OpCode = 0xaf
	OpNOP1                OpCode = 0xb0
	OpCHECKLOCKTIMEVERIFY OpCode = 0xb1
	OpCHECKSEQUENCEVERIFY OpCode = 0xb2
	OpNOP4                OpCode = 0xb3
	OpNOP5                OpCode = 0xb4
	OpNOP6                OpCode = 0xb5
	OpNOP7                OpCode = 0xb6
	OpNOP8                OpCode = 0xb7
	OpNOP9                OpCode = 0xb8
	OpNOP10               OpCode = 0xb9
	OpUNKNOWN186          OpCode = 0xba
	OpUNKNOWN187          OpCode = 0xbb
	OpUNKNOWN188          OpCode = 0xbc
	OpUNKNOWN189          OpCode = 0xbd
	OpUNKNOWN190          OpCode = 0xbe
	OpUNKNOWN191          OpCode = 0xbf
	OpUNKNOWN192          OpCode = 0xc0
	OpUNKNOWN193          OpCode = 0xc1
	OpUNKNOWN194          OpCode = 0xc2
	OpUNKNOWN195          OpCode = 0xc3
	OpUNKNOWN196          OpCode = 0xc4
	OpUNKNOWN197          OpCode = 0xc5
	OpUNKNOWN198          OpCode = 0xc6
	OpUNKNOWN199          OpCode = 0xc7
	OpUNKNOWN200          OpCode = 0xc8
	OpUNKNOWN201          OpCode = 0xc9
	OpUNKNOWN202          OpCode = 0xca
	OpUNKNOWN203          OpCode = 0xcb
	OpUNKNOWN204          OpCode = 0xcc
	OpUNKNOWN205          OpCode = 0xcd
	OpUNKNOWN206          OpCode = 0xce
	OpUNKNOWN207          OpCode = 0xcf
	OpUNKNOWN208          OpCode = 0xd0
	OpUNKNOWN209          OpCode = 0xd1
	OpUNKNOWN210          OpCode = 0xd2
	OpUNKNOWN211          OpCode = 0xd3
	OpUNKNOWN212          OpCode = 0xd4
	OpUNKNOWN213          OpCode = 0xd5
	OpUNKNOWN214          OpCode = 0xd6
	OpUNKNOWN215          OpCode = 0xd7
	OpUNKNOWN216          OpCode = 0xd8
	OpUNKNOWN217          OpCode = 0xd9
	OpUNKNOWN218          OpCode = 0xda
	OpUNKNOWN219          OpCode = 0xdb
	OpUNKNOWN220          OpCode = 0xdc
	OpUNKNOWN221          OpCode = 0xdd
	OpUNKNOWN222          OpCode = 0xde
	OpUNKNOWN223          OpCode = 0xdf
	OpUNKNOWN224          OpCode = 0xe0
	OpUNKNOWN225          OpCode = 0xe1
	OpUNKNOWN226          OpCode = 0xe2
	OpUNKNOWN227          OpCode = 0xe3
	OpUNKNOWN228          OpCode = 0xe4
	OpUNKNOWN229          OpCode = 0xe5
	OpUNKNOWN230          OpCode = 0xe6
	OpUNKNOWN231          OpCode = 0xe7
	OpUNKNOWN232          OpCode = 0xe8
	OpUNKNOWN233          OpCode = 0xe9
	OpUNKNOWN234          OpCode = 0xea
	OpUNKNOWN235          OpCode = 0xeb
	OpUNKNOWN236          OpCode = 0xec
	OpUNKNOWN237          OpCode = 0xed
	OpUNKNOWN238          OpCode = 0xee
	OpUNKNOWN239          OpCode = 0xef
	OpUNKNOWN240          OpCode = 0xf0
	OpUNKNOWN241          OpCode = 0xf1
	OpUNKNOWN242          OpCode = 0xf2
	OpUNKNOWN243          OpCode = 0xf3
	OpUNKNOWN244          OpCode = 0xf4
	OpUNKNOWN245          OpCode = 0xf5
	OpUNKNOWN246          OpCode = 0xf6
	OpUNKNOWN247          OpCode = 0xf7
	OpUNKNOWN248          OpCode = 0xf8
	OpUNKNOWN249          OpCode = 0xf9
	OpSMALLINTEGER        OpCode = 0xfa
	OpPUBKEYS             OpCode = 0xfb
	OpUNKNOWN252          OpCode = 0xfc
	OpPUBKEYHASH          OpCode = 0xfd
	OpPUBKEY              OpCode = 0xfe
	OpINVALIDOPCODE       OpCode = 0xff
)

// OpCodeStringMap maps op codes to their developer readable names
var OpCodeStringMap = map[OpCode]string{
	Op0:                   "OP_0",
	OpDATA1:               "OP_DATA_1",
	OpDATA2:               "OP_DATA_2",
	OpDATA3:               "OP_DATA_3",
	OpDATA4:               "OP_DATA_4",
	OpDATA5:               "OP_DATA_5",
	OpDATA6:               "OP_DATA_6",
	OpDATA7:               "OP_DATA_7",
	OpDATA8:               "OP_DATA_8",
	OpDATA9:               "OP_DATA_9",
	OpDATA10:              "OP_DATA_10",
	OpDATA11:              "OP_DATA_11",
	OpDATA12:              "OP_DATA_12",
	OpDATA13:              "OP_DATA_13",
	OpDATA14:              "OP_DATA_14",
	OpDATA15:              "OP_DATA_15",
	OpDATA16:              "OP_DATA_16",
	OpDATA17:              "OP_DATA_17",
	OpDATA18:              "OP_DATA_18",
	OpDATA19:              "OP_DATA_19",
	OpDATA20:              "OP_DATA_20",
	OpDATA21:              "OP_DATA_21",
	OpDATA22:              "OP_DATA_22",
	OpDATA23:              "OP_DATA_23",
	OpDATA24:              "OP_DATA_24",
	OpDATA25:              "OP_DATA_25",
	OpDATA26:              "OP_DATA_26",
	OpDATA27:              "OP_DATA_27",
	OpDATA28:              "OP_DATA_28",
	OpDATA29:              "OP_DATA_29",
	OpDATA30:              "OP_DATA_30",
	OpDATA31:              "OP_DATA_31",
	OpDATA32:              "OP_DATA_32",
	OpDATA33:              "OP_DATA_33",
	OpDATA34:              "OP_DATA_34",
	OpDATA35:              "OP_DATA_35",
	OpDATA36:              "OP_DATA_36",
	OpDATA37:              "OP_DATA_37",
	OpDATA38:              "OP_DATA_38",
	OpDATA39:              "OP_DATA_39",
	OpDATA40:              "OP_DATA_40",
	OpDATA41:              "OP_DATA_41",
	OpDATA42:              "OP_DATA_42",
	OpDATA43:              "OP_DATA_43",
	OpDATA44:              "OP_DATA_44",
	OpDATA45:              "OP_DATA_45",
	OpDATA46:              "OP_DATA_46",
	OpDATA47:              "OP_DATA_47",
	OpDATA48:              "OP_DATA_48",
	OpDATA49:              "OP_DATA_49",
	OpDATA50:              "OP_DATA_50",
	OpDATA51:              "OP_DATA_51",
	OpDATA52:              "OP_DATA_52",
	OpDATA53:              "OP_DATA_53",
	OpDATA54:              "OP_DATA_54",
	OpDATA55:              "OP_DATA_55",
	OpDATA56:              "OP_DATA_56",
	OpDATA57:              "OP_DATA_57",
	OpDATA58:              "OP_DATA_58",
	OpDATA59:              "OP_DATA_59",
	OpDATA60:              "OP_DATA_60",
	OpDATA61:              "OP_DATA_61",
	OpDATA62:              "OP_DATA_62",
	OpDATA63:              "OP_DATA_63",
	OpDATA64:              "OP_DATA_64",
	OpDATA65:              "OP_DATA_65",
	OpDATA66:              "OP_DATA_66",
	OpDATA67:              "OP_DATA_67",
	OpDATA68:              "OP_DATA_68",
	OpDATA69:              "OP_DATA_69",
	OpDATA70:              "OP_DATA_70",
	OpDATA71:              "OP_DATA_71",
	OpDATA72:              "OP_DATA_72",
	OpDATA73:              "OP_DATA_73",
	OpDATA74:              "OP_DATA_74",
	OpDATA75:              "OP_DATA_75",
	OpPUSHDATA1:           "OP_PUSHDATA1",
	OpPUSHDATA2:           "OP_PUSHDATA2",
	OpPUSHDATA4:           "OP_PUSHDATA4",
	Op1NEGATE:             "OP_1NEGATE",
	OpRESERVED:            "OP_RESERVED",
	Op1:                   "OP_1",
	Op2:                   "OP_2",
	Op3:                   "OP_3",
	Op4:                   "OP_4",
	Op5:                   "OP_5",
	Op6:                   "OP_6",
	Op7:                   "OP_7",
	Op8:                   "OP_8",
	Op9:                   "OP_9",
	Op10:                  "OP_10",
	Op11:                  "OP_11",
	Op12:                  "OP_12",
	Op13:                  "OP_13",
	Op14:                  "OP_14",
	Op15:                  "OP_15",
	Op16:                  "OP_16",
	OpNOP:                 "OP_NOP",
	OpVER:                 "OP_VER",
	OpIF:                  "OP_IF",
	OpNOTIF:               "OP_NOTIF",
	OpVERIF:               "OP_VERIF",
	OpVERNOTIF:            "OP_VERNOTIF",
	OpELSE:                "OP_ELSE",
	OpENDIF:               "OP_ENDIF",
	OpVERIFY:              "OP_VERIFY",
	OpRETURN:              "OP_RETURN",
	OpTOALTSTACK:          "OP_TOALTSTACK",
	OpFROMALTSTACK:        "OP_FROMALTSTACK",
	Op2DROP:               "OP_2DROP",
	Op2DUP:                "OP_2DUP",
	Op3DUP:                "OP_3DUP",
	Op2OVER:               "OP_2OVER",
	Op2ROT:                "OP_2ROT",
	Op2SWAP:               "OP_2SWAP",
	OpIFDUP:               "OP_IFDUP",
	OpDEPTH:               "OP_DEPTH",
	OpDROP:                "OP_DROP",
	OpDUP:                 "OP_DUP",
	OpNIP:                 "OP_NIP",
	OpOVER:                "OP_OVER",
	OpPICK:                "OP_PICK",
	OpROLL:                "OP_ROLL",
	OpROT:                 "OP_ROT",
	OpSWAP:                "OP_SWAP",
	OpTUCK:                "OP_TUCK",
	OpCAT:                 "OP_CAT",
	OpSUBSTR:              "OP_SUBSTR",
	OpLEFT:                "OP_LEFT",
	OpRIGHT:               "OP_RIGHT",
	OpSIZE:                "OP_SIZE",
	OpINVERT:              "OP_INVERT",
	OpAND:                 "OP_AND",
	OpOR:                  "OP_OR",
	OpXOR:                 "OP_XOR",
	OpEQUAL:               "OP_EQUAL",
	OpEQUALVERIFY:         "OP_EQUALVERIFY",
	OpRESERVED1:           "OP_RESERVED1",
	OpRESERVED2:           "OP_RESERVED2",
	Op1ADD:                "OP_1ADD",
	Op1SUB:                "OP_1SUB",
	Op2MUL:                "OP_2MUL",
	Op2DIV:                "OP_2DIV",
	OpNEGATE:              "OP_NEGATE",
	OpABS:                 "OP_ABS",
	OpNOT:                 "OP_NOT",
	Op0NOTEQUAL:           "OP_0NOTEQUAL",
	OpADD:                 "OP_ADD",
	OpSUB:                 "OP_SUB",
	OpMUL:                 "OP_MUL",
	OpDIV:                 "OP_DIV",
	OpMOD:                 "OP_MOD",
	OpLSHIFT:              "OP_LSHIFT",
	OpRSHIFT:              "OP_RSHIFT",
	OpBOOLAND:             "OP_BOOLAND",
	OpBOOLOR:              "OP_BOOLOR",
	OpNUMEQUAL:            "OP_NUMEQUAL",
	OpNUMEQUALVERIFY:      "OP_NUMEQUALVERIFY",
	OpNUMNOTEQUAL:         "OP_NUMNOTEQUAL",
	OpLESSTHAN:            "OP_LESSTHAN",
	OpGREATERTHAN:         "OP_GREATERTHAN",
	OpLESSTHANOREQUAL:     "OP_LESSTHANOREQUAL",
	OpGREATERTHANOREQUAL:  "OP_GREATERTHANOREQUAL",
	OpMIN:                 "OP_MIN",
	OpMAX:                 "OP_MAX",
	OpWITHIN:              "OP_WITHIN",
	OpRIPEMD160:           "OP_RIPEMD160",
	OpSHA1:                "OP_SHA1",
	OpSHA256:              "OP_SHA256",
	OpHASH160:             "OP_HASH160",
	OpHASH256:             "OP_HASH256",
	OpCODESEPARATOR:       "OP_CODESEPARATOR",
	OpCHECKSIG:            "OP_CHECKSIG",
	OpCHECKSIGVERIFY:      "OP_CHECKSIGVERIFY",
	OpCHECKMULTISIG:       "OP_CHECKMULTISIG",
	OpCHECKMULTISIGVERIFY: "OP_CHECKMULTISIGVERIFY",
	OpNOP1:                "OP_NOP1",
	OpCHECKLOCKTIMEVERIFY: "OP_CHECKLOCKTIMEVERIFY",
	OpCHECKSEQUENCEVERIFY: "OP_CHECKSEQUENCEVERIFY",
	OpNOP4:                "OP_NOP4",
	OpNOP5:                "OP_NOP5",
	OpNOP6:                "OP_NOP6",
	OpNOP7:                "OP_NOP7",
	OpNOP8:                "OP_NOP8",
	OpNOP9:                "OP_NOP9",
	OpNOP10:               "OP_NOP10",
	OpUNKNOWN186:          "OP_UNKNOWN186",
	OpUNKNOWN187:          "OP_UNKNOWN187",
	OpUNKNOWN188:          "OP_UNKNOWN188",
	OpUNKNOWN189:          "OP_UNKNOWN189",
	OpUNKNOWN190:          "OP_UNKNOWN190",
	OpUNKNOWN191:          "OP_UNKNOWN191",
	OpUNKNOWN192:          "OP_UNKNOWN192",
	OpUNKNOWN193:          "OP_UNKNOWN193",
	OpUNKNOWN194:          "OP_UNKNOWN194",
	OpUNKNOWN195:          "OP_UNKNOWN195",
	OpUNKNOWN196:          "OP_UNKNOWN196",
	OpUNKNOWN197:          "OP_UNKNOWN197",
	OpUNKNOWN198:          "OP_UNKNOWN198",
	OpUNKNOWN199:          "OP_UNKNOWN199",
	OpUNKNOWN200:          "OP_UNKNOWN200",
	OpUNKNOWN201:          "OP_UNKNOWN201",
	OpUNKNOWN202:          "OP_UNKNOWN202",
	OpUNKNOWN203:          "OP_UNKNOWN203",
	OpUNKNOWN204:          "OP_UNKNOWN204",
	OpUNKNOWN205:          "OP_UNKNOWN205",
	OpUNKNOWN206:          "OP_UNKNOWN206",
	OpUNKNOWN207:          "OP_UNKNOWN207",
	OpUNKNOWN208:          "OP_UNKNOWN208",
	OpUNKNOWN209:          "OP_UNKNOWN209",
	OpUNKNOWN210:          "OP_UNKNOWN210",
	OpUNKNOWN211:          "OP_UNKNOWN211",
	OpUNKNOWN212:          "OP_UNKNOWN212",
	OpUNKNOWN213:          "OP_UNKNOWN213",
	OpUNKNOWN214:          "OP_UNKNOWN214",
	OpUNKNOWN215:          "OP_UNKNOWN215",
	OpUNKNOWN216:          "OP_UNKNOWN216",
	OpUNKNOWN217:          "OP_UNKNOWN217",
	OpUNKNOWN218:          "OP_UNKNOWN218",
	OpUNKNOWN219:          "OP_UNKNOWN219",
	OpUNKNOWN220:          "OP_UNKNOWN220",
	OpUNKNOWN221:          "OP_UNKNOWN221",
	OpUNKNOWN222:          "OP_UNKNOWN222",
	OpUNKNOWN223:          "OP_UNKNOWN223",
	OpUNKNOWN224:          "OP_UNKNOWN224",
	OpUNKNOWN225:          "OP_UNKNOWN225",
	OpUNKNOWN226:          "OP_UNKNOWN226",
	OpUNKNOWN227:          "OP_UNKNOWN227",
	OpUNKNOWN228:          "OP_UNKNOWN228",
	OpUNKNOWN229:          "OP_UNKNOWN229",
	OpUNKNOWN230:          "OP_UNKNOWN230",
	OpUNKNOWN231:          "OP_UNKNOWN231",
	OpUNKNOWN232:          "OP_UNKNOWN232",
	OpUNKNOWN233:          "OP_UNKNOWN233",
	OpUNKNOWN234:          "OP_UNKNOWN234",
	OpUNKNOWN235:          "OP_UNKNOWN235",
	OpUNKNOWN236:          "OP_UNKNOWN236",
	OpUNKNOWN237:          "OP_UNKNOWN237",
	OpUNKNOWN238:          "OP_UNKNOWN238",
	OpUNKNOWN239:          "OP_UNKNOWN239",
	OpUNKNOWN240:          "OP_UNKNOWN240",
	OpUNKNOWN241:          "OP_UNKNOWN241",
	OpUNKNOWN242:          "OP_UNKNOWN242",
	OpUNKNOWN243:          "OP_UNKNOWN243",
	OpUNKNOWN244:          "OP_UNKNOWN244",
	OpUNKNOWN245:          "OP_UNKNOWN245",
	OpUNKNOWN246:          "OP_UNKNOWN246",
	OpUNKNOWN247:          "OP_UNKNOWN247",
	OpUNKNOWN248:          "OP_UNKNOWN248",
	OpUNKNOWN249:          "OP_UNKNOWN249",
	OpSMALLINTEGER:        "OP_SMALLINTEGER",
	OpPUBKEYS:             "OP_PUBKEYS",
	OpUNKNOWN252:          "OP_UNKNOWN252",
	OpPUBKEYHASH:          "OP_PUBKEYHASH",
	OpPUBKEY:              "OP_PUBKEY",
	OpINVALIDOPCODE:       "OP_INVALIDOPCODE",
}
