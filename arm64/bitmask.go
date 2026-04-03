package arm64

import "math/bits"

func encodeBitmask(value uint64, hasSF bool) (n, immr, imms uint8) {
	if value == 0 || (hasSF && value == ^uint64(0)) || (!hasSF && value == 0xFFFFFFFF) {
		return
	}

	if !hasSF {
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
