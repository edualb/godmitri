package element

// Moscovium strtuct for moscovium element
type Moscovium struct{}

// GetPeriod returns the moscovium period
func (Moscovium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the moscovium group
func (Moscovium) GetGroup() string {
	var g groupType = a5
	return g.get()
}

// GetCategory returns the moscovium category
func (Moscovium) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the moscovium name
func (Moscovium) GetName() string {
	return "Moscovium"
}

// GetSimbol returns moscovium simbol
func (Moscovium) GetSimbol() string {
	return "Mc"
}

// GetAtomicNumber returns the moscovium atomic number
func (Moscovium) GetAtomicNumber() int {
	return 115
}

// GetAtomicWeight returns the moscovium atomic weight
func (Moscovium) GetAtomicWeight() float32 {
	return 290
}
