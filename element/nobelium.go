package element

// Nobelium strtuct for nobelium element
type Nobelium struct{}

// GetPeriod returns the nobelium period
func (Nobelium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the nobelium group
func (Nobelium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the nobelium category
func (Nobelium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the nobelium name
func (Nobelium) GetName() string {
	return "Nobelium"
}

// GetSimbol returns nobelium simbol
func (Nobelium) GetSimbol() string {
	return "No"
}

// GetAtomicNumber returns the nobelium atomic number
func (Nobelium) GetAtomicNumber() int {
	return 102
}

// GetAtomicWeight returns the nobelium atomic weight
func (Nobelium) GetAtomicWeight() float32 {
	return 259
}
