package element

// Sodium strtuct for sodium element
type Sodium struct{}

// GetPeriod returns the sodium period
func (Sodium) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the sodium group
func (Sodium) GetGroup() string {
	var g groupType = a1
	return g.get()
}

// GetCategory returns the sodium category
func (Sodium) GetCategory() string {
	var c categoryType = alkaliMetal
	return c.get()
}

// GetName returns the sodium name
func (Sodium) GetName() string {
	return "Sodium"
}

// GetSimbol returns sodium simbol
func (Sodium) GetSimbol() string {
	return "Na"
}

// GetAtomicNumber returns the sodium atomic number
func (Sodium) GetAtomicNumber() int {
	return 11
}

// GetAtomicWeight returns the sodium atomic weight
func (Sodium) GetAtomicWeight() float32 {
	return 22.990
}
