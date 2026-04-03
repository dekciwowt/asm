package arm64

import (
	"fmt"
	"log"
	"strings"
)

// DPOpcode represents a 7-bit Data Processing opcode.
//
// The opcode is split into two sub-fields:
//
//	7-6     5-0
//	+------+----------+
//	|  op  |  opcode  |
//	+------+----------+
//
// The op field (bits 6–5) encodes the operation variant:
//
//	00 — ADD  / AND  / ADDI  / ANDI
//	01 — ADDS / ORR  / ADDSI / ORRI
//	10 — SUB  / EOR  / SUBI  / EORI
//	11 — SUBS / ANDS / SUBSI / ANDSI
//
// The opcode field (bits 4–0) identifies the instruction family:
//
//	01010 — logical register  (AND, ORR, EOR, ANDS)
//	01011 — arithmetic register (ADD, ADDS, SUB, SUBS)
//	10001 — arithmetic immediate (ADDI, ADDSI, SUBI, SUBSI)
//	10010 — logical immediate (ANDI, ORRI, EORI, ANDSI)
type DPOpcode uint8

const /* DPOpcode */ (
	OpADD   DPOpcode = 0x0B // op=00 opcode=01011 — add
	OpADDS  DPOpcode = 0x2B // op=01 opcode=01011 — add, set flags
	OpSUB   DPOpcode = 0x4B // op=10 opcode=01011 — subtract
	OpSUBS  DPOpcode = 0x6B // op=11 opcode=01011 — subtract, set flags
	OpAND   DPOpcode = 0x0A // op=00 opcode=01010 — bitwise AND
	OpORR   DPOpcode = 0x2A // op=01 opcode=01010 — bitwise OR
	OpEOR   DPOpcode = 0x4A // op=10 opcode=01010 — bitwise XOR
	OpANDS  DPOpcode = 0x6A // op=11 opcode=01010 — bitwise AND, set flags
	OpADDI  DPOpcode = 0x11 // op=00 opcode=10001 — add immediate
	OpADDSI DPOpcode = 0x31 // op=01 opcode=10001 — add immediate, set flags
	OpSUBI  DPOpcode = 0x51 // op=10 opcode=10001 — subtract immediate
	OpSUBSI DPOpcode = 0x71 // op=11 opcode=10001 — subtract immediate, set flags
	OpANDI  DPOpcode = 0x12 // op=00 opcode=10010 — bitwise AND immediate
	OpANDSI DPOpcode = 0x72 // op=11 opcode=10010 — bitwise AND immediate, set flags
	OpORRI  DPOpcode = 0x32 // op=01 opcode=10010 — bitwise OR immediate
	OpEORI  DPOpcode = 0x52 // op=10 opcode=10010 — bitwise XOR immediate
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

// String returns the canonical ARM64 mnemonic for the opcode.
// Returns a formatted fallback string if the opcode is not recognized
func (o DPOpcode) String() string {
	if name, ok := dpOpcodes[o]; ok {
		return name
	}

	return fmt.Sprintf("DPOpcode(%07b)", o)
}

// Shift represents the shift kind used in shifted-register instructions
type Shift uint8

const (
	ShiftLSL Shift = 0x0 // logical shift left
	ShiftLSR Shift = 0x1 // logical shift right
	ShiftASR Shift = 0x2 // arithmetic shift right
)

var shifts = map[Shift]string{
	ShiftLSL: "lsl",
	ShiftLSR: "lsr",
	ShiftASR: "asr",
}

// String returns the lowercase ARM64 shift mnemonic (lsl, lsr, asr)
// Returns a formatted fallback string if the shift is not recognized
func (s Shift) String() string {
	if name, ok := shifts[s]; ok {
		return name
	}

	return fmt.Sprintf("Shift(%#03X)", uint8(s))
}

// Extension represents the extend/shift option used in extended-register
// and load/store instructions
//
// The UXTW and LSL options are equivalent for 32-bit operands
// The UXTX and LSL options are equivalent for 64-bit operands
type Extension uint8

const (
	ExtUXTB Extension = 0x0 // 000 — zero-extend byte     (8 → 64)
	ExtUXTH Extension = 0x1 // 001 — zero-extend halfword (16 → 64)
	ExtUXTW Extension = 0x2 // 010 — zero-extend word     (32 → 64)
	ExtUXTX Extension = 0x3 // 011 — zero-extend dword 		(64 → 64)
	ExtSXTB Extension = 0x4 // 100 — sign-extend byte     (8 → 64)
	ExtSXTH Extension = 0x5 // 101 — sign-extend halfword (16 → 64)
	ExtSXTW Extension = 0x6 // 110 — sign-extend word     (32 → 64)
	ExtSXTX Extension = 0x7 // 111 — sign-extend dword		(64 → 64)
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

// String returns the lowercase ARM64 extension mnemonic (uxtb, sxtw, etc).
// Returns a formatted fallback string if the extension is not recognized
func (e Extension) String() string {
	if name, ok := exts[e]; ok {
		return name
	}

	return fmt.Sprintf("Extension(%#03X)", uint8(e))
}

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

// Operand is the type constraint for raw field values extracted from or
// written into instruction words
type Operand interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

// get extracts a field from inst by shifting right by shift and masking
// with mask
func get[O Operand, I Instruction](inst I, mask, shift uint32) O {
	return O((uint32(inst) >> shift) & mask)
}

// set writes `op` into instruction at the field defined by mask and shift,
// clearing any existing bits in that field first
func set[O Operand, I Instruction](inst I, op O, mask, shift uint32) I {
	return I((uint32(inst) & ^(mask << shift)) | ((uint32(op) << shift) & (mask << shift)))
}

// DPInstruction is a 32-bit ARM64 Data Processing instruction word.
//
// Four sub-encodings share this type — the caller selects the appropriate
// accessors based on which sub-encoding the instruction uses:
//
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
// With Bitmask Immediate:
//
//	31     30-24      23    22        21-16       15-10       9-5    4-0
//	+------+----------+-----+---------+-----------+-----------+------+------+
//	|  sf  |  opcode  |  0  |  N      |  imms     |  immr     |  Rn  |  Rd  |
//	+------+----------+-----+---------+-----------+-----------+------+------+
//
// With Shifted Register:
//
//	31     30-24      23-22     21    20-16  15-10            9-5    4-0
//	+------+----------+---------+-----+------+----------------+------+------+
//	|  sf  |  opcode  |  shift  |  0  |  Rm  |  imm6          |  Rn  |  Rd  |
//	+------+----------+---------+-----+------+----------------+------+------+
type DPInstruction uint32

// Field sizes, masks and positions for DPInstruction sub-encodings.
// Fields are shared where their bit positions coincide across sub-encodings
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

// Opcode returns the 7-bit Data Processing opcode from bits 30–24
func (i DPInstruction) Opcode() DPOpcode {
	return get[DPOpcode](i, dpOpcodeMask, dpOpcodePos)
}

// WithOpcode encodes opcode into `opcode` field.
// Returns the updated instruction
func (i DPInstruction) WithOpcode(opcode DPOpcode) DPInstruction {
	return set(i, uint8(opcode), dpOpcodeMask, dpOpcodePos)
}

// Shift returns the shift kind (LSL/LSR/ASR) and shift amount for
// shifted-register instructions. For non shifted-register form
// instruction method fatals
func (i DPInstruction) Shift() (Shift, uint8) {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0xA && opcode != 0xB {
		log.Fatalf("cannot get shift due to non register instruction")
	}

	shift := get[Shift](i, dpShiftMask, dpShiftPos)
	ext := get[uint8](i, dpExtMask, dpExtPos)
	if ext == 1 && shift == 0x0 {
		log.Fatalf("cannot get shift due to non shifted register instruction")
	}

	amount := get[uint8](i, dpImm6Mask, dpImm6Pos)

	return shift, amount
}

// WithShift encodes a shifted-register operand into the instruction.
// Sets ext=0, the `shift`, and the shift amount into `imm6` field.
// Returns the updated instruction
func (i DPInstruction) WithShift(shift Shift, amount uint8) DPInstruction {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0xA && opcode != 0xB {
		log.Fatalf("cannot set shift due to non register instruction")
	}

	i = set(i, uint8(0), dpExtMask, dpExtPos)
	i = set(i, uint8(shift), dpShiftMask, dpShiftPos)
	i = set(i, amount, dpImm6Mask, dpImm6Pos)

	return i
}

// Extension returns the extend option and shift amount for extended-register
// instructions. Valid only for extended-register form instructions
func (i DPInstruction) Extension() (Extension, uint8) {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0xA && opcode != 0xB {
		log.Fatalf("cannot get extension due to non register instruction")
	}

	shift := get[Shift](i, dpShiftMask, dpShiftPos)
	ext := get[uint8](i, dpExtMask, dpExtPos)
	if ext == 0 || shift != 0x0 {
		log.Fatalf("cannot get shift due to non extended register instruction")
	}

	opt := get[Extension](i, dpOptMask, dpOptPos)
	amount := get[uint8](i, dpImm3Mask, dpImm3Pos)

	return opt, amount
}

// WithExtension encodes an extended-register operand into the instruction.
// Sets ext=0, shift=0, the extend option into `opt`, and the shift amount bits
// into `imm3` field. Returns the updated instruction
func (i DPInstruction) WithExtension(option Extension, amount uint8) DPInstruction {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0xA && opcode != 0xB {
		log.Fatalf("cannot set extension due to non register instruction")
	}

	i = set(i, uint8(1), dpExtMask, dpExtPos)
	i = set(i, uint8(0), dpShiftMask, dpShiftPos)
	i = set(i, uint8(option), dpOptMask, dpOptPos)
	i = set(i, amount, dpImm3Mask, dpImm3Pos)

	return i
}

// Immediate returns the 12-bit unsigned immediate from bits 21–10
func (i DPInstruction) Immediate() uint16 {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0x11 {
		log.Fatalf("cannot get immediate due to non arithmetic immediate instruction")
	}

	return get[uint16](i, dpImm12Mask, dpImm12Pos)
}

// WithImmediate encodes unsigned immediate into `imm12` field and
// returns the updated instruction
func (i DPInstruction) WithImmediate(imm uint16) DPInstruction {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0x11 {
		log.Fatalf("cannot set immediate due to non arithmetic immediate instruction")
	}

	return set(i, imm, dpImm12Mask, dpImm12Pos)
}

// Bitmask decodes the N/imms/immr fields and returns the 64-bit immediate
// value they represent. Valid only for logical immediate instructions
func (i DPInstruction) Bitmask() uint64 {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0x12 {
		log.Fatalf("cannot get bitmask due to non logical immediate instruction")
	}

	n := get[uint8](i, dpShiftMask>>1, dpShiftPos)
	imms := get[uint8](i, dpImmsMask, dpImmsPos)
	immr := get[uint8](i, dpImmrMask, dpImmrPos)

	return decodeBitmask(n, imms, immr)
}

// WithBitmask encodes value as an ARM64 bitmask immediate into the `N/imms/immr`
// fields and returns the updated instruction. The `sf` bit must be set before
// calling this method as it determines whether to use 32 or 64-bit encoding
func (i DPInstruction) WithBitmask(bitmask uint64) DPInstruction {
	if opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode != 0x12 {
		log.Fatalf("cannot set bitmask due to non logical immediate instruction")
	}

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

	switch opcode := (uint8(i.Opcode()) >> 0) & 0x1F; opcode {
	case 0xA, 0xB: // register family (AND/OR/EOR/ANDS/ADD/SUB and variants)
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

	case 0x11: // arithmetic immediate family (ADDI/ADDSI/SUBI/SUBSI)
		// shift bit at position 23 indicates imm12 << 12 when set
		if shift := get[uint8](i, dpShiftMask>>1, dpShiftPos); shift == 0 {
			fmt.Fprintf(&b, "#%#X", i.Immediate())
			break
		}

		fmt.Fprintf(&b, "#%#X, lsl #12", i.Immediate())

	case 0x12: // logical bitmask immediate family (ANDI/ORRI/EORI/ANDSI)
		fmt.Fprintf(&b, "#%#X", i.Bitmask())

	default:
		log.Fatalf("[BUG] Must be unreachable. Instruction: DPInstruction(%032b)", i)
	}

	return b.String()
}
