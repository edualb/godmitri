package element

// Iridium strtuct for iridium element
type Iridium struct{}

// GetPeriod returns the iridium period
func (Iridium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the iridium group
func (Iridium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the iridium category
func (Iridium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the iridium name
func (Iridium) GetName() string {
	return "Iridium"
}

// GetSimbol returns iridium simbol
func (Iridium) GetSimbol() string {
	return "Ir"
}

// GetAtomicNumber returns the iridium atomic number
func (Iridium) GetAtomicNumber() int {
	return 77
}

// GetAtomicWeight returns the iridium atomic weight
func (Iridium) GetAtomicWeight() float32 {
	return 192.22
}
