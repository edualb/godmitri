package element

// Osmium strtuct for osmium element
type Osmium struct{}

// GetPeriod returns the osmium period
func (Osmium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the osmium group
func (Osmium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the osmium category
func (Osmium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the osmium name
func (Osmium) GetName() string {
	return "Osmium"
}

// GetSimbol returns osmium simbol
func (Osmium) GetSimbol() string {
	return "Os"
}

// GetAtomicNumber returns the osmium atomic number
func (Osmium) GetAtomicNumber() int {
	return 76
}

// GetAtomicWeight returns the osmium atomic weight
func (Osmium) GetAtomicWeight() float32 {
	return 190.23
}
