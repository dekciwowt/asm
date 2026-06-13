package aarch64

import "fmt"

type DataProc2CRC32 Instruction

// CRC32B encodes a CRC32B instruction
func CRC32B(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 31)
	i = set[uint8](i, 0x10, 6, 10)

	return i
}

// CRC32H encodes a CRC32H instruction
func CRC32H(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 31)
	i = set[uint8](i, 0x11, 6, 10)

	return i
}

// CRC32W encodes a CRC32W instruction
func CRC32W(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 31)
	i = set[uint8](i, 0x12, 6, 10)

	return i
}

// CRC32CB encodes a CRC32CB instruction
func CRC32CB(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 31)
	i = set[uint8](i, 0x14, 6, 10)

	return i
}

// CRC32CH encodes a CRC32CH instruction
func CRC32CH(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 31)
	i = set[uint8](i, 0x15, 6, 10)

	return i
}

// CRC32CW encodes a CRC32CW instruction
func CRC32CW(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x0, 1, 31)
	i = set[uint8](i, 0x16, 6, 10)

	return i
}

// CRC32X encodes a CRC32X instruction
func CRC32X(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x13, 6, 10)

	return i
}

// CRC32CX encodes a CRC32CX instruction
func CRC32CX(rd, rn, rm Register) DataProc2CRC32 {
	var i DataProc2CRC32 = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x17, 6, 10)

	return i
}

func (i DataProc2CRC32) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc2CRC32) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc2CRC32) Rm() Register {
	return getReg(i, 5, 16)
}

func (i DataProc2CRC32) Mnemonic() string {
	sf := get[uint8](i, 1, 31)
	opcode := get[uint8](i, 6, 10)

	switch {
	case sf == 0x0 && opcode == 0x10:
		return "CRC32B"
	case sf == 0x0 && opcode == 0x11:
		return "CRC32H"
	case sf == 0x0 && opcode == 0x12:
		return "CRC32W"
	case sf == 0x0 && opcode == 0x14:
		return "CRC32CB"
	case sf == 0x0 && opcode == 0x15:
		return "CRC32CH"
	case sf == 0x0 && opcode == 0x16:
		return "CRC32CW"
	case sf == 0x1 && opcode == 0x13:
		return "CRC32X"
	case sf == 0x1 && opcode == 0x17:
		return "CRC32CX"
	}

	return "UNALLOCATED"
}

func (i DataProc2CRC32) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
