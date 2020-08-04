package element

// Oxygen strtuct for oxygen element
type Oxygen struct{}

// GetPeriod returns the oxygen period
func (Oxygen) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the oxygen group
func (Oxygen) GetGroup() string {
	var g groupType = a6
	return g.get()
}

// GetCategory returns the oxygen category
func (Oxygen) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the oxygen name
func (Oxygen) GetName() string {
	return "Oxygen"
}

// GetSimbol returns oxygen simbol
func (Oxygen) GetSimbol() string {
	return "O"
}

// GetAtomicNumber returns the oxygen atomic number
func (Oxygen) GetAtomicNumber() int {
	return 8
}

// GetAtomicWeight returns the oxygen atomic weight
func (Oxygen) GetAtomicWeight() float32 {
	return 15.999
}
