package element

// Chlorine strtuct for chlorine element
type Chlorine struct{}

// GetPeriod returns the chlorine period
func (Chlorine) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the chlorine group
func (Chlorine) GetGroup() string {
	var g groupType = a7
	return g.get()
}

// GetCategory returns the chlorine category
func (Chlorine) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the chlorine name
func (Chlorine) GetName() string {
	return "Chlorine"
}

// GetSimbol returns chlorine simbol
func (Chlorine) GetSimbol() string {
	return "Cl"
}

// GetAtomicNumber returns the chlorine atomic number
func (Chlorine) GetAtomicNumber() int {
	return 17
}

// GetAtomicWeight returns the chlorine atomic weight
func (Chlorine) GetAtomicWeight() float32 {
	return 35.45
}
