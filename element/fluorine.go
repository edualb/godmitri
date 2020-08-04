package element

// Oxygen strtuct for fluorine element
type Fluorine struct{}

// GetPeriod returns the fluorine period
func (Fluorine) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the fluorine group
func (Fluorine) GetGroup() string {
	var g groupType = a7
	return g.get()
}

// GetCategory returns the fluorine category
func (Fluorine) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the fluorine name
func (Fluorine) GetName() string {
	return "Fluorine"
}

// GetSimbol returns fluorine simbol
func (Fluorine) GetSimbol() string {
	return "F"
}

// GetAtomicNumber returns the fluorine atomic number
func (Fluorine) GetAtomicNumber() int {
	return 9
}

// GetAtomicWeight returns the fluorine atomic weight
func (Fluorine) GetAtomicWeight() float32 {
	return 18.998
}
