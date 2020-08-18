package element

// Rutherfordium strtuct for rutherfordium element
type Rutherfordium struct{}

// GetPeriod returns the rutherfordium period
func (Rutherfordium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the rutherfordium group
func (Rutherfordium) GetGroup() string {
	var g groupType = b4
	return g.get()
}

// GetCategory returns the rutherfordium category
func (Rutherfordium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the rutherfordium name
func (Rutherfordium) GetName() string {
	return "Rutherfordium"
}

// GetSimbol returns rutherfordium simbol
func (Rutherfordium) GetSimbol() string {
	return "Rf"
}

// GetAtomicNumber returns the rutherfordium atomic number
func (Rutherfordium) GetAtomicNumber() int {
	return 104
}

// GetAtomicWeight returns the rutherfordium atomic weight
func (Rutherfordium) GetAtomicWeight() float32 {
	return 267
}
