package aarch64

import "fmt"

type DataProc1CSSC Instruction

// CTZ encodes a CTZ instruction
func CTZ(rd, rn Register) DataProc1CSSC {
	var i DataProc1CSSC = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x6, 6, 10)

	return i
}

// CNT encodes a CNT instruction
func CNT(rd, rn Register) DataProc1CSSC {
	var i DataProc1CSSC = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x7, 6, 10)

	return i
}

// ABS encodes an ABS instruction
func ABS(rd, rn Register) DataProc1CSSC {
	var i DataProc1CSSC = 0x5AC00000

	i = setSF(i, rd, rn)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x0, 5, 16)
	i = set[uint8](i, 0x8, 6, 10)

	return i
}

func (i DataProc1CSSC) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc1CSSC) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc1CSSC) Mnemonic() string {
	opcode2 := get[uint8](i, 5, 16)
	opcode := get[uint8](i, 6, 10)

	switch {
	case opcode2 == 0x0 && opcode == 0x6:
		return "CTZ"
	case opcode2 == 0x0 && opcode == 0x7:
		return "CNT"
	case opcode2 == 0x0 && opcode == 0x8:
		return "ABS"
	}

	return "UNALLOCATED"
}

func (i DataProc1CSSC) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)

	return fmt.Sprintf("%s %s, %s", mnemonic, rd, rn)
}
