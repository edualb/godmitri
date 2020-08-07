package element

// Strontium strtuct for strontium element
type Strontium struct{}

// GetPeriod returns the strontium period
func (Strontium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the strontium group
func (Strontium) GetGroup() string {
	var g groupType = a2
	return g.get()
}

// GetCategory returns the strontium category
func (Strontium) GetCategory() string {
	var c categoryType = alkalineEarthMetal
	return c.get()
}

// GetName returns the strontium name
func (Strontium) GetName() string {
	return "Strontium"
}

// GetSimbol returns strontium simbol
func (Strontium) GetSimbol() string {
	return "Sr"
}

// GetAtomicNumber returns the strontium atomic number
func (Strontium) GetAtomicNumber() int {
	return 38
}

// GetAtomicWeight returns the strontium atomic weight
func (Strontium) GetAtomicWeight() float32 {
	return 87.62
}
