package element

// Silicon strtuct for silicon element
type Silicon struct{}

// GetPeriod returns the silicon period
func (Silicon) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the silicon group
func (Silicon) GetGroup() string {
	var g groupType = a4
	return g.get()
}

// GetCategory returns the silicon category
func (Silicon) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the silicon name
func (Silicon) GetName() string {
	return "Silicon"
}

// GetSimbol returns silicon simbol
func (Silicon) GetSimbol() string {
	return "Si"
}

// GetAtomicNumber returns the silicon atomic number
func (Silicon) GetAtomicNumber() int {
	return 14
}

// GetAtomicWeight returns the silicon atomic weight
func (Silicon) GetAtomicWeight() float32 {
	return 28.085
}
