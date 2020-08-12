package element

// Erbium strtuct for erbium element
type Erbium struct{}

// GetPeriod returns the erbium period
func (Erbium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the erbium group
func (Erbium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the erbium category
func (Erbium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the erbium name
func (Erbium) GetName() string {
	return "Erbium"
}

// GetSimbol returns erbium simbol
func (Erbium) GetSimbol() string {
	return "Er"
}

// GetAtomicNumber returns the erbium atomic number
func (Erbium) GetAtomicNumber() int {
	return 68
}

// GetAtomicWeight returns the erbium atomic weight
func (Erbium) GetAtomicWeight() float32 {
	return 167.26
}
