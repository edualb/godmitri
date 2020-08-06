package element

// Calcium strtuct for calcium element
type Calcium struct{}

// GetPeriod returns the calcium period
func (Calcium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the calcium group
func (Calcium) GetGroup() string {
	var g groupType = a2
	return g.get()
}

// GetCategory returns the calcium category
func (Calcium) GetCategory() string {
	var c categoryType = alkalineEarthMetal
	return c.get()
}

// GetName returns the calcium name
func (Calcium) GetName() string {
	return "Calcium"
}

// GetSimbol returns calcium simbol
func (Calcium) GetSimbol() string {
	return "Ca"
}

// GetAtomicNumber returns the calcium atomic number
func (Calcium) GetAtomicNumber() int {
	return 20
}

// GetAtomicWeight returns the calcium atomic weight
func (Calcium) GetAtomicWeight() float32 {
	return 40.078
}
