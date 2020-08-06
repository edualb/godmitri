package element

// Manganese strtuct for manganese element
type Manganese struct{}

// GetPeriod returns the manganese period
func (Manganese) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the manganese group
func (Manganese) GetGroup() string {
	var g groupType = b7
	return g.get()
}

// GetCategory returns the manganese category
func (Manganese) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the manganese name
func (Manganese) GetName() string {
	return "Manganese"
}

// GetSimbol returns manganese simbol
func (Manganese) GetSimbol() string {
	return "Mn"
}

// GetAtomicNumber returns the manganese atomic number
func (Manganese) GetAtomicNumber() int {
	return 25
}

// GetAtomicWeight returns the manganese atomic weight
func (Manganese) GetAtomicWeight() float32 {
	return 54.938
}
