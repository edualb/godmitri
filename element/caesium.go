package element

// Caesium strtuct for caesium element
type Caesium struct{}

// GetPeriod returns the caesium period
func (Caesium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the caesium group
func (Caesium) GetGroup() string {
	var g groupType = a1
	return g.get()
}

// GetCategory returns the caesium category
func (Caesium) GetCategory() string {
	var c categoryType = alkaliMetal
	return c.get()
}

// GetName returns the caesium name
func (Caesium) GetName() string {
	return "Caesium"
}

// GetSimbol returns caesium simbol
func (Caesium) GetSimbol() string {
	return "Cs"
}

// GetAtomicNumber returns the caesium atomic number
func (Caesium) GetAtomicNumber() int {
	return 55
}

// GetAtomicWeight returns the caesium atomic weight
func (Caesium) GetAtomicWeight() float32 {
	return 132.91
}
