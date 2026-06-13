package aarch64

func ADD(rd, rn, rm Register) ArithShift {
	return ADDShift(rd, rn, rm, ShiftLSL, 0x0)
}

func ADDS(rd, rn, rm Register) ArithShift {
	return ADDSShift(rd, rn, rm, ShiftLSL, 0x0)
}

func SUB(rd, rn, rm Register) ArithShift {
	return SUBShift(rd, rn, rm, ShiftLSL, 0x0)
}

func SUBS(rd, rn, rm Register) ArithShift {
	return SUBSShift(rd, rn, rm, ShiftLSL, 0x0)
}
