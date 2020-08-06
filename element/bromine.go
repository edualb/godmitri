package element

// Bromine strtuct for bromine element
type Bromine struct{}

// GetPeriod returns the bromine period
func (Bromine) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the bromine group
func (Bromine) GetGroup() string {
	var g groupType = a7
	return g.get()
}

// GetCategory returns the bromine category
func (Bromine) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the bromine name
func (Bromine) GetName() string {
	return "Bromine"
}

// GetSimbol returns bromine simbol
func (Bromine) GetSimbol() string {
	return "Br"
}

// GetAtomicNumber returns the bromine atomic number
func (Bromine) GetAtomicNumber() int {
	return 35
}

// GetAtomicWeight returns the bromine atomic weight
func (Bromine) GetAtomicWeight() float32 {
	return 79.904
}
