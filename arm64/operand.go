package arm64

import "fmt"

type Register uint8

const (
	W0 Register = iota
	W1
	W2
	W3
	W4
	W5
	W6
	W7
	W8
	W9
	W10
	W11
	W12
	W13
	W14
	W15
	W16
	W17
	W18
	W19
	W20
	W21
	W22
	W23
	W24
	W25
	W26
	W27
	W28
	W29
	W30
	w31

	X0
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	x29
	x30
	x31
)

const (
	FP  Register = x29
	LR  Register = x30
	SP  Register = x31
	XZR Register = x31
	WZR Register = w31
)

func (r Register) String() string {
	if r < X0 {
		return fmt.Sprintf("w%d", r)
	}

	return fmt.Sprintf("x%d", r-X0)
}

// Shift represents the shift kind used in shifted-register instructions
type Shift uint8

const (
	ShiftLSL Shift = 0x0
	ShiftLSR Shift = 0x1
	ShiftASR Shift = 0x2
	ShiftROR Shift = 0x3
)

var shifts = map[Shift]string{
	ShiftLSL: "lsl",
	ShiftLSR: "lsr",
	ShiftASR: "asr",
	ShiftROR: "ror",
}

// String returns the lowercase ARM64 shift mnemonic
// Returns a formatted fallback string if the shift is not recognized
func (s Shift) String() string {
	if name, ok := shifts[s]; ok {
		return name
	}

	return fmt.Sprintf("Shift(%#03X)", uint8(s))
}

// Extension represents the extend/shift option used in extended-register
// and load/store instructions
type Extension uint8

const (
	ExtUXTB Extension = 0x0 // zero-extend byte
	ExtUXTH Extension = 0x1 // zero-extend half
	ExtUXTW Extension = 0x2 // zero-extend word
	ExtUXTX Extension = 0x3 // zero-extend dword
	ExtSXTB Extension = 0x4 // sign-extend byte
	ExtSXTH Extension = 0x5 // sign-extend half
	ExtSXTW Extension = 0x6 // sign-extend word
	ExtSXTX Extension = 0x7 // sign-extend dword

	ExtWLSL Extension = ExtUXTW // LSL for 32-bit operands
	ExtXLSL Extension = ExtUXTX // LSL for 64-bit operands
)

var exts = map[Extension]string{
	ExtUXTB: "uxtb", ExtSXTB: "sxtb",
	ExtUXTH: "uxth", ExtSXTH: "sxth",
	ExtUXTW: "uxtw", ExtSXTW: "sxtw",
	ExtUXTX: "uxtx", ExtSXTX: "sxtx",
}

// String returns the lowercase ARM64 extension mnemonic (uxtb, sxtw, etc).
// Returns a formatted fallback string if the extension is not recognized
func (e Extension) String() string {
	if name, ok := exts[e]; ok {
		return name
	}

	return fmt.Sprintf("Extension(%#03X)", uint8(e))
}

// Condition represents ARM64 condition codes used in conditional
// compare, conditional select, and branch instructions.
type Condition uint8

const (
	CondEQ Condition = 0x0 // equal (Z=1)
	CondNE Condition = 0x1 // not equal (Z=0)
	CondCS Condition = 0x2 // carry set / unsigned higher or same (C=1)
	CondHS Condition = 0x2 // alias for CS
	CondCC Condition = 0x3 // carry clear / unsigned lower (C=0)
	CondLO Condition = 0x3 // alias for CC
	CondMI Condition = 0x4 // minus / negative (N=1)
	CondPL Condition = 0x5 // plus / positive or zero (N=0)
	CondVS Condition = 0x6 // overflow (V=1)
	CondVC Condition = 0x7 // no overflow (V=0)
	CondHI Condition = 0x8 // unsigned higher (C=1 && Z=0)
	CondLS Condition = 0x9 // unsigned lower or same (!(C=1 && Z=0))
	CondGE Condition = 0xA // signed greater or equal (N=V)
	CondLT Condition = 0xB // signed less than (N!=V)
	CondGT Condition = 0xC // signed greater than (Z=0 && N=V)
	CondLE Condition = 0xD // signed less or equal (!(Z=0 && N=V))
	CondAL Condition = 0xE // always
	CondNV Condition = 0xF // always (identical to AL)
)

var conds = map[Condition]string{
	CondEQ: "eq", CondNE: "ne",
	CondCS: "cs", CondCC: "cc",
	CondMI: "mi", CondPL: "pl",
	CondVS: "vs", CondVC: "vc",
	CondHI: "hi", CondLS: "ls",
	CondGE: "ge", CondLT: "lt",
	CondGT: "gt", CondLE: "le",
	CondAL: "al", CondNV: "nv",
}

func (c Condition) String() string {
	if name, ok := conds[c]; ok {
		return name
	}

	return fmt.Sprintf("Condition(%#03X)", uint8(c))
}
