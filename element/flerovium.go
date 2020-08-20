package element

// Flerovium strtuct for flerovium element
type Flerovium struct{}

// GetPeriod returns the flerovium period
func (Flerovium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the flerovium group
func (Flerovium) GetGroup() string {
	var g groupType = a4
	return g.get()
}

// GetCategory returns the flerovium category
func (Flerovium) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the flerovium name
func (Flerovium) GetName() string {
	return "Flerovium"
}

// GetSimbol returns flerovium simbol
func (Flerovium) GetSimbol() string {
	return "Fl"
}

// GetAtomicNumber returns the flerovium atomic number
func (Flerovium) GetAtomicNumber() int {
	return 114
}

// GetAtomicWeight returns the flerovium atomic weight
func (Flerovium) GetAtomicWeight() float32 {
	return 289
}
