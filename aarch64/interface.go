package aarch64

import (
	"fmt"
	"log"
	"math/bits"
)

type Extension uint8

const (
	ExtUXTB Extension = 0x0 // zero-extend byte
	ExtUXTH Extension = 0x1 // zero-extend half
	ExtUXTW Extension = 0x2 // zero-extend word
	ExtUXTX Extension = 0x3 // zero-extend dword
	ExtSXTB Extension = 0x4 // sign-extend byte
	ExtSXTH Extension = 0x5 // sign-extend half
	ExtSXTW Extension = 0x6 // sign-extend word
	ExtSXTX Extension = 0x7 // sign-extend dword

	ExtWLSL Extension = ExtUXTW // LSL for 32-bit operands
	ExtXLSL Extension = ExtUXTX // LSL for 64-bit operands
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

type Register uint8

const (
	W0 Register = iota
	W1
	W2
	W3
	W4
	W5
	W6
	W7
	W8
	W9
	W10
	W11
	W12
	W13
	W14
	W15
	W16
	W17
	W18
	W19
	W20
	W21
	W22
	W23
	W24
	W25
	W26
	W27
	W28
	W29
	W30
	w31

	X0
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	x29
	x30
	x31
)

const (
	FP  Register = x29
	LR  Register = x30
	SP  Register = x31
	XZR Register = x31
	WZR Register = w31
)

func (r Register) String() string {
	if r < X0 {
		return fmt.Sprintf("w%d", r)
	}

	return fmt.Sprintf("x%d", r-X0)
}

// Shift represents the shift kind used in shifted-register instructions
type Shift uint8

const (
	ShiftLSL Shift = 0x0
	ShiftLSR Shift = 0x1
	ShiftASR Shift = 0x2
	ShiftROR Shift = 0x3
)

var shifts = map[Shift]string{
	ShiftLSL: "lsl",
	ShiftLSR: "lsr",
	ShiftASR: "asr",
	ShiftROR: "ror",
}

// String returns the lowercase ARM64 shift mnemonic
// Returns a formatted fallback string if the shift is not recognized
func (s Shift) String() string {
	if name, ok := shifts[s]; ok {
		return name
	}

	return fmt.Sprintf("Shift(%#03X)", uint8(s))
}

type Instruction uint32

type operand interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

func get[O, I operand](inst I, size, shift uint32) O {
	mask := (uint32(1) << size) - 1
	return O((uint32(inst) >> shift) & mask)
}

func set[O, I operand](inst I, op O, size, shift uint32) I {
	mask := (uint32(1) << size) - 1
	if uintOp := uint32(op); uintOp > mask {
		log.Fatalf("expected %d-bit number, but got %#X", size, uintOp)
	}

	return I((uint32(inst) & ^(mask << shift)) | ((uint32(op) << shift) & (mask << shift)))
}

func getReg[I ~uint32](inst I, size, pos uint32) Register {
	reg := get[Register](inst, size, pos)
	if sf := get[uint8](inst, 1, 31); sf == 0 {
		return reg
	}

	return reg + X0
}

func setReg[I ~uint32](inst I, reg Register, size, pos uint32) I {
	return set(inst, reg%X0, size, pos)
}

func getBool[I ~uint32](inst I, size, pos uint32) bool {
	flag := get[uint8](inst, size, pos)
	return flag == 0x1
}

func setBool[I ~uint32](inst I, flag bool, size, pos uint32) I {
	if flag {
		return set[uint8](inst, 1, size, pos)
	}

	return inst
}

func getBitmask[I ~uint32](inst I, size, basePos uint32) uint64 {
	immSize := (size - 1) / 2

	imms := get[uint8](inst, immSize, basePos)
	immr := get[uint8](inst, immSize, basePos+immSize)
	n := get[uint8](inst, 1, basePos+2*immSize)

	return decodeBitmask(n, immr, imms)
}

func decodeBitmask(n, immr, imms uint8) uint64 {
	nimms := (uint8(n) << 6) | (imms & 0x3F)

	len := 7 - bits.LeadingZeros8(^nimms)
	if len < 1 {
		return 0
	}

	elmSize := uint8(1) << len
	immsMask := uint8(elmSize - 1)

	ones := uint8(imms&immsMask) + 1
	if ones == elmSize {
		return 0
	}

	elm := ^uint64(0)
	if ones != 64 {
		elm = (uint64(1) << ones) - 1
	}

	if rotation := immr & (elmSize - 1); rotation > 0 {
		elmMask := (uint64(1) << elmSize) - 1
		elm = ((elm << rotation) | (elm >> (elmSize - rotation))) & elmMask
	}

	value := elm
	for size := elmSize; size < 64; size *= 2 {
		value |= value << size
	}

	return value
}

func setBitmask[I ~uint32](inst I, value uint64, size, basePos uint32) I {
	n, immr, imms := encodeBitmask(value, getSF(inst))
	immSize := (size - 1) / 2

	inst = set(inst, imms, immSize, basePos)
	inst = set(inst, immr, immSize, basePos+immSize)
	inst = set(inst, n, 1, basePos+2*immSize)

	return inst
}

func encodeBitmask(value uint64, hasSF bool) (n, immr, imms uint8) {
	if value == 0 ||
		(hasSF && value == ^uint64(0)) ||
		(!hasSF && value == 0xFFFFFFFF) {
		return
	}

	if !hasSF {
		value32 := uint32(value)
		value = uint64(value32) | uint64(value32)<<32
	}

	elmSize := uint64(64)
	for size := uint64(2); size <= 32; size *= 2 {
		if bits.RotateLeft64(value, int(size)) == value {
			elmSize = size
			break
		}
	}

	elmMask := ^uint64(0)
	if elmSize != 64 {
		elmMask = (uint64(1) << elmSize) - 1
	}

	elm := value & elmMask
	zeros := uint64(bits.TrailingZeros64(elm))
	rotated := (elm>>zeros | elm<<(elmSize-zeros)) & elmMask
	ones := bits.TrailingZeros64(^rotated & elmMask)

	if rotated != (uint64(1)<<ones)-1 {
		return
	}

	n, imms = 1, uint8(ones-1)
	if elmSize != 64 {
		log2s := bits.Len(uint(elmSize)) - 1
		tag := uint8((0x3F<<log2s)&0x3F) & ^uint8(1<<log2s)
		n, imms = 0, tag|uint8(ones-1)
	}

	immr = uint8(((elmSize - zeros) % elmSize) & 0x3F)
	return
}

func isSF(regs ...Register) bool {
	if len(regs) == 0 {
		return false
	}

	var flags uint8
	for i, reg := range regs {
		if w31 < reg {
			flags = set[uint8](flags, 1, 1, uint32(i))
		}
	}

	if ones := (uint8(1) << len(regs)) - 1; ones == flags {
		return true
	}

	return false
}

func getSF[I ~uint32](inst I) bool {
	return getBool(inst, 1, 31)
}

func setSF[I ~uint32](inst I, regs ...Register) I {
	flag := isSF(regs...)
	return setBool(inst, flag, 1, 31)
}
