package arm64

func MADD(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source
	if i = instMADDw; W30 < rd {
		i = instMADDx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func MSUB(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source
	if i = instMSUBw; W30 < rd {
		i = instMSUBx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func MUL(rd, rn, rm Register) DataProc3Source {
	if W30 < rd {
		return MADD(rd, rn, rm, XZR)
	}
	return MADD(rd, rn, rm, WZR)
}

func MNEG(rd, rn, rm Register) DataProc3Source {
	if W30 < rd {
		return MSUB(rd, rn, rm, XZR)
	}
	return MSUB(rd, rn, rm, WZR)
}

func SMADDL(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source = instSMADDL
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func SMSUBL(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source = instSMSUBL
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func SMULH(rd, rn, rm Register) DataProc3Source {
	var i DataProc3Source = instSMULH
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(XZR)
}

func UMADDL(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source = instUMADDL
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func UMSUBL(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source = instUMSUBL
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func UMULH(rd, rn, rm Register) DataProc3Source {
	var i DataProc3Source = instUMULH
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(XZR)
}

func SMULL(rd, rn, rm Register) DataProc3Source {
	return SMADDL(rd, rn, rm, XZR)
}

func UMULL(rd, rn, rm Register) DataProc3Source {
	return UMADDL(rd, rn, rm, XZR)
}

func MADDPT(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source = instMADDPT
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func MSUBPT(rd, rn, rm, ra Register) DataProc3Source {
	var i DataProc3Source = instMSUBPT
	return i.WithRd(rd).WithRn(rn).WithRm(rm).WithRa(ra)
}

func UDIV(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instUDIVw; W30 < rd {
		i = instUDIVx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func SDIV(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instSDIVw; W30 < rd {
		i = instSDIVx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func LSLV(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instLSLVw; W30 < rd {
		i = instLSLVx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func LSRV(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instLSRVw; W30 < rd {
		i = instLSRVx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func ASRV(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instASRVw; W30 < rd {
		i = instASRVx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func RORV(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instRORVw; W30 < rd {
		i = instRORVx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32B(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32B
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32H(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32H
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32W(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32W
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32CB(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32CB
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32CH(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32CH
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32CW(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32CW
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32X(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32X
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func CRC32CX(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instCRC32CX
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func SMAX(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instSMAXw; W30 < rd {
		i = instSMAXx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func UMAX(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instUMAXw; W30 < rd {
		i = instUMAXx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func SMIN(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instSMINw; W30 < rd {
		i = instSMINx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func UMIN(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source
	if i = instUMINw; W30 < rd {
		i = instUMINx
	}
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func SUBP(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instSUBP
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func SUBPS(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instSUBPS
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func IRG(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instIRG
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func GMI(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instGMI
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func PACGA(rd, rn, rm Register) DataProc2Source {
	var i DataProc2Source = instPACGA
	return i.WithRd(rd).WithRn(rn).WithRm(rm)
}

func RBIT(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instRBITw; W30 < rd {
		i = instRBITx
	}
	return i.WithRd(rd).WithRn(rn)
}

func REV16(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instREV16w; W30 < rd {
		i = instREV16x
	}
	return i.WithRd(rd).WithRn(rn)
}

func REV32(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instREV32x
	return i.WithRd(rd).WithRn(rn)
}

func REV(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instREVw; W30 < rd {
		i = instREVx
	}
	return i.WithRd(rd).WithRn(rn)
}

func CLZ(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instCLZw; W30 < rd {
		i = instCLZx
	}
	return i.WithRd(rd).WithRn(rn)
}

func CLS(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instCLSw; W30 < rd {
		i = instCLSx
	}
	return i.WithRd(rd).WithRn(rn)
}

func CTZ(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instCTZw; W30 < rd {
		i = instCTZx
	}
	return i.WithRd(rd).WithRn(rn)
}

func CNT(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instCNTw; W30 < rd {
		i = instCNTx
	}
	return i.WithRd(rd).WithRn(rn)
}

func ABS(rd, rn Register) DataProc1Source {
	var i DataProc1Source
	if i = instABSw; W30 < rd {
		i = instABSx
	}
	return i.WithRd(rd).WithRn(rn)
}

func PACIA(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instPACIA
	return i.WithRd(rd).WithRn(rn)
}

func PACIB(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instPACIB
	return i.WithRd(rd).WithRn(rn)
}

func PACDA(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instPACDA
	return i.WithRd(rd).WithRn(rn)
}

func PACDB(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instPACDB
	return i.WithRd(rd).WithRn(rn)
}

func AUTIA(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instAUTIA
	return i.WithRd(rd).WithRn(rn)
}

func AUTIB(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instAUTIB
	return i.WithRd(rd).WithRn(rn)
}

func AUTDA(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instAUTDA
	return i.WithRd(rd).WithRn(rn)
}

func AUTDB(rd, rn Register) DataProc1Source {
	var i DataProc1Source = instAUTDB
	return i.WithRd(rd).WithRn(rn)
}

func PACIZA(rd Register) DataProc1Source {
	var i DataProc1Source = instPACIZA
	return i.WithRd(rd)
}

func PACIZB(rd Register) DataProc1Source {
	var i DataProc1Source = instPACIZB
	return i.WithRd(rd)
}

func PACDZA(rd Register) DataProc1Source {
	var i DataProc1Source = instPACDZA
	return i.WithRd(rd)
}

func PACDZB(rd Register) DataProc1Source {
	var i DataProc1Source = instPACDZB
	return i.WithRd(rd)
}

func AUTIZA(rd Register) DataProc1Source {
	var i DataProc1Source = instAUTIZA
	return i.WithRd(rd)
}

func AUTIZB(rd Register) DataProc1Source {
	var i DataProc1Source = instAUTIZB
	return i.WithRd(rd)
}

func AUTDZA(rd Register) DataProc1Source {
	var i DataProc1Source = instAUTDZA
	return i.WithRd(rd)
}

func AUTDZB(rd Register) DataProc1Source {
	var i DataProc1Source = instAUTDZB
	return i.WithRd(rd)
}

func XPACI(rd Register) DataProc1Source {
	var i DataProc1Source = instXPACI
	return i.WithRd(rd)
}

func XPACD(rd Register) DataProc1Source {
	var i DataProc1Source = instXPACD
	return i.WithRd(rd)
}
