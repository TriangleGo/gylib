package gender

import "strings"

type Gender string

const (
	Unknown Gender = "unknown"
	Male Gender = "male"
	Female Gender = "female"
)

func FromString(in string) Gender {
	lowIn := strings.ToLower(in)
	genderIn := Gender(lowIn)
	switch genderIn {
	case Male:
		return Male
	case Female:
		return Female
	default:
		return Unknown
	}
}