package element

// Europium strtuct for europium element
type Europium struct{}

// GetPeriod returns the europium period
func (Europium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the europium group
func (Europium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the europium category
func (Europium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the europium name
func (Europium) GetName() string {
	return "Europium"
}

// GetSimbol returns europium simbol
func (Europium) GetSimbol() string {
	return "Eu"
}

// GetAtomicNumber returns the europium atomic number
func (Europium) GetAtomicNumber() int {
	return 63
}

// GetAtomicWeight returns the chromium atomic weight
func (Europium) GetAtomicWeight() float32 {
	return 151.96
}
