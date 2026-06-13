package aarch64

import "fmt"

type ArithImm Instruction

// ADDImm encodes an ADD (immediate) instruction
func ADDImm(rd, rn Register, imm uint16, shift bool) ArithImm {
	var i ArithImm = 0x11000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = set[uint16](i, imm, 12, 10)
	i = setBool(i, shift, 1, 22)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// ADDSImm encodes an ADD (immediate, setting flags) instruction
func ADDSImm(rd, rn Register, imm uint16, shift bool) ArithImm {
	var i ArithImm = 0x11000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = set[uint16](i, imm, 12, 10)
	i = setBool(i, shift, 1, 22)

	i = set[uint8](i, 0x0, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

// SUBImm encodes a SUB (immediate) instruction
func SUBImm(rd, rn Register, imm uint16, shift bool) ArithImm {
	var i ArithImm = 0x11000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = set[uint16](i, imm, 12, 10)
	i = setBool(i, shift, 1, 22)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x0, 1, 29)

	return i
}

// SUBSImm encodes a SUB (immediate, setting flags) instruction
func SUBSImm(rd, rn Register, imm uint16, shift bool) ArithImm {
	var i ArithImm = 0x11000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = set[uint16](i, imm, 12, 10)
	i = setBool(i, shift, 1, 22)

	i = set[uint8](i, 0x1, 1, 30)
	i = set[uint8](i, 0x1, 1, 29)

	return i
}

func (i ArithImm) Rd() Register {
	return getReg(i, 5, 0)
}

func (i ArithImm) Rn() Register {
	return getReg(i, 5, 5)
}

func (i ArithImm) Imm() uint16 {
	return get[uint16](i, 12, 10)
}

func (i ArithImm) Shift() bool {
	return getBool(i, 1, 22)
}

func (i ArithImm) Mnemonic() string {
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

func (i ArithImm) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	imm := get[uint16](i, 12, 10)
	shift := getBool(i, 1, 22)

	if shift {
		return fmt.Sprintf("%s %s, %s, #%d, LSL #12", mnemonic, rd, rn, imm)
	}

	return fmt.Sprintf("%s %s, %s, #%d", mnemonic, rd, rn, imm)
}
