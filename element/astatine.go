package element

// Astatine strtuct for astatine element
type Astatine struct{}

// GetPeriod returns the astatine period
func (Astatine) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the astatine group
func (Astatine) GetGroup() string {
	var g groupType = a7
	return g.get()
}

// GetCategory returns the astatine category
func (Astatine) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the astatine name
func (Astatine) GetName() string {
	return "Astatine"
}

// GetSimbol returns astatine simbol
func (Astatine) GetSimbol() string {
	return "At"
}

// GetAtomicNumber returns the astatine atomic number
func (Astatine) GetAtomicNumber() int {
	return 85
}

// GetAtomicWeight returns the astatine atomic weight
func (Astatine) GetAtomicWeight() float32 {
	return 210
}
