package element

type categoryType int

const (
	nonMetal = iota
	nobleGas
)

// Get gets the categoryType
func (ct categoryType) get() string {
	return [...]string{
		"non-metal",
		"noble gas",
	}[ct]
}
