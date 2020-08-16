package element

// Radium strtuct for radium element
type Radium struct{}

// GetPeriod returns the radium period
func (Radium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the radium group
func (Radium) GetGroup() string {
	var g groupType = a2
	return g.get()
}

// GetCategory returns the radium category
func (Radium) GetCategory() string {
	var c categoryType = alkalineEarthMetal
	return c.get()
}

// GetName returns the radium name
func (Radium) GetName() string {
	return "Radium"
}

// GetSimbol returns radium simbol
func (Radium) GetSimbol() string {
	return "Ra"
}

// GetAtomicNumber returns the radium atomic number
func (Radium) GetAtomicNumber() int {
	return 88
}

// GetAtomicWeight returns the radium atomic weight
func (Radium) GetAtomicWeight() float32 {
	return 226
}
