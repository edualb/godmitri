package element

// Thulium strtuct for thulium element
type Thulium struct{}

// GetPeriod returns the thulium period
func (Thulium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the thulium group
func (Thulium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the thulium category
func (Thulium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the thulium name
func (Thulium) GetName() string {
	return "Thulium"
}

// GetSimbol returns thulium simbol
func (Thulium) GetSimbol() string {
	return "Tm"
}

// GetAtomicNumber returns the thulium atomic number
func (Thulium) GetAtomicNumber() int {
	return 69
}

// GetAtomicWeight returns the thulium atomic weight
func (Thulium) GetAtomicWeight() float32 {
	return 168.93
}
