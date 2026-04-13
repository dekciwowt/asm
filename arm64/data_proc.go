package arm64

import "fmt"

// DataProc3Source represents a Data-Processing (3 source) instruction of
// the ARM64 instruction set
//
// Layout:
//
//		31   30  29  28-24     23-21  20-16      15  14-10      9-5        4-0
//		+----+---+---+---------+------+----------+---+----------+----------+----------+
//		| sf | 0 | 0 |  11011  | opt2 |    Rm    | o |    Ra    |    Rn    |    Rd    |
//		+----+---+---+---------+------+----------+---+----------+----------+----------+
//
//	  - o encoded as first bit of the opt4 (opt4 & 0x20) and used as identity of
//	    the operation kind
type DataProc3Source Instruction

var (
	instMADDw  = identity[DataProc3Source](0, 0, 0, catDataProc3Sources, 0x0, 0x0, 0x00, 0x0, 0x0) // 32-bit multiply-add
	instMADDx  = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x0, 0x0, 0x00, 0x0, 0x0) // 64-bit multiply-add
	instMSUBw  = identity[DataProc3Source](0, 0, 0, catDataProc3Sources, 0x0, 0x0, 0x20, 0x0, 0x0) // 32-bit multiply-subtract
	instMSUBx  = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x0, 0x0, 0x20, 0x0, 0x0) // 64-bit multiply-subtract
	instSMADDL = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x1, 0x0, 0x00, 0x0, 0x0) // signed multiply-add long
	instSMSUBL = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x1, 0x0, 0x20, 0x0, 0x0) // signed multiply-subtract long
	instSMULH  = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x2, 0x0, 0x00, 0x0, 0x0) // signed multiply high
	instUMADDL = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x5, 0x0, 0x00, 0x0, 0x0) // unsigned multiply-add long
	instUMSUBL = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x5, 0x0, 0x20, 0x0, 0x0) // unsigned multiply-subtract long
	instUMULH  = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x6, 0x0, 0x00, 0x0, 0x0) // unsigned multiply high
)

var ( // feature CPA
	instMADDPT = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x3, 0x0, 0x00, 0x0, 0x0)
	instMSUBPT = identity[DataProc3Source](1, 0, 0, catDataProc3Sources, 0x3, 0x0, 0x20, 0x0, 0x0)
)

// Identity returns instruction with zeroed operands
func (i DataProc3Source) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)

	return identity[Instruction](sf, 0, 0, opt1, opt2, 0x0, opt4&0x20, 0x0, 0x0)
}

func (i DataProc3Source) WithRd(rd Register) DataProc3Source {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProc3Source) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProc3Source) WithRn(rn Register) DataProc3Source {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProc3Source) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProc3Source) WithRa(ra Register) DataProc3Source {
	return setReg(i, ra, opt4Mask>>1, opt4Pos)
}

func (i DataProc3Source) Ra() Register {
	return getReg(i, opt4Mask>>1, opt4Pos)
}

func (i DataProc3Source) WithRm(rm Register) DataProc3Source {
	return setReg(i, rm, opt3Mask, opt3Pos)
}

func (i DataProc3Source) Rm() Register {
	return getReg(i, opt3Mask, opt3Pos)
}

func (i DataProc3Source) Feature() Feature {
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	if opt2 == 0x3 {
		return FeatCPA
	}

	return FeatNone
}

var dataProc3SrcMnemonics = map[Instruction]string{
	Instruction(instMADDw):  "MADD",
	Instruction(instMADDx):  "MADD",
	Instruction(instMSUBw):  "MSUB",
	Instruction(instMSUBx):  "MSUB",
	Instruction(instSMADDL): "SMADDL",
	Instruction(instSMSUBL): "SMSUBL",
	Instruction(instSMULH):  "SMULH",
	Instruction(instUMADDL): "UMADDL",
	Instruction(instUMSUBL): "UMSUBL",
	Instruction(instUMULH):  "UMULH",
	Instruction(instMADDPT): "MADDPT",
	Instruction(instMSUBPT): "MSUBPT",
}

func (i DataProc3Source) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProc3SrcMnemonics[ident]; ok {
		return fmt.Sprintf("%s %s, %s, %s, %s", mnemonic, i.Rd(), i.Rn(), i.Ra(), i.Rm())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProc2Source represents a Data-Processing (2 source) instruction of the ARM64 instruction set
//
// Layout:
//
//	31   30  29  28-24        23-21   20-16      15-10      9-5        4-0
//	+----+---+---+------------+-------+----------+----------+----------+----------+
//	| sf | 0 | S |   11010    |  110  |    Rm    |  opcode  |    Rn    |    Rd    |
//	+----+---+---+------------+-------+----------+----------+----------+----------+
type DataProc2Source Instruction

var (
	instUDIVw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x2, 0x0, 0x0) // 32-bit unsigned divide
	instSDIVw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x3, 0x0, 0x0) // 32-bit signed divide
	instLSLVw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x8, 0x0, 0x0) // 32-bit logical shift left variable
	instLSRVw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x9, 0x0, 0x0) // 32-bit logical shift right variable
	instASRVw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0xA, 0x0, 0x0) // 32-bit arithmetic shift right variable
	instRORVw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0xB, 0x0, 0x0) // 32-bit rotate right variable
	instUDIVx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x2, 0x0, 0x0) // 64-bit unsigned divide
	instSDIVx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x3, 0x0, 0x0) // 64-bit signed divide
	instLSLVx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x8, 0x0, 0x0) // 64-bit logical shift left variable
	instLSRVx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x9, 0x0, 0x0) // 64-bit logical shift right variable
	instASRVx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0xA, 0x0, 0x0) // 64-bit arithmetic shift right variable
	instRORVx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0xB, 0x0, 0x0) // 64-bit rotate right variable
)

var ( // feature CRC32
	instCRC32B  = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x10, 0x0, 0x0) // CRC32 checksum
	instCRC32H  = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x11, 0x0, 0x0)
	instCRC32W  = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x12, 0x0, 0x0)
	instCRC32CB = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x14, 0x0, 0x0)
	instCRC32CH = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x15, 0x0, 0x0)
	instCRC32CW = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x16, 0x0, 0x0)
	instCRC32X  = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x13, 0x0, 0x0)
	instCRC32CX = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x17, 0x0, 0x0)
)

var ( // feature CSSC
	instSMAXw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x18, 0x0, 0x0) // 32-bit signed maximum
	instUMAXw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x19, 0x0, 0x0) // 32-bit unsigned maximum
	instSMINw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x1A, 0x0, 0x0) // 32-bit signed minimum
	instUMINw = identity[DataProc2Source](0, 0, 0, catDataProcNSources, 0x6, 0x0, 0x1B, 0x0, 0x0) // 32-bit unsigned minimum
	instSMAXx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x18, 0x0, 0x0) // 64-bit signed maximum
	instUMAXx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x19, 0x0, 0x0) // 64-bit unsigned maximum
	instSMINx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x1A, 0x0, 0x0) // 64-bit signed minimum
	instUMINx = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x1B, 0x0, 0x0) // 64-bit unsigned minimum
)

var ( // feature MTE
	instSUBP  = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x0, 0x0, 0x0) // subtract pointer
	instSUBPS = identity[DataProc2Source](1, 0, 1, catDataProcNSources, 0x6, 0x0, 0x0, 0x0, 0x0) // subtract pointer, setting flags
	instIRG   = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x4, 0x0, 0x0) // insert random tag
	instGMI   = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0x5, 0x0, 0x0) // tag mask insert
)

var ( // feature PAuth
	instPACGA = identity[DataProc2Source](1, 0, 0, catDataProcNSources, 0x6, 0x0, 0xC, 0x0, 0x0) // pointer auth code, using generic key
)

func (i DataProc2Source) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)

	return identity[Instruction](sf, 0, s, opt1, opt2, 0x0, opt4, 0x0, 0x0)
}

func (i DataProc2Source) WithRd(rd Register) DataProc2Source {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProc2Source) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProc2Source) WithRn(rn Register) DataProc2Source {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProc2Source) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProc2Source) WithRm(rm Register) DataProc2Source {
	return setReg(i, rm, opt3Mask, opt3Pos)
}

func (i DataProc2Source) Rm() Register {
	return getReg(i, opt3Mask, opt3Pos)
}

func (i DataProc2Source) Feature() Feature {
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)
	if opt2 == 0x6 {
		if 0x10 <= opt4 && opt4 <= 0x17 {
			return FeatCRC32
		}

		if 0x18 <= opt4 && opt4 <= 0x1B {
			return FeatCSSC
		}

		if opt4 == 0x0 || 0x4 <= opt4 && opt4 <= 0x5 {
			return FeatMTE
		}

		if opt4 == 0xC {
			return FeatPAuth
		}
	}

	if opt2 == 0x0 && opt4 == 0x8 {
		return FeatCPA
	}

	return FeatNone
}

var dataProc2SrcMnemonics = map[Instruction]string{
	Instruction(instUDIVw):   "UDIV",
	Instruction(instSDIVw):   "SDIV",
	Instruction(instLSLVw):   "LSLV",
	Instruction(instLSRVw):   "LSRV",
	Instruction(instASRVw):   "ASRV",
	Instruction(instRORVw):   "RORV",
	Instruction(instUDIVx):   "UDIV",
	Instruction(instSDIVx):   "SDIV",
	Instruction(instLSLVx):   "LSLV",
	Instruction(instLSRVx):   "LSRV",
	Instruction(instASRVx):   "ASRV",
	Instruction(instRORVx):   "RORV",
	Instruction(instCRC32B):  "CRC32B",
	Instruction(instCRC32H):  "CRC32H",
	Instruction(instCRC32W):  "CRC32W",
	Instruction(instCRC32CB): "CRC32CB",
	Instruction(instCRC32CH): "CRC32CH",
	Instruction(instCRC32CW): "CRC32CW",
	Instruction(instCRC32X):  "CRC32X",
	Instruction(instCRC32CX): "CRC32CX",
	Instruction(instSMAXw):   "SMAX",
	Instruction(instUMAXw):   "UMAX",
	Instruction(instSMINw):   "SMIN",
	Instruction(instUMINw):   "UMIN",
	Instruction(instSMAXx):   "SMAX",
	Instruction(instUMAXx):   "UMAX",
	Instruction(instSMINx):   "SMIN",
	Instruction(instUMINx):   "UMIN",
	Instruction(instSUBP):    "SUBP",
	Instruction(instSUBPS):   "SUBPS",
	Instruction(instIRG):     "IRG",
	Instruction(instGMI):     "GMI",
	Instruction(instPACGA):   "PACGA",
	Instruction(instADDPT):   "ADDPT",
	Instruction(instSUBPT):   "SUBPT",
}

func (i DataProc2Source) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProc2SrcMnemonics[ident]; ok {
		return fmt.Sprintf("%s %s, %s, %s", mnemonic, i.Rd(), i.Rn(), i.Rm())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProc1Source represents a Data-Processing (1 source) instruction of
// the ARM64 instruction set
//
// Layout:
//
//	31   30  29  28-24        23-21   20-16      15-10      9-5        4-0
//	+----+---+---+------------+-------+----------+----------+----------+----------+
//	| sf | 1 | 0 |   11010    |  110  |    op    |  opcode  |    Rn    |    Rd    |
//	+----+---+---+------------+-------+----------+----------+----------+----------+
type DataProc1Source Instruction

var (
	instRBITw  = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x0, 0x0, 0x0) // 32-bit reverse bits
	instREV16w = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x1, 0x0, 0x0)
	instREVw   = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x2, 0x0, 0x0)
	instCLZw   = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x4, 0x0, 0x0)
	instCLSw   = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x5, 0x0, 0x0)
	instRBITx  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x0, 0x0, 0x0) // 64-bit reverse bits
	instREV16x = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x1, 0x0, 0x0)
	instREV32x = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x2, 0x0, 0x0)
	instREVx   = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x3, 0x0, 0x0)
	instCLZx   = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x4, 0x0, 0x0)
	instCLSx   = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x5, 0x0, 0x0)
)

var ( // feature CSSC
	instCTZw = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x6, 0x0, 0x0)
	instCNTw = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x7, 0x0, 0x0)
	instABSw = identity[DataProc1Source](0, 1, 0, catDataProcNSources, 0x6, 0x0, 0x8, 0x0, 0x0)
	instCTZx = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x6, 0x0, 0x0)
	instCNTx = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x7, 0x0, 0x0)
	instABSx = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x0, 0x8, 0x0, 0x0)
)

var ( // feature PAuth
	instPACIA  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x00, 0x00, 0x0)
	instPACIB  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x01, 0x00, 0x0)
	instPACDA  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x02, 0x00, 0x0)
	instPACDB  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x03, 0x00, 0x0)
	instAUTIA  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x04, 0x00, 0x0)
	instAUTIB  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x05, 0x00, 0x0)
	instAUTDA  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x06, 0x00, 0x0)
	instAUTDB  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x07, 0x00, 0x0)
	instPACIZA = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x08, 0x1F, 0x0)
	instPACIZB = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x09, 0x1F, 0x0)
	instPACDZA = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x0A, 0x1F, 0x0)
	instPACDZB = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x0B, 0x1F, 0x0)
	instAUTIZA = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x0C, 0x1F, 0x0)
	instAUTIZB = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x0D, 0x1F, 0x0)
	instAUTDZA = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x0E, 0x1F, 0x0)
	instAUTDZB = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x0F, 0x1F, 0x0)
	instXPACI  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x10, 0x1F, 0x0)
	instXPACD  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x11, 0x1F, 0x0)
)

var ( // feature PAuthLR
	instPACNBIASPPC = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x20, 0x1F, 0x1E)
	instPACNBIBSPPC = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x21, 0x1F, 0x1E)
	instPACIA171615 = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x22, 0x1F, 0x1E)
	instPACIB171615 = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x23, 0x1F, 0x1E)
	instAUTIASPPCR  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x24, 0x1F, 0x1E)
	instAUTIBSPPCR  = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x25, 0x1F, 0x1E)
	instPACIASPPC   = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x28, 0x1F, 0x1E)
	instPACIBSPPC   = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x29, 0x1F, 0x1E)
	instAUTIA171615 = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x2E, 0x1F, 0x1E)
	instAUTIB171615 = identity[DataProc1Source](1, 1, 0, catDataProcNSources, 0x6, 0x1, 0x2F, 0x1F, 0x1E)
)

func (i DataProc1Source) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt3 := get[uint8](i, opt3Mask, opt3Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)

	opt5 := uint8(0x0)
	if opt3 == 0x1 && opt4 <= 0x11 {
		opt5 = get[uint8](i, opt5Mask, opt5Pos)
	}

	opt6 := uint8(0x0)
	if opt3 == 0x1 && ((0x20 <= opt4 && opt4 <= 0x29) || (0x2E <= opt4 && opt4 <= 0x2F)) {
		opt5 = get[uint8](i, opt5Mask, opt5Pos)
		opt6 = get[uint8](i, opt6Mask, opt6Pos)
	}

	return identity[Instruction](sf, op, 0, opt1, opt2, opt3, opt4, opt5, opt6)
}

func (i DataProc1Source) WithRd(rd Register) DataProc1Source {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProc1Source) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProc1Source) WithRn(rn Register) DataProc1Source {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProc1Source) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProc1Source) Feature() Feature {
	opt3 := get[uint8](i, opt3Mask, opt3Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)

	if opt3 == 0x0 && (0x6 <= opt4 && opt4 <= 0x8) {
		return FeatCSSC
	}

	if opt3 == 0x1 && opt4 <= 0x11 {
		return FeatPAuth
	}

	if opt3 == 0x1 && ((0x20 <= opt4 && opt4 <= 0x29) || (0x2E <= opt4 && opt4 <= 0x2F)) {
		return FeatPAuthLR
	}

	return FeatNone
}

var dataProc1SrcMnemonics = map[Instruction]string{
	Instruction(instRBITw):       "RBIT",
	Instruction(instREV16w):      "REV16",
	Instruction(instREVw):        "REV",
	Instruction(instCLZw):        "CLZ",
	Instruction(instCLSw):        "CLS",
	Instruction(instRBITx):       "RBIT",
	Instruction(instREV16x):      "REV16",
	Instruction(instREV32x):      "REV32",
	Instruction(instREVx):        "REV",
	Instruction(instCLZx):        "CLZ",
	Instruction(instCLSx):        "CLS",
	Instruction(instCTZw):        "CTZ",
	Instruction(instCNTw):        "CNT",
	Instruction(instABSw):        "ABS",
	Instruction(instCTZx):        "CTZ",
	Instruction(instCNTx):        "CNT",
	Instruction(instABSx):        "ABS",
	Instruction(instPACIA):       "PACIA",
	Instruction(instPACIB):       "PACIB",
	Instruction(instPACDA):       "PACDA",
	Instruction(instPACDB):       "PACDB",
	Instruction(instAUTIA):       "AUTIA",
	Instruction(instAUTIB):       "AUTIB",
	Instruction(instAUTDA):       "AUTDA",
	Instruction(instAUTDB):       "AUTDB",
	Instruction(instPACIZA):      "PACIZA",
	Instruction(instPACIZB):      "PACIZB",
	Instruction(instPACDZA):      "PACDZA",
	Instruction(instPACDZB):      "PACDZB",
	Instruction(instAUTIZA):      "AUTIZA",
	Instruction(instAUTIZB):      "AUTIZB",
	Instruction(instAUTDZA):      "AUTDZA",
	Instruction(instAUTDZB):      "AUTDZB",
	Instruction(instXPACI):       "XPACI",
	Instruction(instXPACD):       "XPACD",
	Instruction(instPACNBIASPPC): "PACNBIASPPC",
	Instruction(instPACNBIBSPPC): "PACNBIBSPPC",
	Instruction(instPACIA171615): "PACIA171615",
	Instruction(instPACIB171615): "PACIB171615",
	Instruction(instAUTIASPPCR):  "AUTIASPPCR",
	Instruction(instAUTIBSPPCR):  "AUTIBSPPCR",
	Instruction(instPACIASPPC):   "PACIASPPC",
	Instruction(instPACIBSPPC):   "PACIBSPPC",
	Instruction(instAUTIA171615): "AUTIA171615",
	Instruction(instAUTIB171615): "AUTIB171615",
}

func (i DataProc1Source) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProc1SrcMnemonics[ident]; ok {
		opt3 := get[uint8](i, opt3Mask, opt3Pos)
		opt4 := get[uint8](i, opt4Mask, opt4Pos)

		if opt3 == 0x1 && opt4 <= 0x11 {
			return fmt.Sprintf("%s %s", mnemonic, i.Rd())
		}

		if opt3 == 0x1 && ((0x20 <= opt4 && opt4 <= 0x29) || (0x2E <= opt4 && opt4 <= 0x2F)) {
			return fmt.Sprintf("%s", mnemonic)
		}

		return fmt.Sprintf("%s %s, %s", mnemonic, i.Rd(), i.Rn())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcLogicReg represents a Logical (shifted register) instruction of
// the ARM64 instruction set
//
// Layout:
//
//		31   30-29 28-24     23-22  21    20-16      15-10      9-5        4-0
//		+----+-----+---------+------+-----+----------+----------+----------+----------+
//		| sf | opc |  01010  |  sh  |  N  |    Rm    |   imm6   |    Rn    |    Rd    |
//		+----+-----+---------+------+-----+----------+----------+----------+----------+
//
//	  - N encoded as last bit of the opt2 (opt2 & 0x1) and used as identity of
//	    the operation kind
type DataProcLogicReg Instruction

var (
	instANDw  = identity[DataProcLogicReg](0, 0, 0, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instBICw  = identity[DataProcLogicReg](0, 0, 0, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instORRw  = identity[DataProcLogicReg](0, 0, 1, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instORNw  = identity[DataProcLogicReg](0, 0, 1, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instEORw  = identity[DataProcLogicReg](0, 1, 0, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instEONw  = identity[DataProcLogicReg](0, 1, 0, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instANDSw = identity[DataProcLogicReg](0, 1, 1, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instBICSw = identity[DataProcLogicReg](0, 1, 1, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instANDx  = identity[DataProcLogicReg](1, 0, 0, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instBICx  = identity[DataProcLogicReg](1, 0, 0, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instORRx  = identity[DataProcLogicReg](1, 0, 1, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instORNx  = identity[DataProcLogicReg](1, 0, 1, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instEORx  = identity[DataProcLogicReg](1, 1, 0, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instEONx  = identity[DataProcLogicReg](1, 1, 0, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
	instANDSx = identity[DataProcLogicReg](1, 1, 1, catDataProcLogicReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instBICSx = identity[DataProcLogicReg](1, 1, 1, catDataProcLogicReg, 0x1, 0x0, 0x0, 0x0, 0x0)
)

func (i DataProcLogicReg) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)

	return identity[Instruction](sf, op, s, opt1, opt2&0x1, 0x0, 0x0, 0x0, 0x0)
}

func (i DataProcLogicReg) WithRd(rd Register) DataProcLogicReg {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProcLogicReg) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProcLogicReg) WithRn(rn Register) DataProcLogicReg {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcLogicReg) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcLogicReg) WithRm(rm Register) DataProcLogicReg {
	return setReg(i, rm, opt3Mask, opt3Pos)
}

func (i DataProcLogicReg) Rm() Register {
	return getReg(i, opt3Mask, opt3Pos)
}

func (i DataProcLogicReg) WithShift(shift Shift, amount uint8) DataProcLogicReg {
	i = set(i, shift, 0x6, opt2Pos)
	i = set(i, amount, opt4Mask, opt4Pos)
	return i
}

func (i DataProcLogicReg) Shift() (Shift, uint8) {
	shift := get[Shift](i, 0x6, opt2Pos)
	amount := get[uint8](i, opt4Mask, opt4Pos)
	return shift, amount
}

func (i DataProcLogicReg) Feature() Feature {
	return FeatNone
}

var dataProcLogicRegMnemonics = map[Instruction]string{
	Instruction(instANDw):  "AND",
	Instruction(instBICw):  "BIC",
	Instruction(instORRw):  "ORR",
	Instruction(instORNw):  "ORN",
	Instruction(instEORw):  "EOR",
	Instruction(instEONw):  "EON",
	Instruction(instANDSw): "ANDS",
	Instruction(instBICSw): "BICS",
	Instruction(instANDx):  "AND",
	Instruction(instBICx):  "BIC",
	Instruction(instORRx):  "ORR",
	Instruction(instORNx):  "ORN",
	Instruction(instEORx):  "EOR",
	Instruction(instEONx):  "EON",
	Instruction(instANDSx): "ANDS",
	Instruction(instBICSx): "BICS",
}

func (i DataProcLogicReg) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcLogicRegMnemonics[ident]; ok {
		shift, amount := i.Shift()
		if shift == 0x0 && amount == 0x0 {
			return fmt.Sprintf("%s %s, %s, %s", mnemonic, i.Rd(), i.Rn(), i.Rm())
		}

		return fmt.Sprintf("%s %s, %s, %s, %s #%#X", mnemonic, i.Rd(), i.Rn(), i.Rm(), shift, amount)
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcArithReg represents an Add/substract instruction of
// the ARM64 instruction set
//
// Layout:
//
//	 N = 0, shifted register
//
//		31   30   29  28-24       23-22  21  20-16    15-10         9-5      4-0
//		+----+----+---+-----------+------+---+--------+-------------+--------+--------+
//		| sf | op | S |   01011   |  sh  | N |   Rm   |    imm6     |   Rn   |   Rd   |
//		+----+----+---+-----------+------+---+--------+-------------+--------+--------+
//
//	 N = 1, extended register
//
//		31   30   29  28-24       23-22  21  20-16    15-13  12-10  9-5      4-0
//		+----+----+---+-----------+------+---+--------+------+------+--------+--------+
//		| sf | op | S |   01011   |  00  | N |   Rm   | ext  | imm3 |   Rn   |   Rd   |
//		+----+----+---+-----------+------+---+--------+------+------+--------+--------+
type DataProcArithReg Instruction

var (
	instADDw  = identity[DataProcArithReg](0, 0, 0, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instADDSw = identity[DataProcArithReg](0, 0, 1, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSUBw  = identity[DataProcArithReg](0, 1, 0, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSUBSw = identity[DataProcArithReg](0, 1, 1, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instADDx  = identity[DataProcArithReg](1, 0, 0, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instADDSx = identity[DataProcArithReg](1, 0, 1, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSUBx  = identity[DataProcArithReg](1, 1, 0, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSUBSx = identity[DataProcArithReg](1, 1, 1, catDataProcArithReg, 0x0, 0x0, 0x0, 0x0, 0x0)
)

func (i DataProcArithReg) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)

	return identity[Instruction](sf, op, s, opt1, 0x0, 0x0, 0x0, 0x0, 0x0)
}

func (i DataProcArithReg) WithRd(rd Register) DataProcArithReg {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProcArithReg) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProcArithReg) WithRn(rn Register) DataProcArithReg {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcArithReg) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcArithReg) WithRm(rm Register) DataProcArithReg {
	return setReg(i, rm, opt3Mask, opt3Pos)
}

func (i DataProcArithReg) Rm() Register {
	return getReg(i, opt3Mask, opt3Pos)
}

func (i DataProcArithReg) WithShift(shift Shift, amount uint8) DataProcArithReg {
	i = set(i, uint8(shift<<1), opt2Mask, opt2Pos)
	i = set(i, amount, opt4Mask, opt4Pos)
	return i
}

func (i DataProcArithReg) Shift() (Shift, uint8) {
	shift := get[Shift](i, 0x6, opt2Pos)
	amount := get[uint8](i, opt4Mask, opt4Pos)
	return shift, amount
}

func (i DataProcArithReg) WithExtension(ext Extension, amount uint8) DataProcArithReg {
	i = set(i, uint8(0x1), opt2Mask, opt2Pos)
	i = set(i, amount, 0x7, 10)
	i = set(i, ext, 0x7, 13)
	return i
}

func (i DataProcArithReg) Extension() (Extension, uint8) {
	amount := get[uint8](i, 0x7, 10)
	ext := get[Extension](i, 0x7, 13)
	return ext, amount
}

func (i DataProcArithReg) Feature() Feature {
	return FeatNone
}

var dataProcArithRegMnemonics = map[Instruction]string{
	Instruction(instADDw):  "ADD",
	Instruction(instADDSw): "ADDS",
	Instruction(instSUBw):  "SUB",
	Instruction(instSUBSw): "SUBS",
	Instruction(instADDx):  "ADD",
	Instruction(instADDSx): "ADDS",
	Instruction(instSUBx):  "SUB",
	Instruction(instSUBSx): "SUBS",
}

func (i DataProcArithReg) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcArithRegMnemonics[ident]; ok {
		n := get[uint8](i, 0x1, 21)

		if shift, amount := i.Shift(); n == 0 && amount != 0x0 {
			return fmt.Sprintf("%s %s, %s, %s, %s #%#X", mnemonic, i.Rd(), i.Rn(), i.Rm(), shift, amount)
		}

		if ext, amount := i.Extension(); n == 1 && amount != 0x0 {
			return fmt.Sprintf("%s %s, %s, %s, %s #%#X", mnemonic, i.Rd(), i.Rn(), i.Rm(), ext, amount)
		}

		return fmt.Sprintf("%s %s, %s, %s", mnemonic, i.Rd(), i.Rn(), i.Rm())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcArithWithCarry represents an Add/substract (with carry) instruction of
// the ARM64 instruction set
//
// Layout:
//
//	31   30   29   28-24      23-21   20-16    15-10            9-5      4-0
//	+----+----+---+-----------+-------+--------+----------------+--------+--------+
//	| sf | op | S |   11010   |  000  |   Rm   |    000000      |   Rn   |   Rd   |
//	+----+----+---+-----------+-------+--------+----------------+--------+--------+
type DataProcArithWithCarry Instruction

var (
	instADCw  = identity[DataProcArithWithCarry](0, 0, 0, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instADCSw = identity[DataProcArithWithCarry](0, 0, 1, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSBCw  = identity[DataProcArithWithCarry](0, 1, 0, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSBCSw = identity[DataProcArithWithCarry](0, 1, 1, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instADCx  = identity[DataProcArithWithCarry](1, 0, 0, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instADCSx = identity[DataProcArithWithCarry](1, 0, 1, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSBCx  = identity[DataProcArithWithCarry](1, 1, 0, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
	instSBCSx = identity[DataProcArithWithCarry](1, 1, 1, catDataProcNSources, 0x0, 0x0, 0x0, 0x0, 0x0)
)

func (i DataProcArithWithCarry) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)

	return identity[Instruction](sf, op, s, opt1, 0x0, 0x0, 0x0, 0x0, 0x0)
}

func (i DataProcArithWithCarry) WithRd(rd Register) DataProcArithWithCarry {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProcArithWithCarry) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProcArithWithCarry) WithRn(rn Register) DataProcArithWithCarry {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcArithWithCarry) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcArithWithCarry) WithRm(rm Register) DataProcArithWithCarry {
	return setReg(i, rm, opt3Mask, opt3Pos)
}

func (i DataProcArithWithCarry) Rm() Register {
	return getReg(i, opt3Mask, opt3Pos)
}

func (i DataProcArithWithCarry) Feature() Feature {
	return FeatNone
}

var dataProcArithWithCarryMnemonics = map[Instruction]string{
	Instruction(instADCw):  "ADC",
	Instruction(instADCSw): "ADCS",
	Instruction(instSBCw):  "SBC",
	Instruction(instSBCSw): "SBCS",
	Instruction(instADCx):  "ADC",
	Instruction(instADCSx): "ADCS",
	Instruction(instSBCx):  "SBC",
	Instruction(instSBCSx): "SBCS",
}

func (i DataProcArithWithCarry) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcArithWithCarryMnemonics[ident]; ok {
		return fmt.Sprintf("%s %s, %s, %s", mnemonic, i.Rd(), i.Rn(), i.Rm())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcArithCkPtr represents an Add/substract (checked pointer) instruction of
// the ARM64 instruction set
//
// Layout:
//
//	31  30   29  28-24       23-21    20-16    15-13   12-10    9-5      4-0
//	+---+----+---+-----------+--------+--------+-------+--------+--------+--------+
//	| 1 | op | S |   11010   |  000   |   Rm   |  001  |  imm3  |   Rn   |   Rd   |
//	+---+----+---+-----------+--------+--------+-------+--------+--------+--------+
type DataProcArithCkPtr Instruction

var ( // feature CPA
	instADDPT = identity[DataProcArithCkPtr](1, 0, 0, catDataProcNSources, 0x0, 0x0, 0x8, 0x0, 0x0)
	instSUBPT = identity[DataProcArithCkPtr](1, 1, 0, catDataProcNSources, 0x0, 0x0, 0x8, 0x0, 0x0)
)

func (i DataProcArithCkPtr) Identity() Instruction {
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)

	return identity[Instruction](1, op, s, opt1, 0x0, 0x0, 0x8, 0x0, 0x0)
}

func (i DataProcArithCkPtr) WithRd(rd Register) DataProcArithCkPtr {
	return setReg(i, rd, opt6Mask, opt6Pos)
}

func (i DataProcArithCkPtr) Rd() Register {
	return getReg(i, opt6Mask, opt6Pos)
}

func (i DataProcArithCkPtr) WithRn(rn Register) DataProcArithCkPtr {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcArithCkPtr) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcArithCkPtr) WithRm(rm Register) DataProcArithCkPtr {
	return setReg(i, rm, opt3Mask, opt3Pos)
}

func (i DataProcArithCkPtr) Rm() Register {
	return getReg(i, opt3Mask, opt3Pos)
}

func (i DataProcArithCkPtr) WithShift(amount uint8) DataProcArithCkPtr {
	i = set(i, amount, 0x7, 10)
	return i
}

func (i DataProcArithCkPtr) Shift() uint8 {
	return get[uint8](i, 0x7, 10)
}

func (i DataProcArithCkPtr) Feature() Feature {
	return FeatCPA
}

var dataProcArithCkPtrMnemonics = map[Instruction]string{
	Instruction(instADDPT): "ADDPT",
	Instruction(instSUBPT): "SUBPT",
}

func (i DataProcArithCkPtr) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcArithCkPtrMnemonics[ident]; ok {
		if amount := i.Shift(); amount != 0x0 {
			return fmt.Sprintf("%s %s, %s, %s, lsl #%#X", mnemonic, i.Rd(), i.Rn(), i.Rm(), amount)
		}

		return fmt.Sprintf("%s %s, %s, %s", mnemonic, i.Rd(), i.Rn(), i.Rm())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcRotate represents a rotate right into flags instruction of
// the ARM64 instruction set
//
// Layout:
//
//	31   30   29  28-24        23-21   20-15      14-10     9-5      4   3-0
//	+----+----+---+------------+-------+----------+---------+--------+---+--------+
//	| sf | op | S |   11010    |  000  |   imm6   |  00001  |   Rn   | 0 |  mask  |
//	+----+----+---+------------+-------+----------+---------+--------+---+--------+
type DataProcRotate Instruction

var ( // feature FlagM
	instRMIF = identity[DataProcRotate](1, 0, 1, catDataProcNSources, 0x0, 0x0, 0x1, 0x0, 0x0)
)

func (i DataProcRotate) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)

	return identity[Instruction](sf, op, s, opt1, 0x0, 0x0, 0x1, 0x0, 0x0)
}

func (i DataProcRotate) WithMask(mask uint8) DataProcRotate {
	return set(i, mask, 0xF, opt6Pos)
}

func (i DataProcRotate) Mask() uint8 {
	return get[uint8](i, 0xF, opt6Pos)
}

func (i DataProcRotate) WithRn(rn Register) DataProcRotate {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcRotate) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcRotate) WithShift(shift uint8) DataProcRotate {
	return set(i, shift, 0x1F, 15)
}

func (i DataProcRotate) Shift() uint8 {
	return get[uint8](i, 0x1F, 15)
}

func (i DataProcRotate) Feature() Feature {
	return FeatFlagM
}

var dataProcRotateMnemonics = map[Instruction]string{
	Instruction(instRMIF): "RMIF",
}

func (i DataProcRotate) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcRotateMnemonics[ident]; ok {
		shift := i.Shift()
		mask := i.Mask()
		return fmt.Sprintf("%s %s, #%#X, #%#X", mnemonic, i.Rn(), shift, mask)
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcRotate represents a rotate right into flags instruction of
// the ARM64 instruction set
//
// Layout:
//
//	31   30   29  28-24      23-21   20-15    14   13-10    9-5      4   3-0
//	+----+----+---+----------+-------+--------+----+--------+--------+---+--------+
//	| sf | op | S |  11010   |  000  | opcode | sz |  0010  |   Rn   | 0 |  mask  |
//	+----+----+---+----------+-------+--------+----+--------+--------+---+--------+
type DataProcEvaluate Instruction

var ( // feature FlagM
	instSETF8  = identity[DataProcEvaluate](0, 0, 1, catDataProcNSources, 0x0, 0x0, 0x02, 0x0, 0xD)
	instSETF16 = identity[DataProcEvaluate](0, 0, 1, catDataProcNSources, 0x0, 0x0, 0x12, 0x0, 0xD)
)

func (i DataProcEvaluate) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)
	opt6 := get[uint8](i, opt6Mask, opt6Pos)

	return identity[Instruction](sf, op, s, opt1, 0x0, 0x0, opt4, 0x0, opt6)
}

func (i DataProcEvaluate) WithRn(rn Register) DataProcEvaluate {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcEvaluate) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcEvaluate) Feature() Feature {
	return FeatFlagM
}

var dataProcEvaluateMnemonics = map[Instruction]string{
	Instruction(instSETF8):  "SETF8",
	Instruction(instSETF16): "SETF16",
}

func (i DataProcEvaluate) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcEvaluateMnemonics[ident]; ok {
		return fmt.Sprintf("%s %s", mnemonic, i.Rn())
	}

	return fmt.Sprintf("%032b", i)
}

// DataProcCondCompReg represents a conditional compare (register) instruction
// of the ARM64 instruction set
//
// Layout:
//
//	31   30   29  28-24     23-21   20-16    15-12  11  10   9-5    4    3-0
//	+----+----+---+---------+-------+--------+------+---+----+------+----+--------+
//	| sf | op | S |  11010  |  010  | opcode | cond | 0 | o2 |  Rn  | o3 |  mask  |
//	+----+----+---+---------+-------+--------+----------+----+------+----+--------+
type DataProcCondCompReg Instruction

var ( // feature FlagM
	instCCMNw = identity[DataProcCondCompReg](0, 0, 1, catDataProcNSources, 0x2, 0x0, 0x0, 0x0, 0x0)
	instCCMPw = identity[DataProcCondCompReg](0, 1, 1, catDataProcNSources, 0x2, 0x0, 0x0, 0x0, 0x0)
	instCCMNx = identity[DataProcCondCompReg](1, 0, 1, catDataProcNSources, 0x2, 0x0, 0x0, 0x0, 0x0)
	instCCMPx = identity[DataProcCondCompReg](1, 1, 1, catDataProcNSources, 0x2, 0x0, 0x0, 0x0, 0x0)
)

func (i DataProcCondCompReg) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)

	return identity[Instruction](sf, op, s, opt1, 0x2, 0x0, 0x0, 0x0, 0x0)
}

func (i DataProcCondCompReg) WithRn(rn Register) DataProcCondCompReg {
	return setReg(i, rn, opt5Mask, opt5Pos)
}

func (i DataProcCondCompReg) Rn() Register {
	return getReg(i, opt5Mask, opt5Pos)
}

func (i DataProcCondCompReg) WithCondition(mask uint8, cond Condition) DataProcCondCompReg {
	i = set(i, cond, 0xF, 12)
	i = set(i, mask, 0xF, 0)
	return i
}

func (i DataProcCondCompReg) Condition() (uint8, Condition) {
	cond := get[Condition](i, 0xF, 12)
	mask := get[uint8](i, 0xF, 0)
	return mask, cond
}

func (i DataProcCondCompReg) Feature() Feature {
	return FeatNone
}

var dataProcCondCompRegMnemonics = map[Instruction]string{
	Instruction(instCCMNw): "CCMN",
	Instruction(instCCMPw): "CCMP",
	Instruction(instCCMNx): "CCMN",
	Instruction(instCCMPx): "CCMP",
}

func (i DataProcCondCompReg) String() string {
	ident := i.Identity()
	if mnemonic, ok := dataProcCondCompRegMnemonics[ident]; ok {
		mask, cond := i.Condition()
		return fmt.Sprintf("%s %s, #%#X, %s", mnemonic, i.Rn(), mask, cond)
	}

	return fmt.Sprintf("%032b", i)
}

type DataProcCondCompImm Instruction

type DataProcCondSelect Instruction
