package element

// Copernicium strtuct for copernicium element
type Copernicium struct{}

// GetPeriod returns the copernicium period
func (Copernicium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the copernicium group
func (Copernicium) GetGroup() string {
	var g groupType = b2
	return g.get()
}

// GetCategory returns the copernicium category
func (Copernicium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the copernicium name
func (Copernicium) GetName() string {
	return "Copernicium"
}

// GetSimbol returns copernicium simbol
func (Copernicium) GetSimbol() string {
	return "Cn"
}

// GetAtomicNumber returns the copernicium atomic number
func (Copernicium) GetAtomicNumber() int {
	return 112
}

// GetAtomicWeight returns the copernicium atomic weight
func (Copernicium) GetAtomicWeight() float32 {
	return 285
}
