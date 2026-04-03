package arm64

func ADD(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpADD).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ADDS(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpADDS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func SUB(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpSUB).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func SUBS(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpSUBS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func AND(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpAND).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ANDS(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpANDS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ORR(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpORR).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func EOR(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpEOR).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

func ADDI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpADDI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

func ADDSI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpADDSI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

func SUBI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpSUBI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

func SUBSI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpSUBSI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

func ANDI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpANDI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

func ANDSI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpANDSI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

func ORRI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpORRI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

func EORI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpEORI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}
