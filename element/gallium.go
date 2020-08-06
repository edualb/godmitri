package element

// Gallium strtuct for gallium element
type Gallium struct{}

// GetPeriod returns the gallium period
func (Gallium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the gallium group
func (Gallium) GetGroup() string {
	var g groupType = a3
	return g.get()
}

// GetCategory returns the gallium category
func (Gallium) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the gallium name
func (Gallium) GetName() string {
	return "Gallium"
}

// GetSimbol returns gallium simbol
func (Gallium) GetSimbol() string {
	return "Ga"
}

// GetAtomicNumber returns the gallium atomic number
func (Gallium) GetAtomicNumber() int {
	return 31
}

// GetAtomicWeight returns the gallium atomic weight
func (Gallium) GetAtomicWeight() float32 {
	return 69.723
}
