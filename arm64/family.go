package arm64

type Family uint8

const (
	FamilyUND Family = iota
	FamilyDPR        // Data Processing – Register (DPR) family
	FamilyDPI        // Data Processing – Immediate (DPI) family
	FamilyPCR        // PC-relative (PCR) family
	FamilyLSR        // Load/Store – Register (LSR) family
)

func resolveFamily(opcode uint8) Family {
	switch fewbits := opcode >> 3 & 0x1F; fewbits {
	// Data Processing – Register family
	case 0x01, 0x05, 0x09, 0x0D:
		return FamilyDPR

	// Data Processing – Immediate family
	case 0x02, 0x06, 0x0A, 0x0E:
		// ADR instruction part of the DPI family, but uses another format
		if opcode&0x1F == 0x10 {
			return FamilyPCR
		}

		return FamilyDPI

	// Load/Store – Register family
	case 0x1C:
		return FamilyLSR
	}

	return FamilyUND
}
