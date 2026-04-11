package arm64

import (
	"fmt"
)

// Data-Processing instruction layout
//
//	31   30   29  28-24      23-21  20-16      15-10        9-5        4-0
//	+----+----+---+----------+------+----------+------------+----------+----------+
//	| sf | op | S |   opt1   | opt2 |   opt3   |    opt4    |   opt5   |   opt6   |
//	+----+----+---+----------+------+----------+------------+----------+----------+
type Instruction uint32

const (
	opt6Size uint32 = 5
	opt5Size uint32 = 5
	opt4Size uint32 = 6
	opt3Size uint32 = 5
	opt2Size uint32 = 3
	opt1Size uint32 = 5
	sSize    uint32 = 1
	opSize   uint32 = 1
	sfSize   uint32 = 1

	opt6Mask uint32 = (1 << opt6Size) - 1
	opt5Mask uint32 = (1 << opt5Size) - 1
	opt4Mask uint32 = (1 << opt4Size) - 1
	opt3Mask uint32 = (1 << opt3Size) - 1
	opt2Mask uint32 = (1 << opt2Size) - 1
	opt1Mask uint32 = (1 << opt1Size) - 1
	sMask    uint32 = (1 << sSize) - 1
	opMask   uint32 = (1 << opSize) - 1
	sfMask   uint32 = (1 << sfSize) - 1

	opt6Pos uint32 = 0
	opt5Pos uint32 = 5
	opt4Pos uint32 = 10
	opt3Pos uint32 = 16
	opt2Pos uint32 = 21
	opt1Pos uint32 = 24
	sPos    uint32 = 29
	opPos   uint32 = 30
	sfPos   uint32 = 31
)

type category uint8

const (
	catDataProcLogicReg category = 0x0A // Logical (shifted register)
	catDataProcArithReg category = 0x0B // Add/subtract (shifted/extended register)
	catDataProcArithImm category = 0x11 // Add/subtract (immediate)
	catDataProcLogicImm category = 0x12 // Logical (immediate) / Move wide (immediate)
	catDataProcBitfield category = 0x13 // Bitfield / Extract
	catDataProcNSources category = 0x1A // Add/subtract with carry, Cond compare, Cond select, DP 1/2-source
	catDataProc3Sources category = 0x1B // Data-processing (3 source)
)

var categories = map[category]string{
	catDataProcLogicReg: "Data-Processing – Logical (shifted register)",
	catDataProcArithReg: "Data-Processing – Add/Substract (shifted register)",
	catDataProcArithImm: "Data-Processing – Add/Substract (immediate)",
	catDataProcLogicImm: "Data-Processing – Logical (immediate)",
	catDataProcNSources: "Data-Processing – 1/2 source",
	catDataProc3Sources: "Data-Processing – 3 source",
}

func (c category) String() string {
	if name, ok := categories[c]; ok {
		return name
	}

	return fmt.Sprintf("category(%b)", uint8(c))
}

var ( // Data-Processing (Logical)
	instANDw = identity[Instruction](0x0, 0x0, 0x0, 0x0, 0x0, catDataProcLogicReg, 0, 0, 0)
	instANDx = identity[Instruction](0x0, 0x0, 0x0, 0x0, 0x0, catDataProcLogicReg, 0, 0, 1)
)

func identity[I ~uint32, O operand](opt6, opt5, opt4, opt3, opt2, opt1, s, op, sf O) I {
	var i I

	i = set(i, opt6, opt6Mask, opt6Pos)
	i = set(i, opt5, opt5Mask, opt5Pos)
	i = set(i, opt4, opt4Mask, opt4Pos)
	i = set(i, opt3, opt3Mask, opt3Pos)
	i = set(i, opt2, opt2Mask, opt2Pos)
	i = set(i, opt1, opt1Mask, opt1Pos)
	i = set(i, s, sMask, sPos)
	i = set(i, op, opMask, opPos)
	i = set(i, sf, sfMask, sfPos)

	return i
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

func getReg[I ~uint32](inst I, mask, shift uint32) Register {
	reg := get[Register](inst, mask, shift)
	if sf := get[uint8](inst, sfMask, sfPos); sf == 0 {
		return reg
	}

	return reg + X0
}

func setReg[I ~uint32](inst I, reg Register, mask, shift uint32) I {
	return set(inst, reg%X0, mask, shift)
}
