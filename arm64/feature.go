package arm64

type Feature uint8

const (
	FeatGeneral Feature = iota
	FeatCRC32
	FeatCSSC
	FeatMTE
	FeatPAuth
	FeatPAuthLR
	FeatFlagM
	FeatCPA
)
