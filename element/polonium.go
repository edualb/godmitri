package element

// Polonium strtuct for polonium element
type Polonium struct{}

// GetPeriod returns the polonium period
func (Polonium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the polonium group
func (Polonium) GetGroup() string {
	var g groupType = a6
	return g.get()
}

// GetCategory returns the polonium category
func (Polonium) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the polonium name
func (Polonium) GetName() string {
	return "Polonium"
}

// GetSimbol returns polonium simbol
func (Polonium) GetSimbol() string {
	return "Po"
}

// GetAtomicNumber returns the platinum atomic number
func (Polonium) GetAtomicNumber() int {
	return 84
}

// GetAtomicWeight returns the polonium atomic weight
func (Polonium) GetAtomicWeight() float32 {
	return 209
}
