package element

// Neodymium strtuct for neodymium element
type Neodymium struct{}

// GetPeriod returns the neodymium period
func (Neodymium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the neodymium group
func (Neodymium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the neodymium category
func (Neodymium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the neodymium name
func (Neodymium) GetName() string {
	return "Neodymium"
}

// GetSimbol returns neodymium simbol
func (Neodymium) GetSimbol() string {
	return "Nd"
}

// GetAtomicNumber returns the neodymium atomic number
func (Neodymium) GetAtomicNumber() int {
	return 60
}

// GetAtomicWeight returns the neodymium atomic weight
func (Neodymium) GetAtomicWeight() float32 {
	return 144.24
}
