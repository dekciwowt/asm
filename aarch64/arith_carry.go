package aarch64

import "fmt"

type ArithCarry Instruction

// ADC encodes an ADD with carry instruction
func ADC(rd, rn, rm Register) ArithCarry {
	var i ArithCarry = 0x03400000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// ADCS encodes an ADD with carry (setting flags) instruction
func ADCS(rd, rn, rm Register) ArithCarry {
	var i ArithCarry = 0x03400000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

// SBC encodes a SUB with carry instruction
func SBC(rd, rn, rm Register) ArithCarry {
	var i ArithCarry = 0x03400000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// SBCS encodes a SUB with carry (setting flags) instruction
func SBCS(rd, rn, rm Register) ArithCarry {
	var i ArithCarry = 0x03400000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

func (i ArithCarry) Rd() Register {
	return getReg(i, 5, 0)
}

func (i ArithCarry) Rn() Register {
	return getReg(i, 5, 5)
}

func (i ArithCarry) Rm() Register {
	return getReg(i, 5, 16)
}

func (i ArithCarry) Mnemonic() string {
	op := get[uint8](i, 1, 30)
	S := get[uint8](i, 1, 29)

	switch {
	case op == 0x0 && S == 0x0:
		return "ADC"
	case op == 0x0 && S == 0x1:
		return "ADCS"
	case op == 0x1 && S == 0x0:
		return "SBC"
	case op == 0x1 && S == 0x1:
		return "SBCS"
	}

	return "UNALLOCATED"
}

func (i ArithCarry) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
