package element

// Meitnerium strtuct for meitnerium element
type Meitnerium struct{}

// GetPeriod returns the meitnerium period
func (Meitnerium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the meitnerium group
func (Meitnerium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the meitnerium category
func (Meitnerium) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the meitnerium name
func (Meitnerium) GetName() string {
	return "Meitnerium"
}

// GetSimbol returns meitnerium simbol
func (Meitnerium) GetSimbol() string {
	return "Mt"
}

// GetAtomicNumber returns the meitnerium atomic number
func (Meitnerium) GetAtomicNumber() int {
	return 109
}

// GetAtomicWeight returns the meitnerium atomic weight
func (Meitnerium) GetAtomicWeight() float32 {
	return 278
}
