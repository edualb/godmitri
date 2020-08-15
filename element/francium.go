package element

// Francium strtuct for francium element
type Francium struct{}

// GetPeriod returns the francium period
func (Francium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the francium group
func (Francium) GetGroup() string {
	var g groupType = a1
	return g.get()
}

// GetCategory returns the francium category
func (Francium) GetCategory() string {
	var c categoryType = alkaliMetal
	return c.get()
}

// GetName returns the francium name
func (Francium) GetName() string {
	return "Francium"
}

// GetSimbol returns francium simbol
func (Francium) GetSimbol() string {
	return "Fr"
}

// GetAtomicNumber returns the francium atomic number
func (Francium) GetAtomicNumber() int {
	return 87
}

// GetAtomicWeight returns the francium atomic weight
func (Francium) GetAtomicWeight() float32 {
	return 223
}
