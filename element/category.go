package element

type categoryType int

const (
	nonMetal = iota
	nobleGas
	alkaliMetal
	alkalineEarthMetal
)

// Get gets the categoryType
func (ct categoryType) get() string {
	return [...]string{
		"Non-metal",
		"Noble gas",
		"Alkali metal",
		"Alkaline earth metal",
	}[ct]
}
