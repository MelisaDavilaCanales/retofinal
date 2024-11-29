package enums

type MaritalStatus string

const (
	MarriedCivSpouse    MaritalStatus = "Married-civ-spouse"
	Divorced            MaritalStatus = "Divorced"
	NeverMarried        MaritalStatus = "Never-married"
	Separated           MaritalStatus = "Separated"
	Widowed             MaritalStatus = "Widowed"
	MarriedSpouseAbsent MaritalStatus = "Married-spouse-absent"
	MarriedAFSpouse     MaritalStatus = "Married-AF-spouse"
)
