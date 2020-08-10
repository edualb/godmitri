package element

// Praseodymium strtuct for praseodymium element
type Praseodymium struct{}

// GetPeriod returns the praseodymium period
func (Praseodymium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the praseodymium group
func (Praseodymium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the praseodymium category
func (Praseodymium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the praseodymium name
func (Praseodymium) GetName() string {
	return "Praseodymium"
}

// GetSimbol returns praseodymium simbol
func (Praseodymium) GetSimbol() string {
	return "Pr"
}

// GetAtomicNumber returns the praseodymium atomic number
func (Praseodymium) GetAtomicNumber() int {
	return 59
}

// GetAtomicWeight returns the praseodymium atomic weight
func (Praseodymium) GetAtomicWeight() float32 {
	return 140.91
}
