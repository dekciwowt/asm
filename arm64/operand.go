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
	X29
	X30
)

const (
	FP  Register = X29
	LR  Register = X30
	SP  Register = X30 + 1
	XZR Register = X30 + 1
	WZR Register = W30 + 1
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
	ShiftLSL Shift = 0x0 // logical shift left
	ShiftLSR Shift = 0x1 // logical shift right
	ShiftASR Shift = 0x2 // arithmetic shift right
)

var shifts = map[Shift]string{
	ShiftLSL: "lsl",
	ShiftLSR: "lsr",
	ShiftASR: "asr",
}

// String returns the lowercase ARM64 shift mnemonic (lsl, lsr, asr)
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
	ExtUXTB: "uxtb",
	ExtUXTH: "uxth",
	ExtUXTW: "uxtw",
	ExtUXTX: "uxtx",
	ExtSXTB: "sxtb",
	ExtSXTH: "sxth",
	ExtSXTW: "sxtw",
	ExtSXTX: "sxtx",
}

// String returns the lowercase ARM64 extension mnemonic (uxtb, sxtw, etc).
// Returns a formatted fallback string if the extension is not recognized
func (e Extension) String() string {
	if name, ok := exts[e]; ok {
		return name
	}

	return fmt.Sprintf("Extension(%#03X)", uint8(e))
}
