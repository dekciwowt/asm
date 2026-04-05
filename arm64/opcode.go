package arm64

import "fmt"

// 31    30       29       28-24    23-21    20-16    15-10    9-5      4-0
// +-----+--------+--------+--------+--------+--------+--------+--------+--------+
// |  0  |  oper  |  sign  |  opt1  |  opt2  |  opt3  |  opt4  |  opt5  |  opt6  |
// +-----+--------+--------+--------+--------+--------+--------+--------+--------+
type DPOpcode uint32

const (
	dpOpcodeOpt6Size uint32 = 5
	dpOpcodeOpt5Size uint32 = 5
	dpOpcodeOpt4Size uint32 = 6
	dpOpcodeOpt3Size uint32 = 5
	dpOpcodeOpt2Size uint32 = 3
	dpOpcodeOpt1Size uint32 = 5
	dpOpcodeSignSize uint32 = 1
	dpOpcodeOperSize uint32 = 1

	dpOpcodeOpt6Mask uint32 = (1 << dpOpcodeOpt6Size) - 1
	dpOpcodeOpt5Mask uint32 = (1 << dpOpcodeOpt5Size) - 1
	dpOpcodeOpt4Mask uint32 = (1 << dpOpcodeOpt4Size) - 1
	dpOpcodeOpt3Mask uint32 = (1 << dpOpcodeOpt3Size) - 1
	dpOpcodeOpt2Mask uint32 = (1 << dpOpcodeOpt2Size) - 1
	dpOpcodeOpt1Mask uint32 = (1 << dpOpcodeOpt1Size) - 1
	dpOpcodeSignMask uint32 = (1 << dpOpcodeSignSize) - 1
	dpOpcodeOperMask uint32 = (1 << dpOpcodeOperSize) - 1

	dpOpcodeOpt6Pos uint32 = 0
	dpOpcodeOpt5Pos uint32 = 5
	dpOpcodeOpt4Pos uint32 = 10
	dpOpcodeOpt3Pos uint32 = 16
	dpOpcodeOpt2Pos uint32 = 21
	dpOpcodeOpt1Pos uint32 = 24
	dpOpcodeSignPos uint32 = 29
	dpOpcodeOperPos uint32 = 30
)

func dpOpcode(opt6, opt5, opt4, opt3, opt2, opt1, sign, oper uint8) DPOpcode {
	var opcode DPOpcode

	opcode = set(opcode, opt6, dpOpcodeOpt6Mask, dpOpcodeOpt6Pos)
	opcode = set(opcode, opt5, dpOpcodeOpt5Mask, dpOpcodeOpt5Pos)
	opcode = set(opcode, opt4, dpOpcodeOpt4Mask, dpOpcodeOpt4Pos)
	opcode = set(opcode, opt3, dpOpcodeOpt3Mask, dpOpcodeOpt3Pos)
	opcode = set(opcode, opt2, dpOpcodeOpt2Mask, dpOpcodeOpt2Pos)
	opcode = set(opcode, opt1, dpOpcodeOpt1Mask, dpOpcodeOpt1Pos)
	opcode = set(opcode, sign, dpOpcodeSignMask, dpOpcodeSignPos)
	opcode = set(opcode, oper, dpOpcodeOperMask, dpOpcodeOperPos)

	return opcode
}

const (
	dpCatLogicReg uint8 = 0x0A
	dpCatArithReg uint8 = 0x0B
	dpCatArithImm uint8 = 0x11
	dpCatLogicImm uint8 = 0x12
	dpCat2xSource uint8 = 0x1A
)

var (
	OpAND   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicReg, 0, 0) // bitwise AND
	OpORR   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicReg, 1, 0) // bitwise OR
	OpEOR   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicReg, 0, 1) // bitwise XOR
	OpANDS  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicReg, 1, 1) // bitwise AND, set flags
	OpADD   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithReg, 0, 0) // add
	OpADDS  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithReg, 1, 0) // add, set flags
	OpSUB   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithReg, 0, 1) // subtract
	OpSUBS  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithReg, 1, 1) // subtract, set flags
	OpADDI  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithImm, 0, 0) // add with immediate
	OpADDSI = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithImm, 1, 0) // add with immediate, set flags
	OpSUBI  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithImm, 0, 1) // subtract with immediate
	OpSUBSI = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatArithImm, 1, 1) // subtract with immediate, set flags
	OpANDI  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicImm, 0, 0) // bitwise AND with immediate
	OpORRI  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicImm, 1, 0) // bitwise OR with immediate
	OpEORI  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicImm, 0, 1) // bitwise XOR with immediate
	OpANDSI = dpOpcode(0, 0, 0x0, 0, 0x0, dpCatLogicImm, 1, 1) // bitwise AND with immediate, set flags
	OpADC   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCat2xSource, 0, 0) // add with carry
	OpADCS  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCat2xSource, 1, 0) // add with carry, set flags
	OpSBC   = dpOpcode(0, 0, 0x0, 0, 0x0, dpCat2xSource, 0, 1) // subtract with carry
	OpSBCS  = dpOpcode(0, 0, 0x0, 0, 0x0, dpCat2xSource, 1, 1) // subtract with carry, set flags
	OpADDPT = dpOpcode(0, 0, 0x8, 0, 0x0, dpCat2xSource, 0, 0) // add with checked pointer
	OpSUBPT = dpOpcode(0, 0, 0x8, 0, 0x0, dpCat2xSource, 0, 1) // subtract with checked pointer
	OpUDIV  = dpOpcode(0, 0, 0x2, 0, 0x6, dpCat2xSource, 0, 0) // unsigned divide
	OpSDIV  = dpOpcode(0, 0, 0x3, 0, 0x6, dpCat2xSource, 0, 0) // signed divide
)

var dpOpcodes = map[DPOpcode]string{
	OpADD:   "ADD",
	OpADDS:  "ADDS",
	OpSUB:   "SUB",
	OpSUBS:  "SUBS",
	OpAND:   "AND",
	OpORR:   "ORR",
	OpEOR:   "EOR",
	OpANDS:  "ANDS",
	OpADDI:  "ADD",
	OpADDSI: "ADDS",
	OpSUBI:  "SUB",
	OpSUBSI: "SUBS",
	OpANDI:  "AND",
	OpANDSI: "ANDS",
	OpORRI:  "ORR",
	OpEORI:  "EOR",
	OpADC:   "ADC",
	OpADCS:  "ADCS",
	OpSBC:   "SBC",
	OpSBCS:  "SBCS",
	OpADDPT: "ADDPT",
	OpSUBPT: "SUBPT",
	OpUDIV:  "UDIV",
	OpSDIV:  "SDIV",
}

func (o DPOpcode) String() string {
	if name, ok := dpOpcodes[o]; ok {
		return name
	}

	return fmt.Sprintf("DPOpcode(%032b)", uint32(o))
}
