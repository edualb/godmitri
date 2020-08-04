package element

// Magnesium strtuct for magnesium element
type Magnesium struct{}

// GetPeriod returns the magnesium period
func (Magnesium) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the magnesium group
func (Magnesium) GetGroup() string {
	var g groupType = a2
	return g.get()
}

// GetCategory returns the magnesium category
func (Magnesium) GetCategory() string {
	var c categoryType = alkalineEarthMetal
	return c.get()
}

// GetName returns the magnesium name
func (Magnesium) GetName() string {
	return "Magnesium"
}

// GetSimbol returns magnesium simbol
func (Magnesium) GetSimbol() string {
	return "Mg"
}

// GetAtomicNumber returns the magnesium atomic number
func (Magnesium) GetAtomicNumber() int {
	return 12
}

// GetAtomicWeight returns the magnesium atomic weight
func (Magnesium) GetAtomicWeight() float32 {
	return 24.305
}
