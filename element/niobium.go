package element

// Niobium strtuct for niobium element
type Niobium struct{}

// GetPeriod returns the niobium period
func (Niobium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the niobium group
func (Niobium) GetGroup() string {
	var g groupType = b5
	return g.get()
}

// GetCategory returns the niobium category
func (Niobium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the niobium name
func (Niobium) GetName() string {
	return "Niobium"
}

// GetSimbol returns niobium simbol
func (Niobium) GetSimbol() string {
	return "Nb"
}

// GetAtomicNumber returns the niobium atomic number
func (Niobium) GetAtomicNumber() int {
	return 41
}

// GetAtomicWeight returns the niobium atomic weight
func (Niobium) GetAtomicWeight() float32 {
	return 92.906
}
