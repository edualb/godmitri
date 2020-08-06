package element

// Germanium strtuct for germanium element
type Germanium struct{}

// GetPeriod returns the germanium period
func (Germanium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the germanium group
func (Germanium) GetGroup() string {
	var g groupType = a4
	return g.get()
}

// GetCategory returns the germanium category
func (Germanium) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the germanium name
func (Germanium) GetName() string {
	return "Germanium"
}

// GetSimbol returns germanium simbol
func (Germanium) GetSimbol() string {
	return "Ge"
}

// GetAtomicNumber returns the germanium atomic number
func (Germanium) GetAtomicNumber() int {
	return 32
}

// GetAtomicWeight returns the germanium atomic weight
func (Germanium) GetAtomicWeight() float32 {
	return 72.630
}
