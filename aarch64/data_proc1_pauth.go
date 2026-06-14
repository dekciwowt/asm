package aarch64

import "fmt"

type DataProc1PAUTH Instruction

// PACIA encodes a PACIA instruction
func PACIA(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x0, 6, 10)

	return i
}

// PACIB encodes a PACIB instruction
func PACIB(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x1, 6, 10)

	return i
}

// PACDA encodes a PACDA instruction
func PACDA(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x2, 6, 10)

	return i
}

// PACDB encodes a PACDB instruction
func PACDB(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x3, 6, 10)

	return i
}

// AUTIA encodes an AUTIA instruction
func AUTIA(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x4, 6, 10)

	return i
}

// AUTIB encodes an AUTIB instruction
func AUTIB(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x5, 6, 10)

	return i
}

// AUTDA encodes an AUTDA instruction
func AUTDA(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x6, 6, 10)

	return i
}

// AUTDB encodes an AUTDB instruction
func AUTDB(rd, rn Register) DataProc1PAUTH {
	var i DataProc1PAUTH = 0x5AC00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0x1, 5, 16)
	i = set[uint8](i, 0x7, 6, 10)

	return i
}

func (i DataProc1PAUTH) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc1PAUTH) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc1PAUTH) Mnemonic() string {
	sf := get[uint8](i, 1, 31)
	opcode2 := get[uint8](i, 5, 16)
	opcode := get[uint8](i, 6, 10)

	switch {
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x0:
		return "PACIA"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x1:
		return "PACIB"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x2:
		return "PACDA"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x3:
		return "PACDB"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x4:
		return "AUTIA"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x5:
		return "AUTIB"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x6:
		return "AUTDA"
	case sf == 0x1 && opcode2 == 0x1 && opcode == 0x7:
		return "AUTDB"
	}

	return "UNALLOCATED"
}

func (i DataProc1PAUTH) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)

	return fmt.Sprintf("%s %s, %s", mnemonic, rd, rn)
}
