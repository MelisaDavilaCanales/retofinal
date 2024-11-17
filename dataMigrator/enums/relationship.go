package enums

type Relationship string

const (
	Wife          Relationship = "Wife"
	OwnChild      Relationship = "Own-child"
	Husband       Relationship = "Husband"
	NotInFamily   Relationship = "Not-in-family"
	OtherRelative Relationship = "Other-relative"
	Unmarried     Relationship = "Unmarried"
)
