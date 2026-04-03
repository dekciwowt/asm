package arm64

// The sf bit is derived from the register arguments: if any register is
// an X-register (> W30), the instruction operates in 64-bit mode.
// All registers in a single instruction must be consistently 32 or 64-bit —
// mixing W and X registers in the same instruction is not valid ARM64.

// ADD encodes an ADD (plain register) instruction
//
//	ADD  Wd, Wn, Wm   ; 32-bit
//	ADD  Xd, Xn, Xm   ; 64-bit
func ADD(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpADD).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ADDS encodes an ADDS (plain register) instruction
//
//	ADDS  Wd, Wn, Wm   ; 32-bit
//	ADDS  Xd, Xn, Xm   ; 64-bit
func ADDS(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpADDS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SUB encodes a SUB (plain register) instruction
//
//	SUB  Wd, Wn, Wm   ; 32-bit
//	SUB  Xd, Xn, Xm   ; 64-bit
func SUB(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpSUB).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SUBS encodes a SUBS (plain register) instruction
//
//	SUBS  Wd, Wn, Wm   ; 32-bit
//	SUBS  Xd, Xn, Xm   ; 64-bit
func SUBS(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpSUBS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// AND encodes an AND (plain register) instruction
//
//	AND  Wd, Wn, Wm   ; 32-bit
//	AND  Xd, Xn, Xm   ; 64-bit
func AND(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpAND).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ANDS encodes an ANDS (plain register) instruction
//
//	ANDS  Wd, Wn, Wm   ; 32-bit
//	ANDS  Xd, Xn, Xm   ; 64-bit
func ANDS(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpANDS).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ORR encodes an ORR (plain register) instruction
//
//	ORR  Wd, Wn, Wm   ; 32-bit
//	ORR  Xd, Xn, Xm   ; 64-bit
func ORR(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpORR).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// EOR encodes an EOR (plain register) instruction
//
//	EOR  Wd, Wn, Wm   ; 32-bit
//	EOR  Xd, Xn, Xm   ; 64-bit
func EOR(rd, rn, rm Register) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithOpcode(OpEOR).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ADDI encodes an ADD (immediate) instruction
//
//	ADD  Wd, Wn, #imm   ; 32-bit
//	ADD  Xd, Xn, #imm   ; 64-bit
func ADDI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpADDI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// ADDSI encodes an ADDS (immediate) instruction
//
//	ADDS  Wd, Wn, #imm   ; 32-bit
//	ADDS  Xd, Xn, #imm   ; 64-bit
func ADDSI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpADDSI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// SUBI encodes a SUB (immediate) instruction
//
//	SUB  Wd, Wn, #imm   ; 32-bit
//	SUB  Xd, Xn, #imm   ; 64-bit
func SUBI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpSUBI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// SUBSI encodes a SUBS (immediate) instruction
//
//	SUBS  Wd, Wn, #imm   ; 32-bit
//	SUBS  Xd, Xn, #imm   ; 64-bit
func SUBSI(rd, rn Register, imm uint16) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpSUBSI).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// ANDI encodes an AND (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	AND  Wd, Wn, #bitmask   ; 32-bit
//	AND  Xd, Xn, #bitmask   ; 64-bit
func ANDI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpANDI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// ANDSI encodes an ANDS (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	ANDS  Wd, Wn, #bitmask   ; 32-bit
//	ANDS  Xd, Xn, #bitmask   ; 64-bit
func ANDSI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpANDSI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// ORRI encodes an ORR (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	ORR  Wd, Wn, #bitmask   ; 32-bit
//	ORR  Xd, Xn, #bitmask   ; 64-bit
func ORRI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpORRI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// EORI encodes an EOR (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	EOR  Wd, Wn, #bitmask   ; 32-bit
//	EOR  Xd, Xn, #bitmask   ; 64-bit
func EORI(rd, rn Register, bitmask uint64) DPInstruction {
	var inst DPInstruction
	return inst.
		WithSF(W30 < rd || W30 < rn).
		WithOpcode(OpEORI).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}
