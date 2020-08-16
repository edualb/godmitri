package element

// Curium strtuct for curium element
type Curium struct{}

// GetPeriod returns the curium period
func (Curium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the curium group
func (Curium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the curium category
func (Curium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the curium name
func (Curium) GetName() string {
	return "Curium"
}

// GetSimbol returns curium simbol
func (Curium) GetSimbol() string {
	return "Cm"
}

// GetAtomicNumber returns the curium atomic number
func (Curium) GetAtomicNumber() int {
	return 96
}

// GetAtomicWeight returns the curium atomic weight
func (Curium) GetAtomicWeight() float32 {
	return 247
}
