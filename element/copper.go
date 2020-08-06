package element

// Copper strtuct for copper element
type Copper struct{}

// GetPeriod returns the copper period
func (Copper) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the copper group
func (Copper) GetGroup() string {
	var g groupType = b1
	return g.get()
}

// GetCategory returns the copper category
func (Copper) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the copper name
func (Copper) GetName() string {
	return "Copper"
}

// GetSimbol returns copper simbol
func (Copper) GetSimbol() string {
	return "Cu"
}

// GetAtomicNumber returns the copper atomic number
func (Copper) GetAtomicNumber() int {
	return 29
}

// GetAtomicWeight returns the copper atomic weight
func (Copper) GetAtomicWeight() float32 {
	return 63.546
}
