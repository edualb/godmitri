package element

// Darmstadtium strtuct for darmstadtium element
type Darmstadtium struct{}

// GetPeriod returns the darmstadtium period
func (Darmstadtium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the darmstadtium group
func (Darmstadtium) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the darmstadtium category
func (Darmstadtium) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the darmstadtium name
func (Darmstadtium) GetName() string {
	return "Darmstadtium"
}

// GetSimbol returns darmstadtium simbol
func (Darmstadtium) GetSimbol() string {
	return "Ds"
}

// GetAtomicNumber returns the darmstadtium atomic number
func (Darmstadtium) GetAtomicNumber() int {
	return 110
}

// GetAtomicWeight returns the darmstadtium atomic weight
func (Darmstadtium) GetAtomicWeight() float32 {
	return 281
}
