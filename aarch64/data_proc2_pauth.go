package aarch64

import "fmt"

type DataProc2PAUTH Instruction

// PACGA encodes a SUBP instruction
func PACGA(rd, rn, rm Register) DataProc2PAUTH {
	var i DataProc2PAUTH = 0x0DA00000

	i = setReg(i, rd, 5, 0)
	i = setReg(i, rn, 5, 5)
	i = setReg(i, rm, 5, 16)

	i = set[uint8](i, 0x1, 1, 31)
	i = set[uint8](i, 0xC, 6, 10)

	return i
}

func (i DataProc2PAUTH) Rd() Register {
	return getReg(i, 5, 0)
}

func (i DataProc2PAUTH) Rn() Register {
	return getReg(i, 5, 5)
}

func (i DataProc2PAUTH) Rm() Register {
	return getReg(i, 5, 16)
}

func (i DataProc2PAUTH) Mnemonic() string {
	sf := get[uint8](i, 1, 31)
	opcode := get[uint8](i, 6, 10)

	switch {
	case sf == 0x1 && opcode == 0xC:
		return "PACGA"
	}

	return "UNALLOCATED"
}

func (i DataProc2PAUTH) String() string {
	mnemonic := i.Mnemonic()

	rd := getReg(i, 5, 0)
	rn := getReg(i, 5, 5)
	rm := getReg(i, 5, 16)

	return fmt.Sprintf("%s %s, %s, %s", mnemonic, rd, rn, rm)
}
