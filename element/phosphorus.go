package element

// Phosphorus strtuct for phosphorus element
type Phosphorus struct{}

// GetPeriod returns the phosphorus period
func (Phosphorus) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the phosphorus group
func (Phosphorus) GetGroup() string {
	var g groupType = a5
	return g.get()
}

// GetCategory returns the phosphorus category
func (Phosphorus) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the phosphorus name
func (Phosphorus) GetName() string {
	return "Phosphorus"
}

// GetSimbol returns phosphorus simbol
func (Phosphorus) GetSimbol() string {
	return "P"
}

// GetAtomicNumber returns the phosphorus atomic number
func (Phosphorus) GetAtomicNumber() int {
	return 15
}

// GetAtomicWeight returns the phosphorus atomic weight
func (Phosphorus) GetAtomicWeight() float32 {
	return 30.974
}
