package element

// Nitrogen strtuct for nitrogen element
type Nitrogen struct{}

// GetPeriod returns the nitrogen period
func (Nitrogen) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the nitrogen group
func (Nitrogen) GetGroup() string {
	var g groupType = a5
	return g.get()
}

// GetCategory returns the nitrogen category
func (Nitrogen) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the nitrogen name
func (Nitrogen) GetName() string {
	return "Nitrogen"
}

// GetSimbol returns nitrogen simbol
func (Nitrogen) GetSimbol() string {
	return "N"
}

// GetAtomicNumber returns the nitrogen atomic number
func (Nitrogen) GetAtomicNumber() int {
	return 7
}

// GetAtomicWeight returns the nitrogen atomic weight
func (Nitrogen) GetAtomicWeight() float32 {
	return 14.007
}
