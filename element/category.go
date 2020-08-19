package element

type categoryType int

const (
	nonMetal = iota
	nobleGas
	alkaliMetal
	alkalineEarthMetal
	metalloid
	postTransitionMetal
	transitionMetal
	lanthanoid
	actinoid
	unknown
)

// Get gets the categoryType
func (ct categoryType) get() string {
	return [...]string{
		"Non-metal",
		"Noble gas",
		"Alkali metal",
		"Alkaline earth metal",
		"Metalloid",
		"Post-transition metal",
		"Transition metal",
		"Lanthanoid",
		"Actinoid",
		"Unknown",
	}[ct]
}
