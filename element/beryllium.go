package element

// Beryllium strtuct for beryllium element
type Beryllium struct{}

// GetPeriod returns the beryllium period
func (Beryllium) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the beryllium group
func (Beryllium) GetGroup() string {
	var g groupType = a2
	return g.get()
}

// GetCategory returns the beryllium category
func (Beryllium) GetCategory() string {
	var c categoryType = alkalineEarthMetal
	return c.get()
}

// GetName returns the beryllium name
func (Beryllium) GetName() string {
	return "Beryllium"
}

// GetSimbol returns beryllium simbol
func (Beryllium) GetSimbol() string {
	return "Be"
}

// GetAtomicNumber returns the beryllium atomic number
func (Beryllium) GetAtomicNumber() int {
	return 4
}

// GetAtomicWeight returns the beryllium atomic weight
func (Beryllium) GetAtomicWeight() float32 {
	return 9.0122
}
