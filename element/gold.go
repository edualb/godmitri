package element

// Gold strtuct for gold element
type Gold struct{}

// GetPeriod returns the gold period
func (Gold) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the gold group
func (Gold) GetGroup() string {
	var g groupType = b1
	return g.get()
}

// GetCategory returns the gold category
func (Gold) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the gold name
func (Gold) GetName() string {
	return "Gold"
}

// GetSimbol returns gold simbol
func (Gold) GetSimbol() string {
	return "Au"
}

// GetAtomicNumber returns the gold atomic number
func (Gold) GetAtomicNumber() int {
	return 79
}

// GetAtomicWeight returns the gold atomic weight
func (Gold) GetAtomicWeight() float32 {
	return 196.97
}
