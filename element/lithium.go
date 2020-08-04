package element

// Lithium strtuct for lithium element
type Lithium struct{}

// GetPeriod returns the lithium period
func (Lithium) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the lithium group
func (Lithium) GetGroup() string {
	var g groupType = a1
	return g.get()
}

// GetCategory returns the lithium category
func (Lithium) GetCategory() string {
	var c categoryType = alkaliMetal
	return c.get()
}

// GetName returns the lithium name
func (Lithium) GetName() string {
	return "Lithium"
}

// GetSimbol returns lithium simbol
func (Lithium) GetSimbol() string {
	return "Li"
}

// GetAtomicNumber returns the lithium atomic number
func (Lithium) GetAtomicNumber() int {
	return 3
}

// GetAtomicWeight returns the lithium atomic weight
func (Lithium) GetAtomicWeight() float32 {
	return 6.94
}
