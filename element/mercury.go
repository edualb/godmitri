package element

// Mercury strtuct for mercury element
type Mercury struct{}

// GetPeriod returns the mercury period
func (Mercury) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the mercury group
func (Mercury) GetGroup() string {
	var g groupType = b2
	return g.get()
}

// GetCategory returns the mercury category
func (Mercury) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the mercury name
func (Mercury) GetName() string {
	return "Mercury"
}

// GetSimbol returns mercury simbol
func (Mercury) GetSimbol() string {
	return "Hg"
}

// GetAtomicNumber returns the mercury atomic number
func (Mercury) GetAtomicNumber() int {
	return 80
}

// GetAtomicWeight returns the mercury atomic weight
func (Mercury) GetAtomicWeight() float32 {
	return 200.59
}
