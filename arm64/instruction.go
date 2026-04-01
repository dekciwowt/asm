package arm64

import (
	"fmt"
	"log"
	"math/bits"
)

// Instruction is a single 32-bit ARM64 instruction word.
// All field accessors operate directly on the bits the caller
// is responsible for using the right accessors for the instruction family
//
// Bit layout overview (field positions vary by family):
//
// Data Processing – Register:
//
//	31     30-24      23-22     21    20-16  15-10   9-5    4-0
//	+------+----------+---------+-----+------+-------+------+------+
//	|  sf  |  opcode  |  shift  |  N  |  Rm  |  imm  |  Rn  |  Rd  |
//	+------+----------+---------+-----+------+-------+------+------+
//
// Data Processing – Immediate:
//
//	31     30-24      23        22-10                9-5    4-0
//	+------+----------+---------+--------------------+------+------+
//	|  sf  |  opcode  |  shift  |  imm12             |  Rn  |  Rd  |
//	+------+----------+---------+--------------------+------+------+
//
// PC-relative:
//
//	31     30-29     28-24      23-5                        4-0
//	+------+---------+----------+---------------------------+------+
//	|  sf  |  immlo  |  opcode  |  immhi                    |  Rd  |
//	+------+---------+----------+---------------------------+------+
//
// Load/Store:
//
//	31-30  29-22       21  20-16  15-13   12  11-10  9-5    4-0
//	+------+-----------+---+------+-------+---+------+------+------+
//	| size |  opcode   | 1 |  Rm  |  opt  | S |  10  |  Rn  |  Rd  |
//	+------+-----------+---+------+-------+---+------+------+------+
type Instruction uint32

const (
	// Shared operands

	rdSize = 5
	rnSize = 5
	rmSize = 5
	sfSize = 1

	rdPos = 0
	rnPos = 5
	rmPos = 16
	sfPos = 31

	rdMask = (1 << rdSize) - 1
	rnMask = (1 << rnSize) - 1
	rmMask = (1 << rmSize) - 1
	sfMask = (1 << sfSize) - 1

	// Data Processing – Register

	dprImm6Size   = 6
	dprNSize      = 1
	dprShiftSize  = 2
	dprOpcodeSize = 7

	dprImm6Pos   = rnPos + rnSize
	dprNPos      = rmPos + rmSize
	dprShiftPos  = dprNPos + dprNSize
	dprOpcodePos = dprShiftPos + dprShiftSize

	dprImm6Mask   = (1 << dprImm6Size) - 1
	dprNMask      = (1 << dprNSize) - 1
	dprShiftMask  = (1 << dprShiftSize) - 1
	dprOpcodeMask = (1 << dprOpcodeSize) - 1

	// Data Processing – Immediate

	dpiImm12Size  = 12
	dpiNSize      = 1
	dpiShiftSize  = 1
	dpiOpcodeSize = 7

	dpiImm12Pos  = rnPos + rnSize
	dpiNPos      = dpiImm12Pos + dpiImm12Size
	dpiShiftPos  = dpiNPos + dpiNSize
	dpiOpcodePos = dpiShiftPos + dpiShiftSize

	dpiImm12Mask  = (1 << dpiImm12Size) - 1
	dpiNMask      = (1 << dpiNSize) - 1
	dpiShiftMask  = (1 << dpiShiftSize) - 1
	dpiOpcodeMask = (1 << dpiOpcodeSize) - 1

	// PC-relative

	pcrImmhiSize  = 19
	pcrOpcodeSize = 5
	pcrImmloSize  = 2

	pcrImmhiPos  = rdPos + rdSize
	pcrOpcodePos = pcrImmhiPos + pcrImmhiSize
	pcrImmloPos  = pcrOpcodePos + pcrOpcodeSize

	pcrImmhiMask  = (1 << pcrImmhiSize) - 1
	pcrOpcodeMask = (1 << pcrOpcodeSize) - 1
	pcrImmloMask  = (1 << pcrImmloSize) - 1

	// Load/Store

	lsrShiftSize  = 1
	lsrOptionSize = 3
	lsrOpcodeSize = 8
	lsrSizeSize   = 2

	lsrShiftPos  = rnPos + rnSize + 2
	lsrOptionPos = lsrShiftPos + lsrShiftSize
	lsrOpcodePos = rmPos + rmSize + 1
	lsrSizePos   = lsrOpcodePos + lsrOpcodeSize

	lsrShiftMask  = (1 << lsrShiftSize) - 1
	lsrOptionMask = (1 << lsrOptionSize) - 1
	lsrOpcodeMask = (1 << lsrOpcodeSize) - 1
	lsrSizeMask   = (1 << lsrSizeSize) - 1
)

func (i Instruction) SF() uint8 {
	return decodeOperand[uint8](i, sfMask, sfPos)
}

func (i Instruction) WithSF(sf uint8) Instruction {
	return encodeOperand(i, sf, sfMask, sfPos)
}

func (i Instruction) Family() Family {
	op7 := uint8((uint32(i) >> 24) & 0x7F)
	if family := resolveFamily(op7); family != FamilyUND {
		return family
	}

	op8 := uint8((uint32(i) >> 22) & 0xFF)
	return resolveFamily(op8)
}

func (i Instruction) Opcode() Opcode {
	switch i.Family() {
	case FamilyDPR:
		return decodeOperand[Opcode](i, dprOpcodeMask, dprOpcodePos)

	case FamilyDPI:
		return decodeOperand[Opcode](i, dpiOpcodeMask, dpiOpcodePos)

	case FamilyPCR:
		return decodeOperand[Opcode](i, pcrOpcodeMask, pcrOpcodePos)

	case FamilyLSR:
		return decodeOperand[Opcode](i, lsrOpcodeMask, lsrOpcodePos)

	default:
		return OpUND
	}
}

func (i Instruction) WithOpcode(opcode Opcode) Instruction {
	switch resolveFamily(uint8(opcode)) {
	case FamilyDPR:
		return encodeOperand(i, uint8(opcode), dprOpcodeMask, dprOpcodePos)

	case FamilyDPI:
		return encodeOperand(i, uint8(opcode), dpiOpcodeMask, dpiOpcodePos)

	case FamilyPCR:
		return encodeOperand(i, uint8(opcode), pcrOpcodeMask, pcrOpcodePos)

	case FamilyLSR:
		const lsrDefaults = Instruction(0x200800)
		return encodeOperand(i|lsrDefaults, uint8(opcode), lsrOpcodeMask, lsrOpcodePos)

	default:
		log.Fatalf("unknown opcode: %s", opcode)
		return Instruction(0)
	}
}

func (i Instruction) Shift() uint8 {
	switch i.Family() {
	case FamilyDPR:
		return decodeOperand[uint8](i, dprShiftMask, dprShiftPos)

	case FamilyDPI:
		return decodeOperand[uint8](i, dpiShiftMask, dpiShiftPos)

	case FamilyLSR:
		return decodeOperand[uint8](i, lsrShiftMask, lsrShiftPos)
	}

	return 0
}

func (i Instruction) WithShift(shift uint8) Instruction {
	switch i.Family() {
	case FamilyDPR:
		return encodeOperand(i, shift, dprShiftMask, dprShiftPos)

	case FamilyDPI:
		return encodeOperand(i, shift, dpiShiftMask, dpiShiftPos)

	case FamilyLSR:
		return encodeOperand(i, shift, lsrShiftMask, lsrShiftPos)
	}

	return i
}

var opsWithXdOnly = map[Opcode]bool{
	OpADR: true,
}

func (i Instruction) Rd() Register {
	rd := decodeOperand[Register](i, rdMask, rdPos)
	if _, ok := opsWithXdOnly[i.Opcode()]; ok || i.SF() == 1 {
		return rd + X0
	}

	return rd
}

func (i Instruction) WithRd(rd Register) Instruction {
	return encodeOperand(i, rd%X0, rdMask, rdPos)
}

var opsWithXnOnly = map[Opcode]bool{
	OpSTR: true,
}

func (i Instruction) Rn() Register {
	rn := decodeOperand[Register](i, rnMask, rnPos)
	if _, ok := opsWithXnOnly[i.Opcode()]; ok || i.SF() == 1 {
		return rn + X0
	}

	return rn
}

func (i Instruction) WithRn(rn Register) Instruction {
	return encodeOperand(i, rn%X0, rnMask, rnPos)
}

var opsWithXmOnly = map[Opcode]bool{
	OpSTR: true,
}

func (i Instruction) Rm() Register {
	rm := decodeOperand[Register](i, rmMask, rmPos)
	if _, ok := opsWithXmOnly[i.Opcode()]; ok || i.SF() == 1 {
		return rm + X0
	}

	return rm
}

func (i Instruction) WithRm(rm Register) Instruction {
	return encodeOperand(i, rm%X0, rmMask, rmPos)
}

func (i Instruction) String() string {
	switch i.Family() {
	case FamilyDPR:
		return fmt.Sprintf("%s %s, %s, %s", i.Opcode(), i.Rd(), i.Rn(), i.Rm())

	case FamilyDPI:
		switch i.Opcode() {
		case OpADDI, OpADDSI, OpSUBI, OpSUBSI:
			return fmt.Sprintf("%s %s, %s, #%#X", i.Opcode(), i.Rd(), i.Rn(), i.Imm12())
		}

		return fmt.Sprintf("%s %s, %s, #%#X", i.Opcode(), i.Rd(), i.Rn(), i.Bitmask())

	case FamilyPCR:
		if i.SF() == 1 {
			return fmt.Sprintf("ADRP %s, #%#X", i.Rd(), i.Imm21())
		}

		return fmt.Sprintf("ADR %s, #%#X", i.Rd(), i.Imm21())

	case FamilyLSR:
		return fmt.Sprintf("%s %s, [%s, %s]", i.Opcode(), i.Rd(), i.Rn(), i.Rm())

	default:
		return fmt.Sprintf("Instruction(%#032b)", uint32(i))
	}
}

// Data Processing – Register operands

func (i Instruction) N() uint8 {
	return decodeOperand[uint8](i, dprNMask, dprNPos)
}

func (i Instruction) WithN(n uint8) Instruction {
	return encodeOperand(i, n, dprNMask, dprNPos)
}

func (i Instruction) Imm6() uint8 {
	return decodeOperand[uint8](i, dprImm6Mask, dprImm6Pos)
}

func (i Instruction) WithImm6(imm uint8) Instruction {
	return encodeOperand(i, imm, dprImm6Mask, dprImm6Pos)
}

// Data Processing – Immediate operands

func (i Instruction) Imm12() uint16 {
	return decodeOperand[uint16](i, dpiImm12Mask, dpiImm12Pos)
}

func (i Instruction) WithImm12(imm uint16) Instruction {
	return encodeOperand(i, imm, dpiImm12Mask, dpiImm12Pos)
}

func (i Instruction) Bitmask() uint64 {
	n := decodeOperand[uint8](i, dpiNMask, dpiNPos)
	imm12 := decodeOperand[uint16](i, dpiImm12Mask, dpiImm12Pos)
	return decodeBitmask(n, uint8(imm12>>6), uint8(imm12&0x3F))
}

func (i Instruction) WithBitmask(value uint64) Instruction {
	n, immr, imms := encodeBitmask(value, i.SF() == 1)
	imm12 := uint16(immr)<<6 | uint16(imms)
	i = encodeOperand(i, n, dpiNMask, dpiNPos)
	i = encodeOperand(i, imm12, dpiImm12Mask, dpiImm12Pos)
	return i
}

// PC-relative operands

func (i Instruction) Imm21() int32 {
	immhi := decodeOperand[uint32](i, pcrImmhiMask, pcrImmhiPos)
	immlo := decodeOperand[uint8](i, pcrImmloMask, pcrImmloPos)
	return (int32(immhi<<2) | int32(immlo)) << 11 >> 11
}

func (i Instruction) WithImm21(imm21 int32) Instruction {
	imm := uint32(imm21) & 0x1FFFFF
	i = encodeOperand(i, imm>>2, pcrImmhiMask, pcrImmhiPos)
	i = encodeOperand(i, uint8(imm&0x3), pcrImmloMask, pcrImmloPos)
	return i
}

// Load/Save operands

type Size uint8

const /* Size */ (
	SizeByte  Size = 0x0
	SizeHalf  Size = 0x1
	SizeWord  Size = 0x2
	SizeDWord Size = 0x3
)

func (i Instruction) Size() Size {
	return decodeOperand[Size](i, lsrSizeMask, lsrSizePos)
}

func (i Instruction) WithSize(size Size) Instruction {
	return encodeOperand(i, uint8(size), lsrSizeMask, lsrSizePos)
}

type Option uint8

const /* Option */ (
	OptUXTB Option = 0x0
	OptUXTH Option = 0x1
	OptUXTW Option = 0x2
	OptLSL  Option = 0x3
	OptSXTB Option = 0x4
	OptSXTH Option = 0x5
	OptSXTW Option = 0x6
	OptSXTX Option = 0x7
)

var optNames = map[Option]string{
	OptUXTB: "UXTB",
	OptUXTH: "UXTH",
	OptUXTW: "UXTW",
	OptLSL:  "LSL",
	OptSXTB: "SXTB",
	OptSXTH: "SXTH",
	OptSXTW: "SXTW",
	OptSXTX: "SXTX",
}

func (o Option) String() string {
	if name, ok := optNames[o]; ok {
		return name
	}

	return fmt.Sprintf("Option(%#03b)", uint8(o))
}

func (i Instruction) Option() Option {
	return decodeOperand[Option](i, lsrOptionMask, lsrOptionPos)
}

func (i Instruction) WithOption(ext Option) Instruction {
	return encodeOperand(i, uint8(ext), lsrOptionMask, lsrOptionPos)
}

// Encoding/Decoding functions

func encodeOperand[V ~uint8 | ~uint16 | ~uint32](inst Instruction, val V, mask, shift uint32) Instruction {
	return Instruction((uint32(inst) & ^(mask << shift)) | ((uint32(val) << shift) & (mask << shift)))
}

func decodeOperand[V ~uint8 | ~uint16 | ~uint32](inst Instruction, mask, shift uint32) V {
	return V((uint32(inst) >> shift) & mask)
}

func encodeBitmask(value uint64, isExtMode bool) (n, immr, imms uint8) {
	if value == 0 || (isExtMode && value == ^uint64(0)) || (!isExtMode && value == 0xFFFFFFFF) {
		return
	}

	if !isExtMode {
		value32 := uint32(value)
		value = uint64(value32) | uint64(value32)<<32
	}

	elementSize := uint64(64)
	for size := uint64(2); size <= 32; size *= 2 {
		if bits.RotateLeft64(value, int(size)) == value {
			elementSize = size
			break
		}
	}

	elementMask := ^uint64(0)
	if elementSize != 64 {
		elementMask = (uint64(1) << elementSize) - 1
	}

	element := value & elementMask
	trailingZeros := uint64(bits.TrailingZeros64(element))
	rotated := (element>>trailingZeros | element<<(elementSize-trailingZeros)) & elementMask
	trailingOnes := bits.TrailingZeros64(^rotated & elementMask)

	if rotated != (uint64(1)<<trailingOnes)-1 {
		return
	}

	n, imms = 1, uint8(trailingOnes-1)
	if elementSize != 64 {
		log2s := bits.Len(uint(elementSize)) - 1
		tag := uint8((0x3F<<log2s)&0x3F) & ^uint8(1<<log2s)
		n, imms = 0, tag|uint8(trailingOnes-1)
	}

	immr = uint8(((elementSize - trailingZeros) % elementSize) & 0x3F)
	return
}

func decodeBitmask(n, immr, imms uint8) uint64 {
	nimms := (uint8(n) << 6) | (imms & 0x3F)

	len := 7 - bits.LeadingZeros8(^nimms)
	if len < 1 {
		return 0
	}

	elementSize := uint8(1) << len
	immsMask := uint8(elementSize - 1)

	onesCount := uint8(imms&immsMask) + 1
	if onesCount == elementSize {
		return 0
	}

	element := ^uint64(0)
	if onesCount != 64 {
		element = (uint64(1) << onesCount) - 1
	}

	if rotation := immr & (elementSize - 1); rotation > 0 {
		mask := (uint64(1) << elementSize) - 1
		element = ((element << rotation) | (element >> (elementSize - rotation))) & mask
	}

	value := element
	for size := elementSize; size < 64; size *= 2 {
		value |= value << size
	}

	return value
}
