package element

// Hafnium strtuct for hafnium element
type Hafnium struct{}

// GetPeriod returns the hafnium period
func (Hafnium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the hafnium group
func (Hafnium) GetGroup() string {
	var g groupType = b4
	return g.get()
}

// GetCategory returns the hafnium category
func (Hafnium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the hafnium name
func (Hafnium) GetName() string {
	return "Hafnium"
}

// GetSimbol returns hafnium simbol
func (Hafnium) GetSimbol() string {
	return "Hf"
}

// GetAtomicNumber returns the hafnium atomic number
func (Hafnium) GetAtomicNumber() int {
	return 72
}

// GetAtomicWeight returns the helium atomic weight
func (Hafnium) GetAtomicWeight() float32 {
	return 178.49
}
