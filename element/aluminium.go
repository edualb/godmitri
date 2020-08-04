package element

// Aluminium strtuct for aluminium element
type Aluminium struct{}

// GetPeriod returns the aluminium period
func (Aluminium) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the aluminium group
func (Aluminium) GetGroup() string {
	var g groupType = a3
	return g.get()
}

// GetCategory returns the aluminium category
func (Aluminium) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the aluminium name
func (Aluminium) GetName() string {
	return "Aluminium"
}

// GetSimbol returns aluminium simbol
func (Aluminium) GetSimbol() string {
	return "Al"
}

// GetAtomicNumber returns the aluminium atomic number
func (Aluminium) GetAtomicNumber() int {
	return 13
}

// GetAtomicWeight returns the aluminium atomic weight
func (Aluminium) GetAtomicWeight() float32 {
	return 26.982
}
