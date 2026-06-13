package aarch64

import "fmt"

type LogicShift Instruction

// ANDShift encodes an AND (shifted register) instruction
func ANDShift(rd, rn, rm Register, shift Shift, amount uint8) LogicShift {
	var i LogicShift = 0x0A000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x0, 2, 30)

	return i
}

// ORRShift encodes an ORR (shifted register) instruction
func ORRShift(rd, rn, rm Register, shift Shift, amount uint8) LogicShift {
	var i LogicShift = 0x0A000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x1, 2, 30)

	return i
}

// EORShift encodes an EOR (shifted register) instruction
func EORShift(rd, rn, rm Register, shift Shift, amount uint8) LogicShift {
	var i LogicShift = 0x0A000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x2, 2, 30)

	return i
}

// ANDSShift encodes an AND (shifted register, setting flags) instruction
func ANDSShift(rd, rn, rm Register, shift Shift, amount uint8) LogicShift {
	var i LogicShift = 0x0A000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x3, 2, 30)

	return i
}

func (i LogicShift) Rd() Register {
	return getReg(i, 5, 0)
}

func (i LogicShift) Rn() Register {
	return getReg(i, 5, 5)
}

func (i LogicShift) Rm() Register {
	return getReg(i, 5, 16)
}

func (i LogicShift) Shift() Shift {
	return get[Shift](i, 2, 22)
}

func (i LogicShift) Amount() uint8 {
	return get[uint8](i, 6, 10)
}

func (i LogicShift) Mnemonic() string {
	opc := get[uint8](i, 2, 30)

	switch opc {
	case 0x0:
		return "AND"
	case 0x1:
		return "ORR"
	case 0x2:
		return "EOR"
	case 0x3:
		return "ANDS"
	}

	return "UNALLOCATED"
}

func (i LogicShift) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)
	shift := get[Shift](i, 2, 22)
	amount := get[uint8](i, 6, 10)

	if amount != 0 {
		return fmt.Sprintf("%s %s, %s, %s, %s #%#X", mnemonic, rd, rn, rm, shift, amount)
	}

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
