package element

// Fermium strtuct for fermium element
type Fermium struct{}

// GetPeriod returns the fermium period
func (Fermium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the fermium group
func (Fermium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the fermium category
func (Fermium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the fermium name
func (Fermium) GetName() string {
	return "Fermium"
}

// GetSimbol returns fermium simbol
func (Fermium) GetSimbol() string {
	return "Fm"
}

// GetAtomicNumber returns the fermium atomic number
func (Fermium) GetAtomicNumber() int {
	return 100
}

// GetAtomicWeight returns the fermium atomic weight
func (Fermium) GetAtomicWeight() float32 {
	return 257
}
