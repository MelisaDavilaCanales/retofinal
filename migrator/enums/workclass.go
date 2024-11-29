package enums

type Workclass string

const (
	Private       Workclass = "Private"
	SelfEmpNotInc Workclass = "Self-emp-not-inc"
	SelfEmpInc    Workclass = "Self-emp-inc"
	FederalGov    Workclass = "Federal-gov"
	LocalGov      Workclass = "Local-gov"
	StateGov      Workclass = "State-gov"
	WithoutPay    Workclass = "Without-pay"
	NeverWorked   Workclass = "Never-worked"
	nullWorkclass Workclass = ""
)
