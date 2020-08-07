package element

// Xenon strtuct for xenon element
type Xenon struct{}

// GetPeriod returns the xenon period
func (Xenon) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the xenon group
func (Xenon) GetGroup() string {
	var g groupType = a8
	return g.get()
}

// GetCategory returns the xenon category
func (Xenon) GetCategory() string {
	var c categoryType = nobleGas
	return c.get()
}

// GetName returns the xenon name
func (Xenon) GetName() string {
	return "Xenon"
}

// GetSimbol returns xenon simbol
func (Xenon) GetSimbol() string {
	return "Xe"
}

// GetAtomicNumber returns the xenon atomic number
func (Xenon) GetAtomicNumber() int {
	return 54
}

// GetAtomicWeight returns the xenon atomic weight
func (Xenon) GetAtomicWeight() float32 {
	return 131.29
}
