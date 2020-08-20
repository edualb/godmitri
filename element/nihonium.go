package element

// Nihonium strtuct for nihonium element
type Nihonium struct{}

// GetPeriod returns the nihonium period
func (Nihonium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the nihonium group
func (Nihonium) GetGroup() string {
	var g groupType = a3
	return g.get()
}

// GetCategory returns the nihonium category
func (Nihonium) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the nihonium name
func (Nihonium) GetName() string {
	return "Nihonium"
}

// GetSimbol returns nihonium simbol
func (Nihonium) GetSimbol() string {
	return "Nh"
}

// GetAtomicNumber returns the nihonium atomic number
func (Nihonium) GetAtomicNumber() int {
	return 113
}

// GetAtomicWeight returns the nihonium atomic weight
func (Nihonium) GetAtomicWeight() float32 {
	return 286
}
