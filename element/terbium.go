package element

// Terbium strtuct for terbium element
type Terbium struct{}

// GetPeriod returns the terbium period
func (Terbium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the terbium group
func (Terbium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the terbium category
func (Terbium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the terbium name
func (Terbium) GetName() string {
	return "Terbium"
}

// GetSimbol returns terbium simbol
func (Terbium) GetSimbol() string {
	return "Tb"
}

// GetAtomicNumber returns the terbium atomic number
func (Terbium) GetAtomicNumber() int {
	return 65
}

// GetAtomicWeight returns the terbium atomic weight
func (Terbium) GetAtomicWeight() float32 {
	return 158.93
}
