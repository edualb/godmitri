package element

// Americium strtuct for americium element
type Americium struct{}

// GetPeriod returns the americium period
func (Americium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the americium group
func (Americium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the americium category
func (Americium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the americium name
func (Americium) GetName() string {
	return "Americium"
}

// GetSimbol returns americium simbol
func (Americium) GetSimbol() string {
	return "Am"
}

// GetAtomicNumber returns the americium atomic number
func (Americium) GetAtomicNumber() int {
	return 95
}

// GetAtomicWeight returns the americium atomic weight
func (Americium) GetAtomicWeight() float32 {
	return 243
}
