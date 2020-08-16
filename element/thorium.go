package element

// Thorium strtuct for thorium element
type Thorium struct{}

// GetPeriod returns the thorium period
func (Thorium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the thorium group
func (Thorium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the thorium category
func (Thorium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the thorium name
func (Thorium) GetName() string {
	return "Thorium"
}

// GetSimbol returns thorium simbol
func (Thorium) GetSimbol() string {
	return "Th"
}

// GetAtomicNumber returns the thorium atomic number
func (Thorium) GetAtomicNumber() int {
	return 90
}

// GetAtomicWeight returns the thorium atomic weight
func (Thorium) GetAtomicWeight() float32 {
	return 232.04
}
