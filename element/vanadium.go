package element

// Vanadium strtuct for vanadium element
type Vanadium struct{}

// GetPeriod returns the vanadium period
func (Vanadium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the vanadium group
func (Vanadium) GetGroup() string {
	var g groupType = b5
	return g.get()
}

// GetCategory returns the vanadium category
func (Vanadium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the vanadium name
func (Vanadium) GetName() string {
	return "Vanadium"
}

// GetSimbol returns vanadium simbol
func (Vanadium) GetSimbol() string {
	return "V"
}

// GetAtomicNumber returns the vanadium atomic number
func (Vanadium) GetAtomicNumber() int {
	return 23
}

// GetAtomicWeight returns the vanadium atomic weight
func (Vanadium) GetAtomicWeight() float32 {
	return 50.942
}
