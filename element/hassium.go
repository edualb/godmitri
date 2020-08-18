package element

// Hassium strtuct for hassium element
type Hassium struct{}

// GetPeriod returns the hassium period
func (Hassium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the hassium group
func (Hassium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the hassium category
func (Hassium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the hassium name
func (Hassium) GetName() string {
	return "Hassium"
}

// GetSimbol returns hassium simbol
func (Hassium) GetSimbol() string {
	return "Hs"
}

// GetAtomicNumber returns the hassium atomic number
func (Hassium) GetAtomicNumber() int {
	return 108
}

// GetAtomicWeight returns the hassium atomic weight
func (Hassium) GetAtomicWeight() float32 {
	return 277
}
