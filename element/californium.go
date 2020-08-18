package element

// Californium strtuct for californium element
type Californium struct{}

// GetPeriod returns the californium period
func (Californium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the californium group
func (Californium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the californium category
func (Californium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the californium name
func (Californium) GetName() string {
	return "Californium"
}

// GetSimbol returns californium simbol
func (Californium) GetSimbol() string {
	return "Cf"
}

// GetAtomicNumber returns the californium atomic number
func (Californium) GetAtomicNumber() int {
	return 98
}

// GetAtomicWeight returns the californium atomic weight
func (Californium) GetAtomicWeight() float32 {
	return 251
}
