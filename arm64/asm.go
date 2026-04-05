package arm64

// The `sf` bit is derived from the register arguments: if any register is
// an X-register (> W30), the instruction operates in 64-bit mode.
// All registers in a single instruction must be consistently 32 or 64-bit

// ADD encodes an ADD (plain register) instruction
//
//	ADD <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after ADD:
//
//	ADD(X0, X1, X2).WithRmShift(ShiftLSL, 0x3) // ADD x0, x1, x2, lsl #0x3
//
// To encode an extended-register form, chain WithRmExt after ADD:
//
//	ADD(X0, X1, X2).WithRmExt(ExtSXTW, 0x2) // ADD x0, x1, x2, sxtw #0x2
func ADD(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpADD).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ADDS encodes an ADDS (plain register) instruction
//
//	ADDS <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after ADDS:
//
//	ADDS(X0, X1, X2).WithRmShift(ShiftLSL, 0x3) // ADDS x0, x1, x2, lsl #0x3
//
// To encode an extended-register form, chain WithRmExt after ADDS:
//
//	ADDS(X0, X1, X2).WithRmExt(ExtSXTW, 0x2) // ADDS x0, x1, x2, sxtw #0x2
func ADDS(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpADDS).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SUB encodes a SUB (plain register) instruction
//
//	SUB <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after SUB:
//
//	SUB(X0, X1, X2).WithRmShift(ShiftLSL, 0x3) // SUB x0, x1, x2, lsl #0x3
//
// To encode an extended-register form, chain WithRmExt after SUB:
//
//	SUB(X0, X1, X2).WithRmExt(ExtSXTW, 0x2) // SUB x0, x1, x2, sxtw #0x2
func SUB(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpSUB).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SUBS encodes a SUBS (plain register) instruction
//
//	SUBS <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after SUBS:
//
//	SUBS(X0, X1, X2).WithRmShift(ShiftLSL, 0x3) // SUBS x0, x1, x2, lsl #0x3
//
// To encode an extended-register form, chain WithRmExt after SUBS:
//
//	SUBS(X0, X1, X2).WithRmExt(ExtSXTW, 0x2) // SUBS x0, x1, x2, sxtw #0x2
func SUBS(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpSUBS).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// AND encodes an AND (plain register) instruction
//
//	AND <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after AND:
//
//	AND(X0, X1, X2).WithRmShift(ShiftLSL, 0x4) // AND x0, x1, x2, lsl #0x4
func AND(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpAND).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ANDS encodes an ANDS (plain register) instruction
//
//	ANDS <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after ANDS:
//
//	ANDS(X0, X1, X2).WithRmShift(ShiftLSL, 0x4) // ANDS x0, x1, x2, lsl #0x4
func ANDS(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpANDS).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ORR encodes an ORR (plain register) instruction
//
//	ORR <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after ORR:
//
//	ORR(X0, X1, X2).WithRmShift(ShiftLSL, 0x4)  // ORR x0, x1, x2, lsl #0x4
func ORR(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpORR).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// EOR encodes an EOR (plain register) instruction
//
//	EOR <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
//
// To encode a shifted-register form, chain WithRmShift after EOR:
//
//	EOR(X0, X1, X2).WithRmShift(ShiftLSL, 0x4)  // EOR x0, x1, x2, lsl #0x4
func EOR(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpEOR).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ADDI encodes an ADD (immediate) instruction
//
//	ADD <Wd|Xd>, <Wn|Xn>, #imm
//
// To encode the shifted form (imm << 12), chain WithImmShift after ADDI:
//
//	ADDI(X0, X1, 0x1).WithImmShift(true)  // ADD X0, X1, #0x1 lsl #12
func ADDI(rd, rn Register, imm uint16) DPInstruction {
	return DPInstruction(OpADDI).
		WithSF(W30 < rd || W30 < rn).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// ADDSI encodes an ADDS (immediate) instruction
//
//	ADDS <Wd|Xd>, <Wn|Xn>, #imm
//
// To encode the shifted form (imm << 12), chain WithImmShift after ADDSI:
//
//	ADDSI(X0, X1, 0x1).WithImmShift(true)  // ADDS X0, X1, #0x1 lsl #12
func ADDSI(rd, rn Register, imm uint16) DPInstruction {
	return DPInstruction(OpADDSI).
		WithSF(W30 < rd || W30 < rn).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// SUBI encodes a SUB (immediate) instruction
//
//	SUB <Wd|Xd>, <Wn|Xn>, #imm
//
// To encode the shifted form (imm << 12), chain WithImmShift after SUBI:
//
//	SUBI(X0, X1, 0x1).WithImmShift(true)  // SUB X0, X1, #0x1 lsl #12
func SUBI(rd, rn Register, imm uint16) DPInstruction {
	return DPInstruction(OpSUBI).
		WithSF(W30 < rd || W30 < rn).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// SUBSI encodes a SUBS (immediate) instruction
//
//	SUBS <Wd|Xd>, <Wn|Xn>, #imm
//
// To encode the shifted form (imm << 12), chain WithImmShift after SUBSI:
//
//	SUBSI(X0, X1, 0x1).WithImmShift(true)  // SUBS X0, X1, #0x1 lsl #12
func SUBSI(rd, rn Register, imm uint16) DPInstruction {
	return DPInstruction(OpSUBSI).
		WithSF(W30 < rd || W30 < rn).
		WithImmediate(imm).
		WithRn(rn).
		WithRd(rd)
}

// ANDI encodes an AND (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	AND <Wd|Xd>, <Wn|Xn>, #bitmask
func ANDI(rd, rn Register, bitmask uint64) DPInstruction {
	return DPInstruction(OpANDI).
		WithSF(W30 < rd || W30 < rn).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// ANDSI encodes an ANDS (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	ANDS <Wd|Xd>, <Wn|Xn>, #bitmask
func ANDSI(rd, rn Register, bitmask uint64) DPInstruction {
	return DPInstruction(OpANDSI).
		WithSF(W30 < rd || W30 < rn).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// ORRI encodes an ORR (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	ORR <Wd|Xd>, <Wn|Xn>, #bitmask
func ORRI(rd, rn Register, bitmask uint64) DPInstruction {
	return DPInstruction(OpORRI).
		WithSF(W30 < rd || W30 < rn).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// EORI encodes an EOR (immediate) instruction.
// bitmask must be a valid ARM64 bitmask immediate — a replicated pattern
// of contiguous ones
//
//	EOR <Wd|Xd>, <Wn|Xn>, #bitmask
func EORI(rd, rn Register, bitmask uint64) DPInstruction {
	return DPInstruction(OpEORI).
		WithSF(W30 < rd || W30 < rn).
		WithBitmask(bitmask).
		WithRn(rn).
		WithRd(rd)
}

// ADC encodes an ADC (plain register) instruction
//
//	ADC <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
func ADC(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpADC).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ADCS encodes an ADCS (plain register) instruction
//
//	ADCS <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
func ADCS(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpADCS).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SBC encodes a SBC (plain register) instruction
//
//	SBC <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
func SBC(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpSBC).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SBCS encodes a SBCS (plain register) instruction
//
//	SBCS <Wd|Xd>, <Wn|Xn>, <Wm|Xm>
func SBCS(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpSBCS).
		WithSF(W30 < rd || W30 < rn || W30 < rm).
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// ADDPT encodes a ADDPT (plain register) instruction
//
//	ADDPT <Xd>, <Xn>, <Xm>
//
// To encode a shifted-register form, chain WithRmShift after ADDPT:
//
//	ADDPT(X0, X1, X2).WithRmShift(ShiftLSL, 0x3) // ADDPT x0, x1, x2, lsl #0x3
func ADDPT(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpADDPT).
		WithSF(true). // only 64-bit registers allowed
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}

// SUBPT encodes a SUBPT (plain register) instruction
//
//	SUBPT <Xd>, <Xn>, <Xm>
//
// To encode a shifted-register form, chain WithRmShift after SUBPT:
//
//	SUBPT(X0, X1, X2).WithRmShift(ShiftLSL, 0x3) // SUBPT x0, x1, x2, lsl #0x3
func SUBPT(rd, rn, rm Register) DPInstruction {
	return DPInstruction(OpSUBPT).
		WithSF(true). // only 64-bit registers allowed
		WithRm(rm).
		WithRn(rn).
		WithRd(rd)
}
