package element

// Scandium strtuct for scandium element
type Scandium struct{}

// GetPeriod returns the scandium period
func (Scandium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the scandium group
func (Scandium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the scandium category
func (Scandium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the scandium name
func (Scandium) GetName() string {
	return "Scandium"
}

// GetSimbol returns scandium simbol
func (Scandium) GetSimbol() string {
	return "Sc"
}

// GetAtomicNumber returns the scandium atomic number
func (Scandium) GetAtomicNumber() int {
	return 21
}

// GetAtomicWeight returns the scandium atomic weight
func (Scandium) GetAtomicWeight() float32 {
	return 44.956
}
