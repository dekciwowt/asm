package arm64

import (
	"fmt"
	"strings"
	"testing"
)

type InstructionTests[I Instruction] [][2]uint32

func (ts InstructionTests[I]) Test(t *testing.T) {
	var b strings.Builder

	for _, tuple := range ts {
		inst, exp := tuple[0], tuple[1]

		fmt.Fprintf(&b, "Instruction:     %s\n", I(inst))
		fmt.Fprintf(&b, "Encoded HEX:     %X\n", inst)
		fmt.Fprintf(&b, "Expected HEX:    %X\n", exp)
		fmt.Fprintf(&b, "Encoded Binary:  %032b\n", inst)
		fmt.Fprintf(&b, "Expected Binary: %032b\n", exp)
		fmt.Fprintf(&b, "Correct:         %t\n", inst == exp)

		t.Logf("\n%s\n", b.String())

		b.Reset()

		if inst != exp {
			t.Fail()
		}
	}
}

func TestDPInstructions(t *testing.T) {
	tests := InstructionTests[DPInstruction]{
		// 32-bit (sf=0)
		{uint32(ADD(W0, W1, W2)), 0x0B020020},
		{uint32(ADD(W0, W1, W2).WithExtension(ExtUXTX, 0x1)), 0x0B226420},
		{uint32(ADDS(W0, W1, W2)), 0x2B020020},
		{uint32(SUB(W0, W1, W2)), 0x4B020020},
		{uint32(SUBS(W0, W1, W2)), 0x6B020020},
		{uint32(AND(W0, W1, W2)), 0x0A020020},
		{uint32(ANDS(W0, W1, W2)), 0x6A020020},
		{uint32(ORR(W0, W1, W2)), 0x2A020020},
		{uint32(EOR(W0, W1, W2)), 0x4A020020},
		{uint32(ADDI(W0, W1, 42)), 0x1100A820},
		{uint32(ADDSI(W0, W1, 42)), 0x3100A820},
		{uint32(SUBI(W0, W1, 42)), 0x5100A820},
		{uint32(SUBSI(W0, W1, 42)), 0x7100A820},
		{uint32(ANDI(W0, W1, 0xFF)), 0x12001C20},
		{uint32(ANDSI(W0, W1, 0xFF)), 0x72001C20},
		{uint32(ORRI(W0, W1, 0xFF)), 0x32001C20},
		{uint32(EORI(W0, W1, 0xFF)), 0x52001C20},
		// 64-bit (sf=1)
		{uint32(ADD(X0, X1, X2)), 0x8B020020},
		{uint32(ADDS(X0, X1, X2)), 0xAB020020},
		{uint32(SUB(X0, X1, X2)), 0xCB020020},
		{uint32(SUBS(X0, X1, X2)), 0xEB020020},
		{uint32(AND(X0, X1, X2)), 0x8A020020},
		{uint32(ANDS(X0, X1, X2)), 0xEA020020},
		{uint32(ORR(X0, X1, X2)), 0xAA020020},
		{uint32(EOR(X0, X1, X2)), 0xCA020020},
		{uint32(ADDI(X0, X1, 42)), 0x9100A820},
		{uint32(ADDSI(X0, X1, 42)), 0xB100A820},
		{uint32(SUBI(X0, X1, 42)), 0xD100A820},
		{uint32(SUBSI(X0, X1, 42)), 0xF100A820},
		{uint32(ANDI(X0, X1, 0xFF)), 0x92401C20},
		{uint32(ANDSI(X0, X1, 0xFF)), 0xF2401C20},
		{uint32(ORRI(X0, X1, 0xFF)), 0xB2401C20},
		{uint32(EORI(X0, X1, 0xFF)), 0xD2401C20},
	}

	tests.Test(t)
}

// func TestPCRInstructions(t *testing.T) {
// 	tests := InstructionTests{
// 		{uint32(ADR(X0, 0)), 0x10000000},
// 		{uint32(ADR(X0, 1)), 0x30000000},
// 		{uint32(ADRP(X0, 0)), 0x90000000},
// 		{uint32(ADRP(X0, 1)), 0xB0000000},
// 	}

// 	tests.Test(t)
// }

// func TestLSRInstructions(t *testing.T) {
// 	tests := InstructionTests{
// 		{uint32(STR(W0, X1, X2)), 0xB8226820},
// 		{uint32(STR(X0, X1, X2)), 0xF8226820},
// 		{uint32(LDR(W0, X1, X2)), 0xB8626820},
// 		{uint32(LDR(X0, X1, X2)), 0xF8626820},
// 	}

// 	tests.Test(t)
// }
