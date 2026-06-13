package aarch64

import "fmt"

type LogicImm Instruction

// ANDImm encodes an AND (immediate) instruction
func ANDImm(rd, rn Register, imm uint64) LogicImm {
	var i LogicImm = 0x12000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setBitmask(i, imm, 13, 10)

	i = set[uint8](i, 0x0, 2, 30)

	return i
}

// ORRImm encodes an ORR (immediate) instruction
func ORRImm(rd, rn Register, imm uint64) LogicImm {
	var i LogicImm = 0x12000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setBitmask(i, imm, 13, 10)

	i = set[uint8](i, 0x1, 2, 30)

	return i
}

// EORImm encodes an EOR (immediate) instruction
func EORImm(rd, rn Register, imm uint64) LogicImm {
	var i LogicImm = 0x12000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setBitmask(i, imm, 13, 10)

	i = set[uint8](i, 0x2, 2, 30)

	return i
}

// ANDSImm encodes an AND (immediate, setting flags) instruction
func ANDSImm(rd, rn Register, imm uint64) LogicImm {
	var i LogicImm = 0x12000000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setBitmask(i, imm, 13, 10)

	i = set[uint8](i, 0x3, 2, 30)

	return i
}

func (i LogicImm) Rd() Register {
	return getReg(i, 5, 0)
}

func (i LogicImm) Rn() Register {
	return getReg(i, 5, 5)
}

func (i LogicImm) Imm() uint64 {
	return getBitmask(i, 13, 10)
}

func (i LogicImm) Mnemonic() string {
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

func (i LogicImm) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	imm := getBitmask(i, 13, 10)

	if imm != 0 {
		return fmt.Sprintf("%s %s, %s, #%#X", mnemonic, rd, rn, imm)
	}

	return fmt.Sprintf("%s %s, %s", mnemonic, rd, rn)
}
