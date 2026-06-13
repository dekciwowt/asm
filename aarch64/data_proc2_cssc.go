package aarch64

import "fmt"

type DataProc2CSSC Instruction

// SMAX encodes a SMAX instruction
func SMAX(rd, rn, rm Register) DataProc2CSSC {
	var i DataProc2CSSC = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x18, 6, 10)

	return i
}

// UMAX encodes an UMAX instruction
func UMAX(rd, rn, rm Register) DataProc2CSSC {
	var i DataProc2CSSC = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x19, 6, 10)

	return i
}

// SMIN encodes a SMIN instruction
func SMIN(rd, rn, rm Register) DataProc2CSSC {
	var i DataProc2CSSC = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1A, 6, 10)

	return i
}

// UMIN encodes an UMIN instruction
func UMIN(rd, rn, rm Register) DataProc2CSSC {
	var i DataProc2CSSC = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1B, 6, 10)

	return i
}

func (i DataProc2CSSC) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc2CSSC) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc2CSSC) Rm() Register {
	return getReg(i, 5, 16)
}

func (i DataProc2CSSC) Mnemonic() string {
	opcode := get[uint8](i, 6, 10)

	switch opcode {
	case 0x18:
		return "SMAX"
	case 0x19:
		return "UMAX"
	case 0x1A:
		return "SMIN"
	case 0x1B:
		return "UMIN"
	}

	return "UNALLOCATED"
}

func (i DataProc2CSSC) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
