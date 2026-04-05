package arm64

import (
	"fmt"
	"log"
	"strings"
)

// Instruction is the type constraint for all ARM64 instruction word types.
// Every concrete instruction type is a named uint32 with a String method
type Instruction interface {
	~uint32

	String() string
}

// Shared field sizes, masks and bit positions used across all instruction
// families
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

// getSF returns the `sf` bit of instruction as a raw uint8 (0 or 1).
// sf=1 indicates 64-bit (X) registers; sf=0 indicates 32-bit (W) registers
func getSF[I Instruction](inst I) uint8 {
	return get[uint8](inst, sfMask, sfPos)
}

// setSF encodes sf into the `sf` field of instruction
func setSF[I Instruction](inst I, sf uint8) I {
	return set(inst, sf, sfMask, sfPos)
}

// getRm extracts the `Rm` register field from instruction.
// Returns an X-register if sf=1, otherwise a W-register
func getRm[I Instruction](inst I) Register {
	rm := get[Register](inst, rmMask, rmPos)
	if getSF(inst) == 0 {
		return rm
	}

	return rm + X0
}

// setRm encodes rm into the `Rm` field of instruction.
// The register index is normalized to [0, 31] via modulo X0
func setRm[I Instruction](inst I, rm Register) I {
	return set(inst, rm%X0, rmMask, rmPos)
}

// getRn extracts the `Rn` register field from instruction.
// Returns an X-register if sf=1, otherwise a W-register
func getRn[I Instruction](inst I) Register {
	rn := get[Register](inst, rnMask, rnPos)
	if getSF(inst) == 0 {
		return rn
	}

	return rn + X0
}

// setRn encodes rn into the `Rn` field of instruction.
// The register index is normalized to [0, 31] via modulo X0
func setRn[I Instruction](inst I, rn Register) I {
	return set(inst, rn%X0, rnMask, rnPos)
}

// getRn extracts the `Rd` register field from instruction.
// Returns an X-register if sf=1, otherwise a W-register
func getRd[I Instruction](inst I) Register {
	rd := get[Register](inst, rdMask, rdPos)
	if getSF(inst) == 0 {
		return rd
	}

	return rd + X0
}

// setRn encodes rd into the `Rd` field of instruction.
// The register index is normalized to [0, 31] via modulo X0
func setRd[I Instruction](inst I, rd Register) I {
	return set(inst, rd%X0, rdMask, rdPos)
}

// operand is the type constraint for raw field values extracted from or written into
// instruction words
type operand interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

// get extracts a field from inst by shifting right by shift and masking with mask
func get[O, I operand](inst I, mask, shift uint32) O {
	return O((uint32(inst) >> shift) & mask)
}

// set writes `op` into instruction at the field defined by mask and shift, clearing any
// existing bits in that field first
func set[O, I operand](inst I, op O, mask, shift uint32) I {
	return I((uint32(inst) & ^(mask << shift)) | ((uint32(op) << shift) & (mask << shift)))
}

// DPInstruction is a 32-bit ARM64 Data Processing instruction word.
//
// Four sub-encodings share this type.
// The caller selects the appropriate accessors based on which sub-encoding
// the instruction uses:
//
// Plain Register:
//
//	31     30     29       28-24      23-22     21    20-16  15-10            9-5    4-0
//	+------+------+--------+----------+---------+-----+------+----------------+------+------+
//	|  sf  |  op  |  sign  |  opcode  |  00     |  0  |  Rm  |  000000        |  Rn  |  Rd  |
//	+------+------+--------+----------+---------+-----+------+----------------+------+------+
//
// Shifted Register:
//
//	31     30     29       28-24      23-22     21    20-16  15-10            9-5    4-0
//	+------+------+--------+----------+---------+-----+------+----------------+------+------+
//	|  sf  |  op  |  sign  |  opcode  |  shift  |  0  |  Rm  |  imm6          |  Rn  |  Rd  |
//	+------+------+--------+----------+---------+-----+------+----------------+------+------+
//
// Extended Register:
//
//	31     30     29       28-24      23-22     21    20-16  15-13   12-10    9-5    4-0
//	+------+------+--------+----------+---------+-----+------+-------+--------+------+------+
//	|  sf  |  op  |  sign  |  opcode  |  00     |  1  |  Rm  |  opt  |  imm3  |  Rn  |  Rd  |
//	+------+------+--------+----------+---------+-----+------+-------+--------+------+------+
//
// Immediate:
//
//	31     30     29       28-24      23    22        21-10                   9-5    4-0
//	+------+------+--------+----------+-----+---------+-----------------------+------+------+
//	|  sf  |  op  |  sign  |  opcode  |  0  |  shift  |  imm12                |  Rn  |  Rd  |
//	+------+------+--------+----------+-----+---------+-----------------------+------+------+
//
// Bitmask Immediate:
//
//	31     30     29       28-24      23    22        21-16       15-10       9-5    4-0
//	+------+------+--------+----------+-----+---------+-----------+-----------+------+------+
//	|  sf  |  op  |  sign  |  opcode  |  0  |  N      |  imms     |  immr     |  Rn  |  Rd  |
//	+------+------+--------+----------+-----+---------+-----------+-----------+------+------+
//
// Checked Pointer:
//
//	31     30     29       28-24      23-21           20-16  15-13   12-10    9-5    4-0
//	+------+------+--------+----------+---------------+------+-------+--------+------+------+
//	|  sf  |  op  |  sign  |  opcode  |  000          |  Rm  |  001  |  imm3  |  Rn  |  Rd  |
//	+------+------+--------+----------+---------------+------+-------+--------+------+------+
type DPInstruction uint32

// Field sizes, masks and positions for DPInstruction sub-encodings.
// Fields are shared where their bit positions coincide across sub-encodings
const (
	dpOpcodeSize uint32 = 5
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

// IsSF returns true if the instruction operates on 64-bit (X) registers
func (i DPInstruction) IsSF() bool {
	return getSF(i) == 1
}

// WithSF sets the `sf` bit and returns the updated instruction.
// Pass true for 64-bit (X) registers, false for 32-bit (W) registers
func (i DPInstruction) WithSF(flag bool) DPInstruction {
	var sf uint8
	if flag {
		sf = 1
	}

	return setSF(i, sf)
}

func (i DPInstruction) Opcode() DPOpcode {
	var mask uint32

	switch cat := get[uint8](i, dpOpcodeMask, dpOpcodePos); cat {
	case dpCatArithReg, dpCatLogicReg:
		mask = (sfMask << sfPos) | (dpShiftMask << dpShiftPos) |
			(dpExtMask << dpExtPos) | (rmMask << rmPos) |
			(dpImm6Mask << dpImm6Pos) | (rnMask << rnPos) |
			(rmMask << rmPos)

	case dpCatArithImm, dpCatLogicImm:
		mask = (sfMask << sfPos) | (dpShiftMask << dpShiftPos) |
			(dpImm12Mask << dpImm6Pos) | (rnMask << rnPos) |
			(rmMask << rmPos)

	case dpCatArithWCarry:
		mask = (sfMask << sfPos) | (dpShiftMask << dpShiftPos) |
			(rmMask << rmPos) | (dpImm3Mask << dpImm3Pos) |
			(rnMask << rnPos) | (rmMask << rmPos)
	}

	return DPOpcode(uint32(i) & ^mask)
}

// Immediate returns the 12-bit unsigned immediate
func (i DPInstruction) Immediate() uint16 {
	return get[uint16](i, dpImm12Mask, dpImm12Pos)
}

// WithImmediate encodes unsigned immediate into `imm12` field and
// returns the updated instruction
func (i DPInstruction) WithImmediate(imm uint16) DPInstruction {
	return set(i, imm, dpImm12Mask, dpImm12Pos)
}

// ImmShift decodes the `shift` bit and returns the shift kind and
// the predefined shift amount
func (i DPInstruction) ImmShift() (Shift, uint8) {
	if shift := get[uint8](i, dpShiftMask>>1, dpShiftPos); shift == 0x0 {
		return ShiftLSL, 0
	}

	return ShiftLSL, 0xC
}

// WithImmShift sets the `shift` bit and returns the updated instruction
func (i DPInstruction) WithImmShift(flag bool) DPInstruction {
	var shift uint8
	if flag {
		shift = 1
	}

	return set(i, shift, dpShiftMask>>1, dpShiftPos)
}

// Bitmask decodes the `N/imms/immr` fields and returns the 64-bit immediate
// value they represent. Valid only for logical immediate instructions
func (i DPInstruction) Bitmask() uint64 {
	n := get[uint8](i, dpShiftMask>>1, dpShiftPos)
	imms := get[uint8](i, dpImmsMask, dpImmsPos)
	immr := get[uint8](i, dpImmrMask, dpImmrPos)

	return decodeBitmask(n, imms, immr)
}

// WithBitmask encodes value as an ARM64 bitmask immediate into the `N/imms/immr`
// fields and returns the updated instruction. The `sf` bit must be set before
// calling this method as it determines whether to use 32 or 64-bit encoding
func (i DPInstruction) WithBitmask(bitmask uint64) DPInstruction {
	n, imms, immr := encodeBitmask(bitmask, i.IsSF())

	i = set(i, n, dpShiftMask>>1, dpShiftPos)
	i = set(i, imms, dpImmsMask, dpImmsPos)
	i = set(i, immr, dpImmrMask, dpImmrPos)

	return i
}

// Rd returns the `Rm` register field, adjusted for the `sf` bit
func (i DPInstruction) Rm() Register {
	return getRm(i)
}

// WithRm encodes rn into the `Rm` field and returns the updated instruction
func (i DPInstruction) WithRm(rm Register) DPInstruction {
	return setRm(i, rm)
}

// RmShift returns the shift kind (LSL/LSR/ASR) and shift amount for
// shifted-register instructions
func (i DPInstruction) RmShift() (Shift, uint8) {
	shift := get[Shift](i, dpShiftMask, dpShiftPos)
	amount := get[uint8](i, dpImm6Mask, dpImm6Pos)

	return shift, amount
}

// WithRmShift encodes a shifted-register operand into the instruction.
// Sets ext=0, the `shift`, and the shift amount into `imm6` field.
// Returns the updated instruction
func (i DPInstruction) WithRmShift(shift Shift, amount uint8) DPInstruction {
	i = set(i, uint8(0), dpExtMask, dpExtPos)
	i = set(i, uint8(shift), dpShiftMask, dpShiftPos)
	i = set(i, amount, dpImm6Mask, dpImm6Pos)

	return i
}

// RmExt returns the extend option and shift amount for extended-register
// instructions. Valid only for extended-register form instructions
func (i DPInstruction) RmExt() (Extension, uint8) {
	opt := get[Extension](i, dpOptMask, dpOptPos)
	amount := get[uint8](i, dpImm3Mask, dpImm3Pos)

	return opt, amount
}

// WithRmExt encodes an extended-register operand into the instruction.
// Sets ext=0, shift=0, the extend option into `opt`, and the shift amount bits
// into `imm3` field. Returns the updated instruction
func (i DPInstruction) WithRmExt(option Extension, amount uint8) DPInstruction {
	i = set(i, uint8(1), dpExtMask, dpExtPos)
	i = set(i, uint8(0), dpShiftMask, dpShiftPos)
	i = set(i, uint8(option), dpOptMask, dpOptPos)
	i = set(i, amount, dpImm3Mask, dpImm3Pos)

	return i
}

// Rd returns the `Rn` register field, adjusted for the `sf` bit
func (i DPInstruction) Rn() Register {
	return getRn(i)
}

// WithRn encodes rn into the `Rn` field and returns the updated instruction
func (i DPInstruction) WithRn(rn Register) DPInstruction {
	return setRn(i, rn)
}

// Rd returns the `Rd` register field, adjusted for the `sf` bit
func (i DPInstruction) Rd() Register {
	return getRd(i)
}

// WithRd encodes rd into the `Rd` field and returns the updated instruction
func (i DPInstruction) WithRd(rd Register) DPInstruction {
	return setRd(i, rd)
}

// String returns a human-readable ARM64 assembly representation of the
// instruction, selecting the correct operand format based on the opcode family
func (i DPInstruction) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%s %s, %s, ", i.Opcode(), i.Rd(), i.Rn())

	switch cat := get[uint8](i, dpOpcodeMask, dpOpcodePos); cat {
	case dpCatArithReg, dpCatLogicReg:
		shift := get[Shift](i, dpShiftMask, dpShiftPos)
		ext := get[uint8](i, dpExtMask, dpExtPos)

		// extended-register form
		if ext == 1 && shift == 0x0 {
			opt := get[Extension](i, dpOptMask, dpOptPos)
			amount := get[uint8](i, dpImm3Mask, dpImm3Pos)

			fmt.Fprintf(&b, "%s, %s #%#X", i.Rm(), opt, amount)
			break
		}

		// shifted-register form
		if amount := get[uint8](i, dpImm6Mask, dpImm6Pos); ext == 0 && amount != 0x0 {
			fmt.Fprintf(&b, "%s, %s #%#X", i.Rm(), shift, amount)
			break
		}

		// plain register form
		fmt.Fprintf(&b, "%s", i.Rm())

	case dpCatArithImm:
		// The `shift` bit indicates imm12 << 12 when set
		if shift := get[uint8](i, dpShiftMask>>1, dpShiftPos); shift == 0 {
			fmt.Fprintf(&b, "#%#X", i.Immediate())
			break
		}

		fmt.Fprintf(&b, "#%#X, lsl #12", i.Immediate())

	case dpCatLogicImm:
		fmt.Fprintf(&b, "#%#X", i.Bitmask())

	case dpCatArithWCarry:
		// plain register form
		fmt.Fprintf(&b, "%s", i.Rm())

	default:
		log.Fatalf("Unknown opcode category: Instruction(%032b)", uint32(i))
	}

	return b.String()
}
