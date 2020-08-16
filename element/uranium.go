package element

// Uranium strtuct for uranium element
type Uranium struct{}

// GetPeriod returns the uranium period
func (Uranium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the uranium group
func (Uranium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the uranium category
func (Uranium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the uranium name
func (Uranium) GetName() string {
	return "Uranium"
}

// GetSimbol returns uranium simbol
func (Uranium) GetSimbol() string {
	return "U"
}

// GetAtomicNumber returns the uranium atomic number
func (Uranium) GetAtomicNumber() int {
	return 92
}

// GetAtomicWeight returns the uranium atomic weight
func (Uranium) GetAtomicWeight() float32 {
	return 238.03
}
