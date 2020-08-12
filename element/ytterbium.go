package element

// Ytterbium strtuct for ytterbium element
type Ytterbium struct{}

// GetPeriod returns the ytterbium period
func (Ytterbium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the ytterbium group
func (Ytterbium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the ytterbium category
func (Ytterbium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the ytterbium name
func (Ytterbium) GetName() string {
	return "Ytterbium"
}

// GetSimbol returns ytterbium simbol
func (Ytterbium) GetSimbol() string {
	return "Yb"
}

// GetAtomicNumber returns the ytterbium atomic number
func (Ytterbium) GetAtomicNumber() int {
	return 70
}

// GetAtomicWeight returns the ytterbium atomic weight
func (Ytterbium) GetAtomicWeight() float32 {
	return 173.05
}
