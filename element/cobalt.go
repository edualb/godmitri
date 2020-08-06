package element

// Cobalt strtuct for cobalt element
type Cobalt struct{}

// GetPeriod returns the cobalt period
func (Cobalt) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the cobalt group
func (Cobalt) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the cobalt category
func (Cobalt) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the cobalt name
func (Cobalt) GetName() string {
	return "Cobalt"
}

// GetSimbol returns cobalt simbol
func (Cobalt) GetSimbol() string {
	return "Co"
}

// GetAtomicNumber returns the cobalt atomic number
func (Cobalt) GetAtomicNumber() int {
	return 27
}

// GetAtomicWeight returns the cobalt atomic weight
func (Cobalt) GetAtomicWeight() float32 {
	return 58.933
}
