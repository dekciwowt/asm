package aarch64

import "fmt"

type ArithShift Instruction

// ADDShift encodes an ADD (shifted register) instruction
func ADDShift(rd, rn, rm Register, shift Shift, amount uint8) ArithShift {
	var i ArithShift = 0x0B000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// ADDSShift encodes an ADD (shifted register, setting flags) instruction
func ADDSShift(rd, rn, rm Register, shift Shift, amount uint8) ArithShift {
	var i ArithShift = 0x0B000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

// SUBShift encodes a SUB (shifted register) instruction
func SUBShift(rd, rn, rm Register, shift Shift, amount uint8) ArithShift {
	var i ArithShift = 0x0B000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// SUBSShift encodes a SUB (shifted register, setting flags) instruction
func SUBSShift(rd, rn, rm Register, shift Shift, amount uint8) ArithShift {
	var i ArithShift = 0x0B000000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Shift](i, shift, 2, 22)
	i = set[uint8](i, amount, 6, 10)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

func (i ArithShift) Rd() Register {
	return getReg(i, 5, 0)
}

func (i ArithShift) Rn() Register {
	return getReg(i, 5, 5)
}

func (i ArithShift) Rm() Register {
	return getReg(i, 5, 16)
}

func (i ArithShift) Shift() Shift {
	return get[Shift](i, 2, 22)
}

func (i ArithShift) Amount() uint8 {
	return get[uint8](i, 6, 10)
}

func (i ArithShift) Mnemonic() string {
	op := get[uint8](i, 1, 30)
	S := get[uint8](i, 1, 29)

	switch {
	case op == 0x0 && S == 0x0:
		return "ADD"
	case op == 0x0 && S == 0x1:
		return "ADDS"
	case op == 0x1 && S == 0x0:
		return "SUB"
	case op == 0x1 && S == 0x1:
		return "SUBS"
	}

	return "UNALLOCATED"
}

func (i ArithShift) String() string {
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
