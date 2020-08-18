package element

// Bohrium strtuct for bohrium element
type Bohrium struct{}

// GetPeriod returns the bohrium period
func (Bohrium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the bohrium group
func (Bohrium) GetGroup() string {
	var g groupType = b7
	return g.get()
}

// GetCategory returns the bohrium category
func (Bohrium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the bohrium name
func (Bohrium) GetName() string {
	return "Bohrium"
}

// GetSimbol returns bohrium simbol
func (Bohrium) GetSimbol() string {
	return "Bh"
}

// GetAtomicNumber returns the bohrium atomic number
func (Bohrium) GetAtomicNumber() int {
	return 107
}

// GetAtomicWeight returns the bohrium atomic weight
func (Bohrium) GetAtomicWeight() float32 {
	return 270
}
