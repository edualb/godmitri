package element

// Tellurium strtuct for tellurium element
type Tellurium struct{}

// GetPeriod returns the tellurium period
func (Tellurium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the tellurium group
func (Tellurium) GetGroup() string {
	var g groupType = a6
	return g.get()
}

// GetCategory returns the tellurium category
func (Tellurium) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the tellurium name
func (Tellurium) GetName() string {
	return "Tellurium"
}

// GetSimbol returns tellurium simbol
func (Tellurium) GetSimbol() string {
	return "Te"
}

// GetAtomicNumber returns the tellurium atomic number
func (Tellurium) GetAtomicNumber() int {
	return 52
}

// GetAtomicWeight returns the tellurium atomic weight
func (Tellurium) GetAtomicWeight() float32 {
	return 127.60
}
