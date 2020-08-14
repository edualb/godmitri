package element

// Bismuth strtuct for bismuth element
type Bismuth struct{}

// GetPeriod returns the bismuth period
func (Bismuth) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the bismuth group
func (Bismuth) GetGroup() string {
	var g groupType = a5
	return g.get()
}

// GetCategory returns the bismuth category
func (Bismuth) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the bismuth name
func (Bismuth) GetName() string {
	return "Bismuth"
}

// GetSimbol returns bismuth simbol
func (Bismuth) GetSimbol() string {
	return "Bi"
}

// GetAtomicNumber returns the bismuth atomic number
func (Bismuth) GetAtomicNumber() int {
	return 83
}

// GetAtomicWeight returns the bismuth atomic weight
func (Bismuth) GetAtomicWeight() float32 {
	return 208.98
}
