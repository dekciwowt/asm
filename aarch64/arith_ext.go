package aarch64

import "fmt"

type ArithExt Instruction

// ADDExt encodes an ADD (extended register) instruction
func ADDExt(rd, rn, rm Register, ext Extension, amount uint8) ArithExt {
	var i ArithExt = 0x0B200000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Extension](i, ext, 3, 13)
	i = set[uint8](i, amount, 3, 10)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// ADDSExt encodes an ADD (extended register, setting flags) instruction
func ADDSExt(rd, rn, rm Register, ext Extension, amount uint8) ArithExt {
	var i ArithExt = 0x0B200000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Extension](i, ext, 3, 13)
	i = set[uint8](i, amount, 3, 10)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

// SUBExt encodes a SUB (extended register) instruction
func SUBExt(rd, rn, rm Register, ext Extension, amount uint8) ArithExt {
	var i ArithExt = 0x0B200000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Extension](i, ext, 3, 13)
	i = set[uint8](i, amount, 3, 10)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// SUBSExt encodes a SUB (extended register, setting flags) instruction
func SUBSExt(rd, rn, rm Register, ext Extension, amount uint8) ArithExt {
	var i ArithExt = 0x0B200000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)
	i = set[Extension](i, ext, 3, 13)
	i = set[uint8](i, amount, 3, 10)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

func (i ArithExt) Rd() Register {
	return getReg(i, 5, 0)
}

func (i ArithExt) Rn() Register {
	return getReg(i, 5, 5)
}

func (i ArithExt) Rm() Register {
	return getReg(i, 5, 16)
}

func (i ArithExt) Ext() Extension {
	return get[Extension](i, 3, 13)
}

func (i ArithExt) Amount() uint8 {
	return get[uint8](i, 3, 10)
}

func (i ArithExt) Mnemonic() string {
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

func (i ArithExt) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)
	ext := get[Extension](i, 3, 13)
	amount := get[uint8](i, 3, 10)

	if amount != 0 {
		return fmt.Sprintf("%s %s, %s, %s, %s #%d", mnemonic, rd, rn, rm, ext, amount)
	}

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
