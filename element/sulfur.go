package element

// Sulfur strtuct for sulfur element
type Sulfur struct{}

// GetPeriod returns the sulfur period
func (Sulfur) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the sulfur group
func (Sulfur) GetGroup() string {
	var g groupType = a6
	return g.get()
}

// GetCategory returns the sulfur category
func (Sulfur) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the sulfur name
func (Sulfur) GetName() string {
	return "Sulfur"
}

// GetSimbol returns sulfur simbol
func (Sulfur) GetSimbol() string {
	return "S"
}

// GetAtomicNumber returns the sulfur atomic number
func (Sulfur) GetAtomicNumber() int {
	return 16
}

// GetAtomicWeight returns the sulfur atomic weight
func (Sulfur) GetAtomicWeight() float32 {
	return 32.06
}
