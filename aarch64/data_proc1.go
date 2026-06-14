package aarch64

import "fmt"

type DataProc1 Instruction

// RBIT encodes a RBIT instruction
func RBIT(rd, rn Register) DataProc1 {
	var i DataProc1 = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x0, 6, 10)

	return i
}

// REV16 encodes a REV16 instruction
func REV16(rd, rn Register) DataProc1 {
	var i DataProc1 = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x1, 6, 10)

	return i
}

// REV encodes a REV instruction
func REV(rd, rn Register) DataProc1 {
	var i DataProc1 = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x2, 6, 10)

	return i
}

// CLZ encodes a CLZ instruction
func CLZ(rd, rn Register) DataProc1 {
	var i DataProc1 = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x4, 6, 10)

	return i
}

// CLS encodes a CLS instruction
func CLS(rd, rn Register) DataProc1 {
	var i DataProc1 = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x5, 6, 10)

	return i
}

func (i DataProc1) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc1) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc1) Mnemonic() string {
	opcode2 := get[uint8](i, 5, 16)
	opcode := get[uint8](i, 6, 10)

	switch {
	case opcode2 == 0x0 && opcode == 0x0:
		return "RBIT"
	case opcode2 == 0x0 && opcode == 0x1:
		return "REV16"
	case opcode2 == 0x0 && opcode == 0x2:
		return "REV"
	case opcode2 == 0x0 && opcode == 0x4:
		return "CLZ"
	case opcode2 == 0x0 && opcode == 0x5:
		return "CLS"
	}

	return "UNALLOCATED"
}

func (i DataProc1) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)

	return fmt.Sprintf("%s %s, %s", mnemonic, rd, rn)
}
