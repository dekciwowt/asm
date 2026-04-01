package arm64

func b2u(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}

func ADD(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpADD).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ADDS(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpADDS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func SUB(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpSUB).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func SUBS(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpSUBS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func AND(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpAND).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ANDS(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpANDS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ORR(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpORR).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func EOR(rd, rn, rm Register) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpEOR).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ADDI(rd, rn Register, imm uint16) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpADDI).
		WithImm12(imm).
		WithRn(rn).
		WithRd(rd)
}

func ADDSI(rd, rn Register, imm uint16) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpADDSI).
		WithImm12(imm).
		WithRn(rn).
		WithRd(rd)
}

func SUBI(rd, rn Register, imm uint16) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpSUBI).
		WithImm12(imm).
		WithRn(rn).
		WithRd(rd)
}

func SUBSI(rd, rn Register, imm uint16) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpSUBSI).
		WithImm12(imm).
		WithRn(rn).
		WithRd(rd)
}

func ANDI(rd, rn Register, imm uint64) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpANDI).
		WithBitmask(imm).
		WithRn(rn).
		WithRd(rd)
}

func ANDSI(rd, rn Register, imm uint64) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpANDSI).
		WithBitmask(imm).
		WithRn(rn).
		WithRd(rd)
}

func ORRI(rd, rn Register, imm uint64) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpORRI).
		WithBitmask(imm).
		WithRn(rn).
		WithRd(rd)
}

func EORI(rd, rn Register, imm uint64) Instruction {
	var inst Instruction
	return inst.
		WithSF(b2u(W31 < rd)).
		WithOpcode(OpEORI).
		WithBitmask(imm).
		WithRn(rn).
		WithRd(rd)
}

func ADR(rd Register, imm21 int32) Instruction {
	var inst Instruction
	return inst.
		WithSF(0).
		WithOpcode(OpADR).
		WithImm21(imm21).
		WithRd(rd)
}

func ADRP(rd Register, imm21 int32) Instruction {
	var inst Instruction
	return inst.
		WithSF(1).
		WithOpcode(OpADR).
		WithImm21(imm21).
		WithRd(rd)
}

func STR(rd, rn, rm Register) Instruction {
	var inst Instruction

	size := SizeWord
	if W31 < rd {
		size = SizeDWord
	}

	return inst.
		WithSize(size).
		WithOpcode(OpSTR).
		WithOption(OptLSL).
		WithShift(0).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func LDR(rd, rn, rm Register) Instruction {
	var inst Instruction

	size := SizeWord
	if W31 < rd {
		size = SizeDWord
	}

	return inst.
		WithSize(size).
		WithOpcode(OpLDR).
		WithOption(OptLSL).
		WithShift(0).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}
