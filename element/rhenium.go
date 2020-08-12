package element

// Rhenium strtuct for rhenium element
type Rhenium struct{}

// GetPeriod returns the rhenium period
func (Rhenium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the rhenium group
func (Rhenium) GetGroup() string {
	var g groupType = b7
	return g.get()
}

// GetCategory returns the rhenium category
func (Rhenium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the rhenium name
func (Rhenium) GetName() string {
	return "Rhenium"
}

// GetSimbol returns rhenium simbol
func (Rhenium) GetSimbol() string {
	return "Re"
}

// GetAtomicNumber returns the rhenium atomic number
func (Rhenium) GetAtomicNumber() int {
	return 75
}

// GetAtomicWeight returns the rhenium atomic weight
func (Rhenium) GetAtomicWeight() float32 {
	return 186.21
}
