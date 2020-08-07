package element

// Rhodium strtuct for rhodium element
type Rhodium struct{}

// GetPeriod returns the rhodium period
func (Rhodium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the rhodium group
func (Rhodium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the rhodium category
func (Rhodium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the rhodium name
func (Rhodium) GetName() string {
	return "Rhodium"
}

// GetSimbol returns rhodium simbol
func (Rhodium) GetSimbol() string {
	return "Rh"
}

// GetAtomicNumber returns the rhodium atomic number
func (Rhodium) GetAtomicNumber() int {
	return 45
}

// GetAtomicWeight returns the rhodium atomic weight
func (Rhodium) GetAtomicWeight() float32 {
	return 102.91
}
