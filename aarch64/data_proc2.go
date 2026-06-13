package aarch64

import "fmt"

type DataProc2 Instruction

// UDIV encodes an UDIV instruction
func UDIV(rd, rn, rm Register) DataProc2 {
	var i DataProc2 = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x2, 6, 10)

	return i
}

// SDIV encodes a SDIV instruction
func SDIV(rd, rn, rm Register) DataProc2 {
	var i DataProc2 = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x3, 6, 10)

	return i
}

// LSLV encodes a LSLV instruction
func LSLV(rd, rn, rm Register) DataProc2 {
	var i DataProc2 = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x8, 6, 10)

	return i
}

// LSRV encodes a LSRV instruction
func LSRV(rd, rn, rm Register) DataProc2 {
	var i DataProc2 = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x9, 6, 10)

	return i
}

// ASRV encodes an ASRV instruction
func ASRV(rd, rn, rm Register) DataProc2 {
	var i DataProc2 = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0xA, 6, 10)

	return i
}

// RORV encodes a RORV instruction
func RORV(rd, rn, rm Register) DataProc2 {
	var i DataProc2 = 0x0DA00000

	i = setSF(i, rd, rn, rm)

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0xB, 6, 10)

	return i
}

func (i DataProc2) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc2) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc2) Rm() Register {
	return getReg(i, 5, 16)
}

func (i DataProc2) Mnemonic() string {
	opcode := get[uint8](i, 6, 10)

	switch opcode {
	case 0x2:
		return "UDIV"
	case 0x3:
		return "SDIV"
	case 0x8:
		return "LSLV"
	case 0x9:
		return "LSRV"
	case 0xA:
		return "ASRV"
	case 0xB:
		return "RORV"
	}

	return "UNALLOCATED"
}

func (i DataProc2) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
