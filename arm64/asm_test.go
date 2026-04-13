package arm64

import (
	"fmt"
	"strings"
	"testing"
)

type instruction interface {
	~uint32

	String() string
	Feature() Feature
}

type instructionTests[I instruction] [][2]I

func (ts instructionTests[I]) Test(t *testing.T) {
	var b strings.Builder

	for _, tuple := range ts {
		inst, exp := tuple[0], tuple[1]

		fmt.Fprintf(&b, "Instruction:     %s\n", inst)
		fmt.Fprintf(&b, "Feature:         %s\n", inst.Feature())
		fmt.Fprintf(&b, "Encoded HEX:     %X\n", uint32(inst))
		fmt.Fprintf(&b, "Expected HEX:    %X\n", uint32(exp))
		fmt.Fprintf(&b, "Encoded Binary:  %032b\n", uint32(inst))
		fmt.Fprintf(&b, "Expected Binary: %032b\n", uint32(exp))
		fmt.Fprintf(&b, "Correct:         %t\n", inst == exp)

		t.Logf("\n%s\n", b.String())

		b.Reset()

		if inst != exp {
			t.Fail()
		}
	}
}

func TestEncoding(t *testing.T) {
	dp3Source := instructionTests[DataProc3Source]{
		{MADD(W0, W1, W2, W3), 0x1B020C20},
		{MSUB(W0, W1, W2, W3), 0x1B028C20},
		{MADD(X0, X1, X2, X3), 0x9B020C20},
		{MSUB(X0, X1, X2, X3), 0x9B028C20},
		{SMADDL(X0, W1, W2, X3), 0x9B220C20},
		{SMSUBL(X0, W1, W2, X3), 0x9B228C20},
		{SMULH(X0, X1, X2), 0x9B427C20},
		{UMADDL(X0, W1, W2, X3), 0x9BA20C20},
		{UMSUBL(X0, W1, W2, X3), 0x9BA28C20},
		{UMULH(X0, X1, X2), 0x9BC27C20},
		{MADDPT(X0, W1, W2, X3), 0x9B620C20},
		{MSUBPT(X0, W1, W2, X3), 0x9B628C20},
	}

	dp3Source.Test(t)

	dp2Source := instructionTests[DataProc2Source]{
		{UDIV(W0, W1, W2), 0x1AC20820},
		{SDIV(W0, W1, W2), 0x1AC20C20},
		{LSLV(W0, W1, W2), 0x1AC22020},
		{LSRV(W0, W1, W2), 0x1AC22420},
		{ASRV(W0, W1, W2), 0x1AC22820},
		{RORV(W0, W1, W2), 0x1AC22C20},
		{UDIV(X0, X1, X2), 0x9AC20820},
		{SDIV(X0, X1, X2), 0x9AC20C20},
		{LSLV(X0, X1, X2), 0x9AC22020},
		{LSRV(X0, X1, X2), 0x9AC22420},
		{ASRV(X0, X1, X2), 0x9AC22820},
		{RORV(X0, X1, X2), 0x9AC22C20},

		// feature CRC32

		{CRC32B(W0, W1, W2), 0x1AC24020},
		{CRC32H(W0, W1, W2), 0x1AC24420},
		{CRC32W(W0, W1, W2), 0x1AC24820},
		{CRC32CB(W0, W1, W2), 0x1AC25020},
		{CRC32CH(W0, W1, W2), 0x1AC25420},
		{CRC32CW(W0, W1, W2), 0x1AC25820},
		{CRC32X(W0, W1, X2), 0x9AC24C20},
		{CRC32CX(W0, W1, X2), 0x9AC25C20},

		// feature CSSC

		{SMAX(W0, W1, W2), 0x1AC26020},
		{UMAX(W0, W1, W2), 0x1AC26420},
		{SMIN(W0, W1, W2), 0x1AC26820},
		{UMIN(W0, W1, W2), 0x1AC26C20},
		{SMAX(X0, X1, X2), 0x9AC26020},
		{UMAX(X0, X1, X2), 0x9AC26420},
		{SMIN(X0, X1, X2), 0x9AC26820},
		{UMIN(X0, X1, X2), 0x9AC26C20},

		// feature MTE

		{SUBP(X0, X1, X2), 0x9AC20020},
		{SUBPS(X0, X1, X2), 0xBAC20020},
		{IRG(X0, X1, X2), 0x9AC21020},
		{GMI(X0, X1, X2), 0x9AC21420},

		// feature PAuth

		{PACGA(X0, X1, X2), 0x9AC23020},
	}

	dp2Source.Test(t)

	dp1Source := instructionTests[DataProc1Source]{
		{RBIT(W0, W1), 0x5AC00020},
		{REV16(W0, W1), 0x5AC00420},
		{REV(W0, W1), 0x5AC00820},
		{CLZ(W0, W1), 0x5AC01020},
		{CLS(W0, W1), 0x5AC01420},
		{RBIT(X0, X1), 0xDAC00020},
		{REV16(X0, X1), 0xDAC00420},
		{REV32(X0, X1), 0xDAC00820},
		{REV(X0, X1), 0xDAC00C20},
		{CLZ(X0, X1), 0xDAC01020},
		{CLS(X0, X1), 0xDAC01420},

		// feature CSSC

		{CTZ(W0, W1), 0x5AC01820},
		{CNT(W0, W1), 0x5AC01C20},
		{ABS(W0, W1), 0x5AC02020},
		{CTZ(X0, X1), 0xDAC01820},
		{CNT(X0, X1), 0xDAC01C20},
		{ABS(X0, X1), 0xDAC02020},

		// feature PAuth

		{PACIA(X0, X1), 0xDAC10020},
		{PACIB(X0, X1), 0xDAC10420},
		{PACDA(X0, X1), 0xDAC10820},
		{PACDB(X0, X1), 0xDAC10C20},
		{AUTIA(X0, X1), 0xDAC11020},
		{AUTIB(X0, X1), 0xDAC11420},
		{AUTDA(X0, X1), 0xDAC11820},
		{AUTDB(X0, X1), 0xDAC11C20},
		{PACIZA(X0), 0xDAC123E0},
		{PACIZB(X0), 0xDAC127E0},
		{PACDZA(X0), 0xDAC12BE0},
		{PACDZB(X0), 0xDAC12FE0},
		{AUTIZA(X0), 0xDAC133E0},
		{AUTIZB(X0), 0xDAC137E0},
		{AUTDZA(X0), 0xDAC13BE0},
		{AUTDZB(X0), 0xDAC13FE0},
		{XPACI(X0), 0xDAC143E0},
		{XPACD(X0), 0xDAC147E0},
		{PACNBIASPPC(), 0xDAC183FE},
		{PACNBIBSPPC(), 0xDAC187FE},
		{PACIA171615(), 0xDAC18BFE},
		{PACIB171615(), 0xDAC18FFE},
		{AUTIASPPCR(), 0xDAC193FE},
		{AUTIBSPPCR(), 0xDAC197FE},
		{PACIASPPC(), 0xDAC1A3FE},
		{PACIBSPPC(), 0xDAC1A7FE},
		{AUTIA171615(), 0xDAC1BBFE},
		{AUTIB171615(), 0xDAC1BFFE},
	}

	dp1Source.Test(t)

	dpLogicReg := instructionTests[DataProcLogicReg]{
		{AND(W0, W1, W2), 0x0A020020},
		{BIC(W0, W1, W2), 0x0A220020},
		{ORR(W0, W1, W2), 0x2A020020},
		{ORN(W0, W1, W2), 0x2A220020},
		{EOR(W0, W1, W2), 0x4A020020},
		{EON(W0, W1, W2), 0x4A220020},
		{ANDS(W0, W1, W2), 0x6A020020},
		{BICS(W0, W1, W2), 0x6A220020},
		{AND(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x0A020420},
		{BIC(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x0A220420},
		{ORR(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x2A020420},
		{ORN(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x2A220420},
		{EOR(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x4A020420},
		{EON(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x4A220420},
		{ANDS(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x6A020420},
		{BICS(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x6A220420},
		{AND(X0, X1, X2), 0x8A020020},
		{BIC(X0, X1, X2), 0x8A220020},
		{ORR(X0, X1, X2), 0xAA020020},
		{ORN(X0, X1, X2), 0xAA220020},
		{EOR(X0, X1, X2), 0xCA020020},
		{EON(X0, X1, X2), 0xCA220020},
		{ANDS(X0, X1, X2), 0xEA020020},
		{BICS(X0, X1, X2), 0xEA220020},
		{AND(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0x8A020420},
		{BIC(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0x8A220420},
		{ORR(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xAA020420},
		{ORN(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xAA220420},
		{EOR(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xCA020420},
		{EON(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xCA220420},
		{ANDS(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xEA020420},
		{BICS(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xEA220420},
	}

	dpLogicReg.Test(t)

	dpArithReg := instructionTests[DataProcArithReg]{
		{ADD(W0, W1, W2), 0x0B020020},
		{ADDS(W0, W1, W2), 0x2B020020},
		{SUB(W0, W1, W2), 0x4B020020},
		{SUBS(W0, W1, W2), 0x6B020020},
		{ADD(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x0B020420},
		{ADDS(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x2B020420},
		{SUB(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x4B020420},
		{SUBS(W0, W1, W2).WithShift(ShiftLSL, 0x1), 0x6B020420},
		{ADD(W0, W1, W2).WithExtension(ExtUXTB, 0x1), 0x0B220420},
		{ADDS(W0, W1, W2).WithExtension(ExtUXTB, 0x1), 0x2B220420},
		{SUB(W0, W1, W2).WithExtension(ExtUXTB, 0x1), 0x4B220420},
		{SUBS(W0, W1, W2).WithExtension(ExtUXTB, 0x1), 0x6B220420},
		{ADD(X0, X1, X2), 0x8B020020},
		{ADDS(X0, X1, X2), 0xAB020020},
		{SUB(X0, X1, X2), 0xCB020020},
		{SUBS(X0, X1, X2), 0xEB020020},
		{ADD(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0x8B020420},
		{ADDS(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xAB020420},
		{SUB(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xCB020420},
		{SUBS(X0, X1, X2).WithShift(ShiftLSL, 0x1), 0xEB020420},
		{ADD(X0, X1, X2).WithExtension(ExtUXTB, 0x1), 0x8B220420},
		{ADDS(X0, X1, X2).WithExtension(ExtUXTB, 0x1), 0xAB220420},
		{SUB(X0, X1, X2).WithExtension(ExtUXTB, 0x1), 0xCB220420},
		{SUBS(X0, X1, X2).WithExtension(ExtUXTB, 0x1), 0xEB220420},
	}

	dpArithReg.Test(t)

	dpArithWithCarry := instructionTests[DataProcArithWithCarry]{
		{ADC(W0, W1, W2), 0x1A020020},
		{ADCS(W0, W1, W2), 0x3A020020},
		{SBC(W0, W1, W2), 0x5A020020},
		{SBCS(W0, W1, W2), 0x7A020020},
		{ADC(X0, X1, X2), 0x9A020020},
		{ADCS(X0, X1, X2), 0xBA020020},
		{SBC(X0, X1, X2), 0xDA020020},
		{SBCS(X0, X1, X2), 0xFA020020},
	}

	dpArithWithCarry.Test(t)

	dpArithCkPtr := instructionTests[DataProcArithCkPtr]{
		{ADDPT(X0, X1, X2), 0x9A022020},
		{SUBPT(X0, X1, X2), 0xDA022020},
		{ADDPT(X0, X1, X2).WithShift(0x1), 0x9A022420},
		{SUBPT(X0, X1, X2).WithShift(0x1), 0xDA022420},
	}

	dpArithCkPtr.Test(t)

	dpRotate := instructionTests[DataProcRotate]{
		{RMIF(X0, 0x1, 0x3), 0xBA008403},
	}

	dpRotate.Test(t)

	dpEvaluate := instructionTests[DataProcEvaluate]{
		{SETF8(X0), 0x3A00080D},
		{SETF16(X0), 0x3A00480D},
	}

	dpEvaluate.Test(t)

	dpCondCompReg := instructionTests[DataProcCondCompReg]{
		{CCMN(W0, 0x1, CondEQ), 0x3A400001},
		{CCMP(W0, 0x1, CondEQ), 0x7A400001},
		{CCMN(X0, 0x1, CondEQ), 0xBA400001},
		{CCMP(X0, 0x1, CondEQ), 0xFA400001},
	}

	dpCondCompReg.Test(t)
}
