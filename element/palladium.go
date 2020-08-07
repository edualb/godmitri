package element

// Palladium strtuct for palladium element
type Palladium struct{}

// GetPeriod returns the palladium period
func (Palladium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the palladium group
func (Palladium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the palladium category
func (Palladium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the palladium name
func (Palladium) GetName() string {
	return "Palladium"
}

// GetSimbol returns palladium simbol
func (Palladium) GetSimbol() string {
	return "Pd"
}

// GetAtomicNumber returns the palladium atomic number
func (Palladium) GetAtomicNumber() int {
	return 46
}

// GetAtomicWeight returns the palladium atomic weight
func (Palladium) GetAtomicWeight() float32 {
	return 106.42
}
