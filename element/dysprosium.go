package element

// Dysprosium strtuct for dysprosium element
type Dysprosium struct{}

// GetPeriod returns the dysprosium period
func (Dysprosium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the dysprosium group
func (Dysprosium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the dysprosium category
func (Dysprosium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the dysprosium name
func (Dysprosium) GetName() string {
	return "Dysprosium"
}

// GetSimbol returns dysprosium simbol
func (Dysprosium) GetSimbol() string {
	return "Dy"
}

// GetAtomicNumber returns the dysprosium atomic number
func (Dysprosium) GetAtomicNumber() int {
	return 66
}

// GetAtomicWeight returns the dysprosium atomic weight
func (Dysprosium) GetAtomicWeight() float32 {
	return 157.25
}
