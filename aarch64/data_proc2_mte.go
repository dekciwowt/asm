package aarch64

import "fmt"

type DataProc2MTE Instruction

// SUBP encodes a SUBP instruction
func SUBP(rd, rn, rm Register) DataProc2MTE {
	var i DataProc2MTE = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x0, 1, 29)
	i = set[uint8](i, 0x0, 6, 10)

	return i
}

// IRG encodes an IRG instruction
func IRG(rd, rn, rm Register) DataProc2MTE {
	var i DataProc2MTE = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x0, 1, 29)
	i = set[uint8](i, 0x4, 6, 10)

	return i
}

// GMI encodes a GMI instruction
func GMI(rd, rn, rm Register) DataProc2MTE {
	var i DataProc2MTE = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x0, 1, 29)
	i = set[uint8](i, 0x5, 6, 10)

	return i
}

// SUBPS encodes a SUBPS instruction
func SUBPS(rd, rn, rm Register) DataProc2MTE {
	var i DataProc2MTE = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 1, 29)
	i = set[uint8](i, 0x0, 6, 10)

	return i
}

func (i DataProc2MTE) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc2MTE) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc2MTE) Rm() Register {
	return getReg(i, 5, 16)
}

func (i DataProc2MTE) Mnemonic() string {
	sf := get[uint8](i, 1, 31)
	S := get[uint8](i, 1, 29)
	opcode := get[uint8](i, 6, 10)

	switch {
	case sf == 0x1 && S == 0x0 && opcode == 0x0:
		return "SUBP"
	case sf == 0x1 && S == 0x0 && opcode == 0x4:
		return "IRG"
	case sf == 0x1 && S == 0x0 && opcode == 0x5:
		return "GMI"
	case sf == 0x1 && S == 0x1 && opcode == 0x0:
		return "SUBPS"
	}

	return "UNALLOCATED"
}

func (i DataProc2MTE) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
