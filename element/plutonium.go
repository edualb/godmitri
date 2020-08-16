package element

// Plutonium strtuct for plutonium element
type Plutonium struct{}

// GetPeriod returns the plutonium period
func (Plutonium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the plutonium group
func (Plutonium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the plutonium category
func (Plutonium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the plutonium name
func (Plutonium) GetName() string {
	return "Plutonium"
}

// GetSimbol returns plutonium simbol
func (Plutonium) GetSimbol() string {
	return "Pu"
}

// GetAtomicNumber returns the plutonium atomic number
func (Plutonium) GetAtomicNumber() int {
	return 94
}

// GetAtomicWeight returns the plutonium atomic weight
func (Plutonium) GetAtomicWeight() float32 {
	return 244
}
