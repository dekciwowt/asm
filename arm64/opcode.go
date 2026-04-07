package arm64

import "fmt"

// 31     30       29       28-24    23-21    20-16    15-10    9-5      4-0
// +------+--------+--------+--------+--------+--------+--------+--------+--------+
// |  sf  |  oper  |  sign  |  opt1  |  opt2  |  opt3  |  opt4  |  opt5  |  opt6  |
// +------+--------+--------+--------+--------+--------+--------+--------+--------+
type Opcode uint32

const (
	opcodeOpt6Size uint32 = 5
	opcodeOpt5Size uint32 = 5
	opcodeOpt4Size uint32 = 6
	opcodeOpt3Size uint32 = 5
	opcodeOpt2Size uint32 = 3
	opcodeOpt1Size uint32 = 5
	opcodeSignSize uint32 = 1
	opcodeOperSize uint32 = 1
	opcodeSFSize   uint32 = 1

	opcodeOpt6Mask uint32 = (1 << opcodeOpt6Size) - 1
	opcodeOpt5Mask uint32 = (1 << opcodeOpt5Size) - 1
	opcodeOpt4Mask uint32 = (1 << opcodeOpt4Size) - 1
	opcodeOpt3Mask uint32 = (1 << opcodeOpt3Size) - 1
	opcodeOpt2Mask uint32 = (1 << opcodeOpt2Size) - 1
	opcodeOpt1Mask uint32 = (1 << opcodeOpt1Size) - 1
	opcodeSignMask uint32 = (1 << opcodeSignSize) - 1
	opcodeOperMask uint32 = (1 << opcodeOperSize) - 1
	opcodeSFMask   uint32 = (1 << opcodeSFSize) - 1

	opcodeOpt6Pos uint32 = 0
	opcodeOpt5Pos uint32 = 5
	opcodeOpt4Pos uint32 = 10
	opcodeOpt3Pos uint32 = 16
	opcodeOpt2Pos uint32 = 21
	opcodeOpt1Pos uint32 = 24
	opcodeSignPos uint32 = 29
	opcodeOperPos uint32 = 30
	opcodeSFPos   uint32 = 31
)

func opcode(opt6, opt5, opt4, opt3, opt2, opt1, sign, oper, sf uint8) Opcode {
	var op Opcode

	op = set(op, opt6, opcodeOpt6Mask, opcodeOpt6Pos)
	op = set(op, opt5, opcodeOpt5Mask, opcodeOpt5Pos)
	op = set(op, opt4, opcodeOpt4Mask, opcodeOpt4Pos)
	op = set(op, opt3, opcodeOpt3Mask, opcodeOpt3Pos)
	op = set(op, opt2, opcodeOpt2Mask, opcodeOpt2Pos)
	op = set(op, opt1, opcodeOpt1Mask, opcodeOpt1Pos)
	op = set(op, sign, opcodeSignMask, opcodeSignPos)
	op = set(op, oper, opcodeOperMask, opcodeOperPos)
	op = set(op, sf, opcodeSFMask, opcodeSFPos)

	return op
}

const (
	catLogicReg uint8 = 0x0A
	catArithReg uint8 = 0x0B
	catArithImm uint8 = 0x11
	catLogicImm uint8 = 0x12
	catNSources uint8 = 0x1A
)

var (
	// Data-Processing (2 source)

	OpUDIVw = opcode(0x0, 0x0, 0x2, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit unsigned divide
	OpSDIVw = opcode(0x0, 0x0, 0x3, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit signed divide
	OpLSLVw = opcode(0x0, 0x0, 0x8, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit logical shift left variable
	OpLSRVw = opcode(0x0, 0x0, 0x9, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit logical shift right variable
	OpASRVw = opcode(0x0, 0x0, 0xA, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit arithmetic shift right variable
	OpRORVw = opcode(0x0, 0x0, 0xB, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit rotate right variable
	OpUDIVx = opcode(0x0, 0x0, 0x2, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit unsigned divide
	OpSDIVx = opcode(0x0, 0x0, 0x3, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit signed divide
	OpLSLVx = opcode(0x0, 0x0, 0x8, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit logical shift left variable
	OpLSRVx = opcode(0x0, 0x0, 0x9, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit logical shift right variable
	OpASRVx = opcode(0x0, 0x0, 0xA, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit arithmetic shift right variable
	OpRORVx = opcode(0x0, 0x0, 0xB, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit rotate right variable

	// Data-Processing (2 source), feature CRC32

	OpCRC32B  = opcode(0x0, 0x0, 0x10, 0x0, 0x6, catNSources, 0, 0, 0) // CRC32 checksum
	OpCRC32H  = opcode(0x0, 0x0, 0x11, 0x0, 0x6, catNSources, 0, 0, 0)
	OpCRC32W  = opcode(0x0, 0x0, 0x12, 0x0, 0x6, catNSources, 0, 0, 0)
	OpCRC32CB = opcode(0x0, 0x0, 0x14, 0x0, 0x6, catNSources, 0, 0, 0)
	OpCRC32CH = opcode(0x0, 0x0, 0x15, 0x0, 0x6, catNSources, 0, 0, 0)
	OpCRC32CW = opcode(0x0, 0x0, 0x16, 0x0, 0x6, catNSources, 0, 0, 0)
	OpCRC32X  = opcode(0x0, 0x0, 0x13, 0x0, 0x6, catNSources, 0, 0, 1)
	OpCRC32CX = opcode(0x0, 0x0, 0x17, 0x0, 0x6, catNSources, 0, 0, 1)

	// Data-Processing (2 source), feature CSSC

	OpSMAXw = opcode(0x0, 0x0, 0x18, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit signed maximum
	OpUMAXw = opcode(0x0, 0x0, 0x19, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit unsigned maximum
	OpSMINw = opcode(0x0, 0x0, 0x1A, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit signed minimum
	OpUMINw = opcode(0x0, 0x0, 0x1B, 0x0, 0x6, catNSources, 0, 0, 0) // 32-bit unsigned minimum
	OpSMAXx = opcode(0x0, 0x0, 0x18, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit signed maximum
	OpUMAXx = opcode(0x0, 0x0, 0x19, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit unsigned maximum
	OpSMINx = opcode(0x0, 0x0, 0x1A, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit signed minimum
	OpUMINx = opcode(0x0, 0x0, 0x1B, 0x0, 0x6, catNSources, 0, 0, 1) // 64-bit unsigned minimum

	// Data-Processing (2 source), feature MTE

	OpSUBPx  = opcode(0x0, 0x0, 0x0, 0x0, 0x6, catNSources, 0, 0, 1) // subtract pointer
	OpSUBPSx = opcode(0x0, 0x0, 0x0, 0x0, 0x6, catNSources, 1, 0, 1) // subtract pointer, setting flags
	OpIRGx   = opcode(0x0, 0x0, 0x4, 0x0, 0x6, catNSources, 0, 0, 1) // insert random tag
	OpGMIx   = opcode(0x0, 0x0, 0x5, 0x0, 0x6, catNSources, 0, 0, 1) // tag mask insert

	// Data-Processing (2 source), feature PAuth

	OpPACGAx = opcode(0x0, 0x0, 0xC, 0x0, 0x6, catNSources, 0, 0, 1) // pointer auth code, using generic key

	// Data-Processing (1 source)

	OpRBITw  = opcode(0x0, 0x0, 0x0, 0x0, 0x6, catNSources, 0, 1, 0) // 32-bit reverse bits
	OpREV16w = opcode(0x0, 0x0, 0x1, 0x0, 0x6, catNSources, 0, 1, 0)
	OpREVw   = opcode(0x0, 0x0, 0x2, 0x0, 0x6, catNSources, 0, 1, 0)
	OpCLZw   = opcode(0x0, 0x0, 0x4, 0x0, 0x6, catNSources, 0, 1, 0)
	OpCLSw   = opcode(0x0, 0x0, 0x5, 0x0, 0x6, catNSources, 0, 1, 0)
	OpRBITx  = opcode(0x0, 0x0, 0x0, 0x0, 0x6, catNSources, 0, 1, 1) // 64-bit reverse bits
	OpREV16x = opcode(0x0, 0x0, 0x1, 0x0, 0x6, catNSources, 0, 1, 1)
	OpREV32x = opcode(0x0, 0x0, 0x2, 0x0, 0x6, catNSources, 0, 1, 1)
	OpREVx   = opcode(0x0, 0x0, 0x3, 0x0, 0x6, catNSources, 0, 1, 1)
	OpCLZx   = opcode(0x0, 0x0, 0x4, 0x0, 0x6, catNSources, 0, 1, 1)
	OpCLSx   = opcode(0x0, 0x0, 0x5, 0x0, 0x6, catNSources, 0, 1, 1)

	// Data-Processing (1 source), feature CSSC

	OpCTZw = opcode(0x0, 0x0, 0x6, 0x0, 0x6, catNSources, 0, 1, 0)
	OpCNTw = opcode(0x0, 0x0, 0x7, 0x0, 0x6, catNSources, 0, 1, 0)
	OpABSw = opcode(0x0, 0x0, 0x8, 0x0, 0x6, catNSources, 0, 1, 0)
	OpCTZx = opcode(0x0, 0x0, 0x6, 0x0, 0x6, catNSources, 0, 1, 1)
	OpCNTx = opcode(0x0, 0x0, 0x7, 0x0, 0x6, catNSources, 0, 1, 1)
	OpABSx = opcode(0x0, 0x0, 0x8, 0x0, 0x6, catNSources, 0, 1, 1)

	// Data-Processing (1 source), feature PAuth

	OpPACIAx  = opcode(0x0, 0x0, 0x0, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACIBx  = opcode(0x0, 0x0, 0x1, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACDAx  = opcode(0x0, 0x0, 0x2, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACDBx  = opcode(0x0, 0x0, 0x3, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTIAx  = opcode(0x0, 0x0, 0x4, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTIBx  = opcode(0x0, 0x0, 0x5, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTDAx  = opcode(0x0, 0x0, 0x6, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTDBx  = opcode(0x0, 0x0, 0x7, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACIZAx = opcode(0x0, 0x0, 0x8, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACIZBx = opcode(0x0, 0x0, 0x9, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACDZAx = opcode(0x0, 0x0, 0xA, 0x1, 0x6, catNSources, 0, 1, 1)
	OpPACDZBx = opcode(0x0, 0x0, 0xB, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTIZAx = opcode(0x0, 0x0, 0xC, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTIZBx = opcode(0x0, 0x0, 0xD, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTDZAx = opcode(0x0, 0x0, 0xE, 0x1, 0x6, catNSources, 0, 1, 1)
	OpAUTDZBx = opcode(0x0, 0x0, 0xF, 0x1, 0x6, catNSources, 0, 1, 1)

	OpAND32   = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 0, 0, 0) // 32-bit bitwise AND
	OpAND64   = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 0, 0, 1) // 64-bit bitwise AND
	OpORR32   = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 1, 0, 0) // 32-bit bitwise OR
	OpORR64   = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 1, 0, 1) // 64-bit bitwise OR
	OpEOR32   = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 0, 1, 0) // 32-bit bitwise XOR
	OpEOR64   = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 0, 1, 1) // 64-bit bitwise XOR
	OpANDS32  = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 1, 1, 0) // bitwise AND, set flags
	OpANDS64  = opcode(0, 0, 0x00, 0, 0x0, catLogicReg, 1, 1, 1) // bitwise AND, set flags
	OpADD32   = opcode(0, 0, 0x00, 0, 0x0, catArithReg, 0, 0, 0) // add
	OpADDS32  = opcode(0, 0, 0x00, 0, 0x0, catArithReg, 1, 0, 0) // add, set flags
	OpSUB32   = opcode(0, 0, 0x00, 0, 0x0, catArithReg, 0, 1, 0) // subtract
	OpSUBS32  = opcode(0, 0, 0x00, 0, 0x0, catArithReg, 1, 1, 0) // subtract, set flags
	OpADDI32  = opcode(0, 0, 0x00, 0, 0x0, catArithImm, 0, 0, 0) // add with immediate
	OpADDSI32 = opcode(0, 0, 0x00, 0, 0x0, catArithImm, 1, 0, 0) // add with immediate, set flags
	OpSUBI32  = opcode(0, 0, 0x00, 0, 0x0, catArithImm, 0, 1, 0) // subtract with immediate
	OpSUBSI32 = opcode(0, 0, 0x00, 0, 0x0, catArithImm, 1, 1, 0) // subtract with immediate, set flags
	OpANDI32  = opcode(0, 0, 0x00, 0, 0x0, catLogicImm, 0, 0, 0) // bitwise AND with immediate
	OpORRI32  = opcode(0, 0, 0x00, 0, 0x0, catLogicImm, 1, 0, 0) // bitwise OR with immediate
	OpEORI32  = opcode(0, 0, 0x00, 0, 0x0, catLogicImm, 0, 1, 0) // bitwise XOR with immediate
	OpANDSI32 = opcode(0, 0, 0x00, 0, 0x0, catLogicImm, 1, 1, 0) // bitwise AND with immediate, set flags
	OpADC32   = opcode(0, 0, 0x00, 0, 0x0, catNSources, 0, 0, 0) // add with carry
	OpADCS32  = opcode(0, 0, 0x00, 0, 0x0, catNSources, 1, 0, 0) // add with carry, set flags
	OpSBC32   = opcode(0, 0, 0x00, 0, 0x0, catNSources, 0, 1, 0) // subtract with carry
	OpSBCS32  = opcode(0, 0, 0x00, 0, 0x0, catNSources, 1, 1, 0) // subtract with carry, set flags
	OpADDPT   = opcode(0, 0, 0x08, 0, 0x0, catNSources, 0, 0, 0) // add with checked pointer
	OpSUBPT   = opcode(0, 0, 0x08, 0, 0x0, catNSources, 0, 1, 0) // subtract with checked pointer
)

var dpOpcodes = map[Opcode]string{
	OpUDIVw: "UDIV",
	OpSDIVw: "SDIV",
	OpLSLVw: "LSLV",
	OpLSRVw: "LSRV",
	OpASRVw: "ASRV",
	OpRORVw: "RORV",
	OpUDIVx: "UDIV",
	OpSDIVx: "SDIV",
	OpLSLVx: "LSLV",
	OpLSRVx: "LSRV",
	OpASRVx: "ASRV",
	OpRORVx: "RORV",

	OpCRC32B:  "CRC32B",
	OpCRC32H:  "CRC32H",
	OpCRC32W:  "CRC32W",
	OpCRC32CB: "CRC32CB",
	OpCRC32CH: "CRC32CH",
	OpCRC32CW: "CRC32CW",
	OpCRC32X:  "CRC32X",
	OpCRC32CX: "CRC32CX",

	OpSMAXw: "SMAX",
	OpUMAXw: "UMAX",
	OpSMINw: "SMIN",
	OpUMINw: "UMIN",
	OpSMAXx: "SMAX",
	OpUMAXx: "UMAX",
	OpSMINx: "SMIN",
	OpUMINx: "UMIN",

	OpSUBPx:  "SUBP",
	OpSUBPSx: "SUBPS",
	OpIRGx:   "IRG",
	OpGMIx:   "GMI",

	OpPACGAx: "PACGA",

	OpRBITw:  "RBIT",
	OpREV16w: "REV16",
	OpREVw:   "REV",
	OpCLZw:   "CLZ",
	OpCLSw:   "CLS",
	OpRBITx:  "RBIT",
	OpREV16x: "REV16",
	OpREV32x: "REV32",
	OpREVx:   "REV",
	OpCLZx:   "CLZ",
	OpCLSx:   "CLS",

	OpCTZw: "CTZ",
	OpCNTw: "CNT",
	OpABSw: "ABS",
	OpCTZx: "CTZ",
	OpCNTx: "CNT",
	OpABSx: "ABS",

	OpPACIAx:  "PACIA",
	OpPACIBx:  "PACIB",
	OpPACDAx:  "PACDA",
	OpPACDBx:  "PACDB",
	OpAUTIAx:  "AUTIA",
	OpAUTIBx:  "AUTIB",
	OpAUTDAx:  "AUTDA",
	OpAUTDBx:  "AUTDB",
	OpPACIZAx: "PACIZA",
	OpPACIZBx: "PACIZB",
	OpPACDZAx: "PACDZA",
	OpPACDZBx: "PACDZB",
	OpAUTIZAx: "AUTIZA",
	OpAUTIZBx: "AUTIZB",
	OpAUTDZAx: "AUTDZA",
	OpAUTDZBx: "AUTDZB",
}

func (o Opcode) String() string {
	if name, ok := dpOpcodes[o]; ok {
		return name
	}

	return fmt.Sprintf("DPOpcode(%032b)", uint32(o))
}
