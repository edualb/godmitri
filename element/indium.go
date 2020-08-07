package element

// Indium strtuct for indium element
type Indium struct{}

// GetPeriod returns the indium period
func (Indium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the indium group
func (Indium) GetGroup() string {
	var g groupType = a3
	return g.get()
}

// GetCategory returns the indium category
func (Indium) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the indium name
func (Indium) GetName() string {
	return "Indium"
}

// GetSimbol returns indium simbol
func (Indium) GetSimbol() string {
	return "In"
}

// GetAtomicNumber returns the indium atomic number
func (Indium) GetAtomicNumber() int {
	return 49
}

// GetAtomicWeight returns the indium atomic weight
func (Indium) GetAtomicWeight() float32 {
	return 114.82
}
