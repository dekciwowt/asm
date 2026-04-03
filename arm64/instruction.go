package arm64

import (
	"fmt"
	"log"
	"strings"
)

// DPOpcode represents Data Processing opcode encoded as 7-bit wide constant
//
// Binary layout:
//
// 7-6     5-0
// +------+----------+
// |  op  |  opcode  |
// +------+----------+
type DPOpcode uint8

const /* DPOpcode */ (
	OpADD   DPOpcode = 0x0B // op=00 opcode=01011
	OpADDS  DPOpcode = 0x2B // op=01 opcode=01011
	OpSUB   DPOpcode = 0x4B // op=10 opcode=01011
	OpSUBS  DPOpcode = 0x6B // op=11 opcode=01011
	OpAND   DPOpcode = 0x0A // op=00 opcode=01010
	OpORR   DPOpcode = 0x2A // op=01 opcode=01010
	OpEOR   DPOpcode = 0x4A // op=10 opcode=01010
	OpANDS  DPOpcode = 0x6A // op=11 opcode=01010
	OpADDI  DPOpcode = 0x11 // op=00 opcode=10001
	OpADDSI DPOpcode = 0x31 // op=01 opcode=10001
	OpSUBI  DPOpcode = 0x51 // op=10 opcode=10001
	OpSUBSI DPOpcode = 0x71 // op=11 opcode=10001
	OpANDI  DPOpcode = 0x12 // op=00 opcode=10010
	OpANDSI DPOpcode = 0x72 // op=11 opcode=10010
	OpORRI  DPOpcode = 0x32 // op=01 opcode=10010
	OpEORI  DPOpcode = 0x52 // op=10 opcode=10010
)

var dpOpcodes = map[DPOpcode]string{
	OpADD:   "ADD",
	OpADDS:  "ADDS",
	OpSUB:   "SUB",
	OpSUBS:  "SUBS",
	OpAND:   "AND",
	OpORR:   "ORR",
	OpEOR:   "EOR",
	OpANDS:  "ANDS",
	OpADDI:  "ADD",
	OpADDSI: "ADDS",
	OpSUBI:  "SUB",
	OpSUBSI: "SUBS",
	OpANDI:  "AND",
	OpANDSI: "ANDS",
	OpORRI:  "ORR",
	OpEORI:  "EOR",
}

func (o DPOpcode) String() string {
	if name, ok := dpOpcodes[o]; ok {
		return name
	}

	return fmt.Sprintf("DPOpcode(%07b)", o)
}

type Shift uint8

const (
	ShiftLSL Shift = 0x0
	ShiftLSR Shift = 0x1
	ShiftASR Shift = 0x2
)

var shifts = map[Shift]string{
	ShiftLSL: "lsl",
	ShiftLSR: "lsr",
	ShiftASR: "asr",
}

func (s Shift) String() string {
	if name, ok := shifts[s]; ok {
		return name
	}

	return fmt.Sprintf("Shift(%#03X)", uint8(s))
}

type Extension uint8

const (
	ExtUXTB Extension = 0x0
	ExtUXTH Extension = 0x1
	ExtUXTW Extension = 0x2
	ExtUXTX Extension = 0x3
	ExtSXTB Extension = 0x4
	ExtSXTH Extension = 0x5
	ExtSXTW Extension = 0x6
	ExtSXTX Extension = 0x7
)

var exts = map[Extension]string{
	ExtUXTB: "uxtb",
	ExtUXTH: "uxth",
	ExtUXTW: "uxtw",
	ExtUXTX: "uxtx",
	ExtSXTB: "sxtb",
	ExtSXTH: "sxth",
	ExtSXTW: "sxtw",
	ExtSXTX: "sxtx",
}

func (e Extension) String() string {
	if name, ok := exts[e]; ok {
		return name
	}

	return fmt.Sprintf("Extension(%#03X)", uint8(e))
}

type Instruction interface {
	~uint32

	String() string
}

const (
	rdSize uint32 = 5
	rnSize uint32 = 5
	rmSize uint32 = 5
	sfSize uint32 = 1

	rdMask uint32 = (1 << rdSize) - 1
	rnMask uint32 = (1 << rnSize) - 1
	rmMask uint32 = (1 << rmSize) - 1
	sfMask uint32 = (1 << sfSize) - 1

	rdPos uint32 = 0
	rnPos uint32 = 5
	rmPos uint32 = 16
	sfPos uint32 = 31
)

func getSF[I Instruction](inst I) uint8 {
	return get[uint8](inst, sfMask, sfPos)
}

func setSF[I Instruction](inst I, sf uint8) I {
	return set(inst, sf, sfMask, sfPos)
}

func getRm[I Instruction](inst I) Register {
	rm := get[Register](inst, rmMask, rmPos)
	if getSF(inst) == 0 {
		return rm
	}

	return rm + X0
}

func setRm[I Instruction](inst I, rm Register) I {
	return set(inst, rm%X0, rmMask, rmPos)
}

func getRn[I Instruction](inst I) Register {
	rn := get[Register](inst, rnMask, rnPos)
	if getSF(inst) == 0 {
		return rn
	}

	return rn + X0
}

func setRn[I Instruction](inst I, rn Register) I {
	return set(inst, rn%X0, rnMask, rnPos)
}

func getRd[I Instruction](inst I) Register {
	rd := get[Register](inst, rdMask, rdPos)
	if getSF(inst) == 0 {
		return rd
	}

	return rd + X0
}

func setRd[I Instruction](inst I, rd Register) I {
	return set(inst, rd%X0, rdMask, rdPos)
}

type Operand interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

func get[O Operand, I Instruction](inst I, mask, shift uint32) O {
	return O((uint32(inst) >> shift) & mask)
}

func set[O Operand, I Instruction](inst I, op O, mask, shift uint32) I {
	return I((uint32(inst) & ^(mask << shift)) | ((uint32(op) << shift) & (mask << shift)))
}

// With Extended Register:
//
//	31     30-24      23-21           20-16  15-13   12-10    9-5    4-0
//	+------+----------+---------------+------+-------+--------+------+------+
//	|  sf  |  opcode  |  001          |  Rm  |  opt  |  imm3  |  Rn  |  Rd  |
//	+------+----------+---------------+------+-------+--------+------+------+
//
// With Immediate:
//
//	31     30-24      23    22        21-10                   9-5    4-0
//	+------+----------+-----+---------+-----------------------+------+------+
//	|  sf  |  opcode  |  0  |  shift  |  imm12                |  Rn  |  Rd  |
//	+------+----------+-----+---------+-----------------------+------+------+
//
// With Bitmask:
//
//	31     30-24      23    22        21-16       15-10       9-5    4-0
//	+------+----------+-----+---------+-----------+-----------+------+------+
//	|  sf  |  opcode  |  0  |  shift  |  imms     |  immr     |  Rn  |  Rd  |
//	+------+----------+-----+---------+-----------+-----------+------+------+
//
// With Shifted Register:
//
//	31     30-24      23-22     21    20-16  15-10            9-5    4-0
//	+------+----------+---------+-----+------+----------------+------+------+
//	|  sf  |  opcode  |  shift  |  0  |  Rm  |  imm6          |  Rn  |  Rd  |
//	+------+----------+---------+-----+------+----------------+------+------+
type DPInstruction uint32

const (
	dpOpcodeSize uint32 = 7
	dpShiftSize  uint32 = 2
	dpExtSize    uint32 = 1
	dpOptSize    uint32 = 3
	dpImmsSize   uint32 = 6
	dpImmrSize   uint32 = 6
	dpImm3Size   uint32 = 3
	dpImm6Size   uint32 = 6
	dpImm12Size  uint32 = 12

	dpOpcodeMask uint32 = (1 << dpOpcodeSize) - 1
	dpShiftMask  uint32 = (1 << dpShiftSize) - 1
	dpExtMask    uint32 = (1 << dpExtSize) - 1
	dpOptMask    uint32 = (1 << dpOptSize) - 1
	dpImmsMask   uint32 = (1 << dpImmsSize) - 1
	dpImmrMask   uint32 = (1 << dpImmrSize) - 1
	dpImm3Mask   uint32 = (1 << dpImm3Size) - 1
	dpImm6Mask   uint32 = (1 << dpImm6Size) - 1
	dpImm12Mask  uint32 = (1 << dpImm12Size) - 1

	dpOpcodePos uint32 = 24
	dpShiftPos  uint32 = 22
	dpExtPos    uint32 = 21
	dpOptPos    uint32 = 13
	dpImmsPos   uint32 = 16
	dpImmrPos   uint32 = 10
	dpImm3Pos   uint32 = 10
	dpImm6Pos   uint32 = 10
	dpImm12Pos  uint32 = 10
)

func (i DPInstruction) HasSF() bool {
	return getSF(i) == 1
}

// WithSF sets the `sixty-four` flag and returns new Data Processing instruction
func (i DPInstruction) WithSF(flag bool) DPInstruction {
	var sf uint8
	if flag {
		sf = 1
	}

	return setSF(i, sf)
}

func (i DPInstruction) Opcode() DPOpcode {
	return get[DPOpcode](i, dpOpcodeMask, dpOpcodePos)
}

func (i DPInstruction) WithOpcode(opcode DPOpcode) DPInstruction {
	return set(i, uint8(opcode), dpOpcodeMask, dpOpcodePos)
}

func (i DPInstruction) Shift() (Shift, uint8) {
	shift := get[Shift](i, dpShiftMask, dpShiftPos)
	ext := get[uint8](i, dpExtMask, dpExtPos)
	if ext == 1 && shift == 0x0 {
		log.Fatalf("This instruction does not contain any shift operands")
	}

	amount := get[uint8](i, dpImm6Mask, dpImm6Pos)

	return shift, amount
}

func (i DPInstruction) WithShift(shift Shift, amount uint8) DPInstruction {
	i = set(i, uint8(0), dpExtMask, dpExtPos)
	i = set(i, uint8(shift), dpShiftMask, dpShiftPos)
	i = set(i, amount, dpImm6Mask, dpImm6Pos)
	return i
}

func (i DPInstruction) Extension() (Extension, uint8) {
	shift := get[Shift](i, dpShiftMask, dpShiftPos)
	ext := get[uint8](i, dpExtMask, dpExtPos)
	if ext == 0 || shift != 0x0 {
		log.Fatalf("This instruction does not contain any extension operands")
	}

	opt := get[Extension](i, dpOptMask, dpOptPos)
	amount := get[uint8](i, dpImm3Mask, dpImm3Pos)

	return opt, amount
}

func (i DPInstruction) WithExtension(option Extension, amount uint8) DPInstruction {
	i = set(i, uint8(1), dpExtMask, dpExtPos)
	i = set(i, uint8(0x0), dpShiftMask, dpShiftPos)
	i = set(i, uint8(option), dpOptMask, dpOptPos)
	i = set(i, amount, dpImm3Mask, dpImm3Pos)
	return i
}

func (i DPInstruction) Immediate() uint16 {
	return get[uint16](i, dpImm12Mask, dpImm12Pos)
}

func (i DPInstruction) WithImmediate(imm uint16) DPInstruction {
	return set(i, imm, dpImm12Mask, dpImm12Pos)
}

func (i DPInstruction) Bitmask() uint64 {
	n := get[uint8](i, dpShiftMask>>1, dpShiftPos)
	imms := get[uint8](i, dpImmsMask, dpImmsPos)
	immr := get[uint8](i, dpImmrMask, dpImmrPos)
	return decodeBitmask(n, imms, immr)
}

func (i DPInstruction) WithBitmask(bitmask uint64) DPInstruction {
	n, imms, immr := encodeBitmask(bitmask, i.HasSF())
	i = set(i, n, dpShiftMask>>1, dpShiftPos)
	i = set(i, imms, dpImmsMask, dpImmsPos)
	i = set(i, immr, dpImmrMask, dpImmrPos)
	return i
}

func (i DPInstruction) Rm() Register {
	return getRm(i)
}

func (i DPInstruction) WithRm(rm Register) DPInstruction {
	return setRm(i, rm)
}

func (i DPInstruction) Rn() Register {
	return getRn(i)
}

func (i DPInstruction) WithRn(rn Register) DPInstruction {
	return setRn(i, rn)
}

func (i DPInstruction) Rd() Register {
	return getRd(i)
}

func (i DPInstruction) WithRd(rd Register) DPInstruction {
	return setRd(i, rd)
}

func (i DPInstruction) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%s %s, %s, ", i.Opcode(), i.Rd(), i.Rn())

	switch opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode {
	case 0xA, 0xB: // Operations on registers
		shift := get[Shift](i, dpShiftMask, dpShiftPos)
		ext := get[uint8](i, dpExtMask, dpExtPos)

		// Extended register
		if ext == 1 && shift == 0x0 {
			opt := get[Extension](i, dpOptMask, dpOptPos)
			amount := get[uint8](i, dpImm3Mask, dpImm3Pos)

			fmt.Fprintf(&b, "%s, %s #%#X", i.Rm(), opt, amount)

			break
		}

		// Shifted register
		if ext == 0 && shift != 0x0 {
			amount := get[uint8](i, dpImm6Mask, dpImm6Pos)

			fmt.Fprintf(&b, "%s, %s #%#X", i.Rm(), shift, amount)

			break
		}

		fmt.Fprintf(&b, "%s", i.Rm())

	case 0x11: // Operations on immediate
		shift := get[uint8](i, dpShiftMask>>1, dpShiftPos)
		if shift == 0 {
			fmt.Fprintf(&b, "#%#X", i.Immediate())
			break
		}

		fmt.Fprintf(&b, "#%#X, lsl #12", i.Immediate())

	case 0x12: // Operations on bitmask
		fmt.Fprintf(&b, "#%#X", i.Bitmask())
	}

	return b.String()
}
