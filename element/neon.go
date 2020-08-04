package element

// Neon strtuct for neon element
type Neon struct{}

// GetPeriod returns the neon period
func (Neon) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the neon group
func (Neon) GetGroup() string {
	var g groupType = a8
	return g.get()
}

// GetCategory returns the neon category
func (Neon) GetCategory() string {
	var c categoryType = nobleGas
	return c.get()
}

// GetName returns the neon name
func (Neon) GetName() string {
	return "Neon"
}

// GetSimbol returns neon simbol
func (Neon) GetSimbol() string {
	return "Ne"
}

// GetAtomicNumber returns the neon atomic number
func (Neon) GetAtomicNumber() int {
	return 10
}

// GetAtomicWeight returns the neon atomic weight
func (Neon) GetAtomicWeight() float32 {
	return 20.180
}
