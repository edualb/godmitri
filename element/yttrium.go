package element

// Yttrium strtuct for yttrium element
type Yttrium struct{}

// GetPeriod returns the yttrium period
func (Yttrium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the yttrium group
func (Yttrium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the yttrium category
func (Yttrium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the yttrium name
func (Yttrium) GetName() string {
	return "Yttrium"
}

// GetSimbol returns yttrium simbol
func (Yttrium) GetSimbol() string {
	return "Y"
}

// GetAtomicNumber returns the yttrium atomic number
func (Yttrium) GetAtomicNumber() int {
	return 39
}

// GetAtomicWeight returns the yttrium atomic weight
func (Yttrium) GetAtomicWeight() float32 {
	return 88.906
}
