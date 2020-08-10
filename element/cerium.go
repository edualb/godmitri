package element

// Cerium strtuct for cerium element
type Cerium struct{}

// GetPeriod returns the cerium period
func (Cerium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the cerium group
func (Cerium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the cerium category
func (Cerium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the cerium name
func (Cerium) GetName() string {
	return "Cerium"
}

// GetSimbol returns cerium simbol
func (Cerium) GetSimbol() string {
	return "Ce"
}

// GetAtomicNumber returns the cerium atomic number
func (Cerium) GetAtomicNumber() int {
	return 58
}

// GetAtomicWeight returns the cerium atomic weight
func (Cerium) GetAtomicWeight() float32 {
	return 140.12
}
