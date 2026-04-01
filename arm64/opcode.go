package arm64

import "fmt"

type Opcode uint8

const /* Opcode */ (
	OpUND Opcode = 0x00

	// Data Processing – Register opcodes

	OpADD  Opcode = 0x0B
	OpADDS Opcode = 0x2B
	OpSUB  Opcode = 0x4B
	OpSUBS Opcode = 0x6B
	OpAND  Opcode = 0x0A
	OpORR  Opcode = 0x2A
	OpEOR  Opcode = 0x4A
	OpANDS Opcode = 0x6A

	// Data Processing – Immediate opcodes

	OpADDI  Opcode = 0x11
	OpADDSI Opcode = 0x31
	OpSUBI  Opcode = 0x51
	OpSUBSI Opcode = 0x71
	OpANDI  Opcode = 0x12
	OpANDSI Opcode = 0x72
	OpORRI  Opcode = 0x32
	OpEORI  Opcode = 0x52

	// PC-relative opcodes

	OpADR Opcode = 0x10

	// Load/Store – Register opcodes
	//
	// Uses 8-bit wide opcodes and encoded as:
	//
	//  7     6-2       1-0
	//	+-----+---------+-----+
	//	|  1  | opcode  | opc |
	//	+-----+---------+-----+

	OpSTR  Opcode = 0xE0
	OpLDR  Opcode = 0xE1
	OpLDRS Opcode = 0xE2
)

var opNames = map[Opcode]string{
	OpUND:   "UND",
	OpADD:   "ADD",
	OpADDS:  "ADDS",
	OpSUB:   "SUB",
	OpSUBS:  "SUBS",
	OpAND:   "AND",
	OpORR:   "ORR",
	OpEOR:   "EOR",
	OpANDS:  "ANDS",
	OpADDI:  "ADD",
	OpADDSI: "ADDS",
	OpSUBI:  "SUB",
	OpSUBSI: "SUBS",
	OpANDI:  "AND",
	OpANDSI: "ANDS",
	OpORRI:  "ORR",
	OpEORI:  "EOR",
	OpADR:   "ADR",
	OpSTR:   "STR",
	OpLDR:   "LDR",
	OpLDRS:  "LDRS",
}

func (o Opcode) String() string {
	if name, ok := opNames[o]; ok {
		return name
	}

	return fmt.Sprintf("Opcode(%#07X)", uint8(o))
}
