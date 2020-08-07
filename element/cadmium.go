package element

// Cadmium strtuct for cadmium element
type Cadmium struct{}

// GetPeriod returns the cadmium period
func (Cadmium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the cadmium group
func (Cadmium) GetGroup() string {
	var g groupType = b2
	return g.get()
}

// GetCategory returns the cadmium category
func (Cadmium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the cadmium name
func (Cadmium) GetName() string {
	return "Cadmium"
}

// GetSimbol returns cadmium simbol
func (Cadmium) GetSimbol() string {
	return "Cd"
}

// GetAtomicNumber returns the cadmium atomic number
func (Cadmium) GetAtomicNumber() int {
	return 48
}

// GetAtomicWeight returns the cadmium atomic weight
func (Cadmium) GetAtomicWeight() float32 {
	return 112.41
}
