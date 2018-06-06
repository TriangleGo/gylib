package boolean

type Boolean string

const (
	Yes Boolean = "Yes"
	No Boolean = "No"
)

func FromString(in string) Boolean {
	inBoolean := Boolean(in)
	if inBoolean == Yes {
		return Yes
	} else {
		return No
	}
}
