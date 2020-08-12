package element

// Tantalum strtuct for tantalum element
type Tantalum struct{}

// GetPeriod returns the tantalum period
func (Tantalum) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the tantalum group
func (Tantalum) GetGroup() string {
	var g groupType = b5
	return g.get()
}

// GetCategory returns the tantalum category
func (Tantalum) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the tantalum name
func (Tantalum) GetName() string {
	return "Tantalum"
}

// GetSimbol returns tantalum simbol
func (Tantalum) GetSimbol() string {
	return "Ta"
}

// GetAtomicNumber returns the tantalum atomic number
func (Tantalum) GetAtomicNumber() int {
	return 73
}

// GetAtomicWeight returns the tantalum atomic weight
func (Tantalum) GetAtomicWeight() float32 {
	return 180.95
}
