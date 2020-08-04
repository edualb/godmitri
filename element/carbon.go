package element

// Carbon strtuct for carbon element
type Carbon struct{}

// GetPeriod returns the carbon period
func (Carbon) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the carbon group
func (Carbon) GetGroup() string {
	var g groupType = a4
	return g.get()
}

// GetCategory returns the carbon category
func (Carbon) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the carbon name
func (Carbon) GetName() string {
	return "Carbon"
}

// GetSimbol returns carbon simbol
func (Carbon) GetSimbol() string {
	return "C"
}

// GetAtomicNumber returns the carbon atomic number
func (Carbon) GetAtomicNumber() int {
	return 6
}

// GetAtomicWeight returns the carbon atomic weight
func (Carbon) GetAtomicWeight() float32 {
	return 12.011
}
