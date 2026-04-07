package arm64

func UDIV(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpUDIVw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpUDIVx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func SDIV(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpSDIVw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpSDIVx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func LSLV(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpLSLVw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpLSLVx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func LSRV(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpLSRVw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpLSRVx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func ASRV(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpASRVw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpASRVx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func RORV(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpRORVw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpRORVx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32B(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32B)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32H(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32H)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32W(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32W)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32CB(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32CB)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32CH(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32CH)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32CW(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32CW)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32X(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32X)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func CRC32CX(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpCRC32CX)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func SMAX(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpSMAXw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpSMAXx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func UMAX(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpUMAXw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpUMAXx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func SMIN(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpSMINw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpSMINx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func UMIN(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpUMINw)
	if W30 < rd || W30 < rn || W30 < rm {
		inst = DPInstruction(OpUMINx)
	}

	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func SUBP(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpSUBPx)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func SUBPS(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpSUBPSx)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func IRG(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpIRGx)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func GMI(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpGMIx)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func PACGA(rd, rn, rm Register) DPInstruction {
	inst := DPInstruction(OpPACGAx)
	return inst.WithRm(rm).WithRn(rn).WithRd(rd)
}

func RBIT(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpRBITw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpRBITx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func REV16(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpREV16w)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpREV16x)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func REV32(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpREV32x)
	return inst.WithRn(rn).WithRd(rd)
}

func REV(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpREVw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpREVx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func CLZ(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpCLZw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpCLZx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func CLS(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpCLSw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpCLSx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func CTZ(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpCTZw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpCTZx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func CNT(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpCNTw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpCNTx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func ABS(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpABSw)
	if W30 < rd || W30 < rn {
		inst = DPInstruction(OpABSx)
	}

	return inst.WithRn(rn).WithRd(rd)
}

func PACIA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACIAx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACIB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACIBx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACDA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACDAx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACDB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACDBx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTIA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTIAx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTIB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTIBx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTDA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTDAx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTDB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTDBx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACIZA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACIZAx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACIZB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACIZBx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACDZA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACDZAx)
	return inst.WithRn(rn).WithRd(rd)
}

func PACDZB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpPACDZBx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTIZA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTIZAx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTIZB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTIZBx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTDZA(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTDZAx)
	return inst.WithRn(rn).WithRd(rd)
}

func AUTDZB(rd, rn Register) DPInstruction {
	inst := DPInstruction(OpAUTDZBx)
	return inst.WithRn(rn).WithRd(rd)
}
