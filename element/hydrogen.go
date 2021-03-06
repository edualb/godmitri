package element

// Hydrogen strtuct for hydrogen element
type Hydrogen struct{}

// GetPeriod returns the hydrogen period
func (Hydrogen) GetPeriod() string {
	var p periodType = firstPeriod
	return p.get()
}

// GetGroup returns the hydrogen group
func (Hydrogen) GetGroup() string {
	return ""
}

// GetCategory returns the hydrogen category
func (Hydrogen) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the hydrogen name
func (Hydrogen) GetName() string {
	return "Hydrogen"
}

// GetSimbol returns hydrogen simbol
func (Hydrogen) GetSimbol() string {
	return "H"
}

// GetAtomicNumber returns the hydrogen atomic number
func (Hydrogen) GetAtomicNumber() int {
	return 1
}

// GetAtomicWeight returns the hydrogen atomic weight
func (Hydrogen) GetAtomicWeight() float32 {
	return 1.008
}
