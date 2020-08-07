package element

// Ruthenium strtuct for ruthenium element
type Ruthenium struct{}

// GetPeriod returns the ruthenium period
func (Ruthenium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the ruthenium group
func (Ruthenium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the ruthenium category
func (Ruthenium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the ruthenium name
func (Ruthenium) GetName() string {
	return "Ruthenium"
}

// GetSimbol returns ruthenium simbol
func (Ruthenium) GetSimbol() string {
	return "Ru"
}

// GetAtomicNumber returns the ruthenium atomic number
func (Ruthenium) GetAtomicNumber() int {
	return 44
}

// GetAtomicWeight returns the ruthenium atomic weight
func (Ruthenium) GetAtomicWeight() float32 {
	return 101.07
}
