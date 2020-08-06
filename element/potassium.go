package element

// Potassium strtuct for potassium element
type Potassium struct{}

// GetPeriod returns the potassium period
func (Potassium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the potassium group
func (Potassium) GetGroup() string {
	var g groupType = a1
	return g.get()
}

// GetCategory returns the potassium category
func (Potassium) GetCategory() string {
	var c categoryType = alkaliMetal
	return c.get()
}

// GetName returns the potassium name
func (Potassium) GetName() string {
	return "Potassium"
}

// GetSimbol returns potassium simbol
func (Potassium) GetSimbol() string {
	return "K"
}

// GetAtomicNumber returns the potassium atomic number
func (Potassium) GetAtomicNumber() int {
	return 19
}

// GetAtomicWeight returns the potassium atomic weight
func (Potassium) GetAtomicWeight() float32 {
	return 39.098
}
