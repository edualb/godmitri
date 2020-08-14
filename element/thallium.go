package element

// Thallium strtuct for thallium element
type Thallium struct{}

// GetPeriod returns the thallium period
func (Thallium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the thallium group
func (Thallium) GetGroup() string {
	var g groupType = a3
	return g.get()
}

// GetCategory returns the thallium category
func (Thallium) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the thallium name
func (Thallium) GetName() string {
	return "Thallium"
}

// GetSimbol returns thallium simbol
func (Thallium) GetSimbol() string {
	return "Tl"
}

// GetAtomicNumber returns the terbium atomic number
func (Thallium) GetAtomicNumber() int {
	return 81
}

// GetAtomicWeight returns the thallium atomic weight
func (Thallium) GetAtomicWeight() float32 {
	return 204.38
}
