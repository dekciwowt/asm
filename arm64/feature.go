package arm64

import "fmt"

type Feature uint8

const (
	FeatNone Feature = iota
	FeatCRC32
	FeatCSSC
	FeatMTE
	FeatPAuth
	FeatPAuthLR
	FeatFlagM
	FeatCPA
)

var features = map[Feature]string{
	FeatNone:    "–",
	FeatCRC32:   "CRC32",
	FeatCSSC:    "CSSC",
	FeatMTE:     "MTE",
	FeatPAuth:   "PAuth",
	FeatPAuthLR: "PAuthLR",
	FeatFlagM:   "FlagM",
	FeatCPA:     "CPA",
}

func (f Feature) String() string {
	if name, ok := features[f]; ok {
		return name
	}

	return fmt.Sprintf("Feature(%d)", uint8(f))
}
