package element

// Radon strtuct for radon element
type Radon struct{}

// GetPeriod returns the radon period
func (Radon) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the radon group
func (Radon) GetGroup() string {
	var g groupType = a8
	return g.get()
}

// GetCategory returns the radon category
func (Radon) GetCategory() string {
	var c categoryType = nobleGas
	return c.get()
}

// GetName returns the radon name
func (Radon) GetName() string {
	return "Radon"
}

// GetSimbol returns radon simbol
func (Radon) GetSimbol() string {
	return "Rn"
}

// GetAtomicNumber returns the radon atomic number
func (Radon) GetAtomicNumber() int {
	return 86
}

// GetAtomicWeight returns the radon atomic weight
func (Radon) GetAtomicWeight() float32 {
	return 222
}
