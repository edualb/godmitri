package element

// Rubidium strtuct for rubidium element
type Rubidium struct{}

// GetPeriod returns the rubidium period
func (Rubidium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the rubidium group
func (Rubidium) GetGroup() string {
	var g groupType = a1
	return g.get()
}

// GetCategory returns the rubidium category
func (Rubidium) GetCategory() string {
	var c categoryType = alkaliMetal
	return c.get()
}

// GetName returns the rubidium name
func (Rubidium) GetName() string {
	return "Rubidium"
}

// GetSimbol returns rubidium simbol
func (Rubidium) GetSimbol() string {
	return "Rb"
}

// GetAtomicNumber returns the rubidium atomic number
func (Rubidium) GetAtomicNumber() int {
	return 37
}

// GetAtomicWeight returns the rubidium atomic weight
func (Rubidium) GetAtomicWeight() float32 {
	return 85.468
}
