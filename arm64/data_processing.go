package arm64

import "fmt"

// Data-Processing (3 source)

var (
	instMADDw  = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x0, catDataProc3Sources, 0, 0, 0) // 32-bit multiply-add
	instMADDx  = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x0, catDataProc3Sources, 0, 0, 1) // 64-bit multiply-add
	instMSUBw  = identity[DataProc3Source](0x0, 0x0, 0x20, 0x0, 0x0, catDataProc3Sources, 0, 0, 0) // 32-bit multiply-subtract
	instMSUBx  = identity[DataProc3Source](0x0, 0x0, 0x20, 0x0, 0x0, catDataProc3Sources, 0, 0, 1) // 64-bit multiply-subtract
	instSMADDL = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x1, catDataProc3Sources, 0, 0, 1) // signed multiply-add long
	instSMSUBL = identity[DataProc3Source](0x0, 0x0, 0x20, 0x0, 0x1, catDataProc3Sources, 0, 0, 1) // signed multiply-subtract long
	instSMULH  = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x2, catDataProc3Sources, 0, 0, 1) // signed multiply high
	instUMADDL = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x5, catDataProc3Sources, 0, 0, 1) // unsigned multiply-add long
	instUMSUBL = identity[DataProc3Source](0x0, 0x0, 0x20, 0x0, 0x5, catDataProc3Sources, 0, 0, 1) // unsigned multiply-subtract long
	instUMULH  = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x6, catDataProc3Sources, 0, 0, 1) // unsigned multiply high
)

var ( // feature CPA
	instMADDPT = identity[DataProc3Source](0x0, 0x0, 0x00, 0x0, 0x3, catDataProc3Sources, 0, 0, 1)
	instMSUBPT = identity[DataProc3Source](0x0, 0x0, 0x20, 0x0, 0x3, catDataProc3Sources, 0, 0, 1)
)

// DataProc3Source represents a Data-Processing (3 source) instruction of the ARM64 instruction set
//
// Layout:
//
//	31   30  29  28-24     23-21  20-16      15  14-10      9-5        4-0
//	+----+---+---+---------+------+----------+---+----------+----------+----------+
//	| sf | 0 | 0 |  11011  | opt2 |    Rm    | o |    Ra    |    Rn    |    Rd    |
//	+----+---+---+---------+------+----------+---+----------+----------+----------+
//
// * o encoded as first bit of the opt4 (opt4 & 0x20) and used as identity of the operation kind
type DataProc3Source Instruction

func (i DataProc3Source) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)

	return identity[Instruction](0x0, 0x0, opt4&0x20, 0x0, opt2, opt1, 0, 0, sf)
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

	return fmt.Sprintf("%032b", ident)
}

// Data-Processing (2 source)

var (
	instUDIVw = identity[DataProc2Source](0x0, 0x0, 0x2, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit unsigned divide
	instSDIVw = identity[DataProc2Source](0x0, 0x0, 0x3, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit signed divide
	instLSLVw = identity[DataProc2Source](0x0, 0x0, 0x8, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit logical shift left variable
	instLSRVw = identity[DataProc2Source](0x0, 0x0, 0x9, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit logical shift right variable
	instASRVw = identity[DataProc2Source](0x0, 0x0, 0xA, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit arithmetic shift right variable
	instRORVw = identity[DataProc2Source](0x0, 0x0, 0xB, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit rotate right variable
	instUDIVx = identity[DataProc2Source](0x0, 0x0, 0x2, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit unsigned divide
	instSDIVx = identity[DataProc2Source](0x0, 0x0, 0x3, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit signed divide
	instLSLVx = identity[DataProc2Source](0x0, 0x0, 0x8, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit logical shift left variable
	instLSRVx = identity[DataProc2Source](0x0, 0x0, 0x9, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit logical shift right variable
	instASRVx = identity[DataProc2Source](0x0, 0x0, 0xA, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit arithmetic shift right variable
	instRORVx = identity[DataProc2Source](0x0, 0x0, 0xB, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit rotate right variable
)

var ( // feature CRC32
	instCRC32B  = identity[DataProc2Source](0x0, 0x0, 0x10, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // CRC32 checksum
	instCRC32H  = identity[DataProc2Source](0x0, 0x0, 0x11, 0x0, 0x6, catDataProcNSources, 0, 0, 0)
	instCRC32W  = identity[DataProc2Source](0x0, 0x0, 0x12, 0x0, 0x6, catDataProcNSources, 0, 0, 0)
	instCRC32CB = identity[DataProc2Source](0x0, 0x0, 0x14, 0x0, 0x6, catDataProcNSources, 0, 0, 0)
	instCRC32CH = identity[DataProc2Source](0x0, 0x0, 0x15, 0x0, 0x6, catDataProcNSources, 0, 0, 0)
	instCRC32CW = identity[DataProc2Source](0x0, 0x0, 0x16, 0x0, 0x6, catDataProcNSources, 0, 0, 0)
	instCRC32X  = identity[DataProc2Source](0x0, 0x0, 0x13, 0x0, 0x6, catDataProcNSources, 0, 0, 1)
	instCRC32CX = identity[DataProc2Source](0x0, 0x0, 0x17, 0x0, 0x6, catDataProcNSources, 0, 0, 1)
)

var ( // feature CSSC
	instSMAXw = identity[DataProc2Source](0x0, 0x0, 0x18, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit signed maximum
	instUMAXw = identity[DataProc2Source](0x0, 0x0, 0x19, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit unsigned maximum
	instSMINw = identity[DataProc2Source](0x0, 0x0, 0x1A, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit signed minimum
	instUMINw = identity[DataProc2Source](0x0, 0x0, 0x1B, 0x0, 0x6, catDataProcNSources, 0, 0, 0) // 32-bit unsigned minimum
	instSMAXx = identity[DataProc2Source](0x0, 0x0, 0x18, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit signed maximum
	instUMAXx = identity[DataProc2Source](0x0, 0x0, 0x19, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit unsigned maximum
	instSMINx = identity[DataProc2Source](0x0, 0x0, 0x1A, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit signed minimum
	instUMINx = identity[DataProc2Source](0x0, 0x0, 0x1B, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // 64-bit unsigned minimum
)

var ( // feature MTE
	instSUBP  = identity[DataProc2Source](0x0, 0x0, 0x0, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // subtract pointer
	instSUBPS = identity[DataProc2Source](0x0, 0x0, 0x0, 0x0, 0x6, catDataProcNSources, 1, 0, 1) // subtract pointer, setting flags
	instIRG   = identity[DataProc2Source](0x0, 0x0, 0x4, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // insert random tag
	instGMI   = identity[DataProc2Source](0x0, 0x0, 0x5, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // tag mask insert
)

var ( // feature PAuth
	instPACGA = identity[DataProc2Source](0x0, 0x0, 0xC, 0x0, 0x6, catDataProcNSources, 0, 0, 1) // pointer auth code, using generic key
)

var ( // feature CPA
	instADDPT = identity[DataProc2Source](0x0, 0x0, 0x8, 0x0, 0x0, catDataProcNSources, 0, 0, 1) // add checked pointer
	instSUBPT = identity[DataProc2Source](0x0, 0x0, 0x8, 0x0, 0x0, catDataProcNSources, 1, 0, 1) // subtract checked pointer
)

// DataProc2Source represents a Data-Processing (2 source) instruction of the ARM64 instruction set
//
// Layout:
//
//	31   30  29  28-24        23-21   20-16      15-10      9-5        4-0
//	+----+---+---+------------+-------+----------+----------+----------+----------+
//	| sf | 0 | S |   11010    |  110  |    Rm    |  opcode  |    Rn    |    Rd    |
//	+----+---+---+------------+-------+----------+----------+----------+----------+
type DataProc2Source Instruction

func (i DataProc2Source) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	s := get[uint8](i, sMask, sPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)

	return identity[Instruction](0x0, 0x0, opt4, 0x0, opt2, opt1, s, 0, sf)
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

	return fmt.Sprintf("%032b", ident)
}

// Data-Processing (1 source)

var (
	instRBITw  = identity[DataProc1Source](0x0, 0x0, 0x0, 0x0, 0x6, catDataProcNSources, 0, 1, 0) // 32-bit reverse bits
	instREV16w = identity[DataProc1Source](0x0, 0x0, 0x1, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instREVw   = identity[DataProc1Source](0x0, 0x0, 0x2, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instCLZw   = identity[DataProc1Source](0x0, 0x0, 0x4, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instCLSw   = identity[DataProc1Source](0x0, 0x0, 0x5, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instRBITx  = identity[DataProc1Source](0x0, 0x0, 0x0, 0x0, 0x6, catDataProcNSources, 0, 1, 1) // 64-bit reverse bits
	instREV16x = identity[DataProc1Source](0x0, 0x0, 0x1, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
	instREV32x = identity[DataProc1Source](0x0, 0x0, 0x2, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
	instREVx   = identity[DataProc1Source](0x0, 0x0, 0x3, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
	instCLZx   = identity[DataProc1Source](0x0, 0x0, 0x4, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
	instCLSx   = identity[DataProc1Source](0x0, 0x0, 0x5, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
)

var ( // feature CSSC
	instCTZw = identity[DataProc1Source](0x0, 0x0, 0x6, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instCNTw = identity[DataProc1Source](0x0, 0x0, 0x7, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instABSw = identity[DataProc1Source](0x0, 0x0, 0x8, 0x0, 0x6, catDataProcNSources, 0, 1, 0)
	instCTZx = identity[DataProc1Source](0x0, 0x0, 0x6, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
	instCNTx = identity[DataProc1Source](0x0, 0x0, 0x7, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
	instABSx = identity[DataProc1Source](0x0, 0x0, 0x8, 0x0, 0x6, catDataProcNSources, 0, 1, 1)
)

var ( // feature PAuth
	instPACIA  = identity[DataProc1Source](0x0, 0x00, 0x00, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIB  = identity[DataProc1Source](0x0, 0x00, 0x01, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACDA  = identity[DataProc1Source](0x0, 0x00, 0x02, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACDB  = identity[DataProc1Source](0x0, 0x00, 0x03, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIA  = identity[DataProc1Source](0x0, 0x00, 0x04, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIB  = identity[DataProc1Source](0x0, 0x00, 0x05, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTDA  = identity[DataProc1Source](0x0, 0x00, 0x06, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTDB  = identity[DataProc1Source](0x0, 0x00, 0x07, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIZA = identity[DataProc1Source](0x0, 0x1F, 0x08, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIZB = identity[DataProc1Source](0x0, 0x1F, 0x09, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACDZA = identity[DataProc1Source](0x0, 0x1F, 0x0A, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACDZB = identity[DataProc1Source](0x0, 0x1F, 0x0B, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIZA = identity[DataProc1Source](0x0, 0x1F, 0x0C, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIZB = identity[DataProc1Source](0x0, 0x1F, 0x0D, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTDZA = identity[DataProc1Source](0x0, 0x1F, 0x0E, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTDZB = identity[DataProc1Source](0x0, 0x1F, 0x0F, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instXPACI  = identity[DataProc1Source](0x0, 0x1F, 0x10, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instXPACD  = identity[DataProc1Source](0x0, 0x1F, 0x11, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
)

var ( // feature PAuthLR
	instPACNBIASPPC = identity[DataProc1Source](0x1E, 0x1F, 0x20, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACNBIBSPPC = identity[DataProc1Source](0x1E, 0x1F, 0x21, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIA171615 = identity[DataProc1Source](0x1E, 0x1F, 0x22, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIB171615 = identity[DataProc1Source](0x1E, 0x1F, 0x23, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIASPPCR  = identity[DataProc1Source](0x1E, 0x1F, 0x24, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIBSPPCR  = identity[DataProc1Source](0x1E, 0x1F, 0x25, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIASPPC   = identity[DataProc1Source](0x1E, 0x1F, 0x28, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instPACIBSPPC   = identity[DataProc1Source](0x1E, 0x1F, 0x29, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIA171615 = identity[DataProc1Source](0x1E, 0x1F, 0x2E, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
	instAUTIB171615 = identity[DataProc1Source](0x1E, 0x1F, 0x2F, 0x1, 0x6, catDataProcNSources, 0, 1, 1)
)

// DataProc1Source represents a Data-Processing (1 source) instruction of the ARM64 instruction set
//
// Layout:
//
//	31   30  29  28-24        23-21   20-16      15-10      9-5        4-0
//	+----+---+---+------------+-------+----------+----------+----------+----------+
//	| sf | 1 | 0 |   11010    |  110  |    op    |  opcode  |    Rn    |    Rd    |
//	+----+---+---+------------+-------+----------+----------+----------+----------+
type DataProc1Source Instruction

func (i DataProc1Source) Identity() Instruction {
	sf := get[uint8](i, sfMask, sfPos)
	op := get[uint8](i, opMask, opPos)
	opt1 := get[uint8](i, opt1Mask, opt1Pos)
	opt2 := get[uint8](i, opt2Mask, opt2Pos)
	opt3 := get[uint8](i, opt3Mask, opt3Pos)
	opt4 := get[uint8](i, opt4Mask, opt4Pos)
	opt5 := uint8(0x0)

	if opt3 == 0x1 && 0x8 <= opt4 && opt4 <= 0xF {
		opt5 = get[uint8](i, opt5Mask, opt5Pos)
	}

	return identity[Instruction](0x0, opt5, opt4, opt3, opt2, opt1, 0, op, sf)
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
		return fmt.Sprintf("%s %s, %s", mnemonic, i.Rd(), i.Rn())
	}

	return fmt.Sprintf("%032b", ident)
}

//
// Data-Processing (shifted register)
//
//	31   30   29  28-24        23-22  21  20-16    15-10        9-5      4-0
//	+----+----+---+------------+------+---+--------+------------+--------+--------+
//	| sf | op | S |   opcode   |  sh  | 0 |   Rm   |    imm6    |   Rn   |   Rd   |
//	+----+----+---+------------+------+---+--------+------------+--------+--------+
//
// Data-Processing (extended register)
//
//	31   30   29  28-24        23-22  21  20-16    15-13 12-10  9-5      4-0
//	+----+----+---+------------+------+---+--------+-----+------+--------+--------+
//	| sf | op | S |   opcode   |  00  | 1 |   Rm   | opt | imm3 |   Rn   |   Rd   |
//	+----+----+---+------------+------+---+--------+-----+------+--------+--------+
//
// Data-Processing (immediate)
//
//	31   30   29  28-24        23  22     21-10                 9-5      4-0
//	+----+----+---+------------+---+------+---------------------+--------+--------+
//	| sf | op | S |   opcode   | 0 |  sh  |        imm12        |   Rn   |   Rd   |
//	+----+----+---+------------+---+------+---------------------+--------+--------+
//
// Data-Processing (bitmask immediate)
//
//	31   30   29  28-24        23  22     21-16      15-10      9-5      4-0
//	+----+----+---+------------+---+------+----------+----------+--------+--------+
//	| sf | op | S |   opcode   | 0 |  N   |   immr   |   imms   |   Rn   |   Rd   |
//	+----+----+---+------------+---+------+----------+----------+--------+--------+
//
// Data-Processing (checked pointer)
//
//	31   30   29  28-24        23-22  21  20-16    15-13 12-10  9-5      4-0
//	+----+----+---+------------+------+---+--------+-----+------+--------+--------+
//	| sf | op | S |   opcode   |  00  | 0 |   Rm   | 001 | imm3 |   Rn   |   Rd   |
//	+----+----+---+------------+------+---+--------+-----+------+--------+--------+
