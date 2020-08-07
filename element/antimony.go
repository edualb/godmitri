package element

// Antimony strtuct for antimony element
type Antimony struct{}

// GetPeriod returns the antimony period
func (Antimony) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the antimony group
func (Antimony) GetGroup() string {
	var g groupType = a5
	return g.get()
}

// GetCategory returns the antimony category
func (Antimony) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the antimony name
func (Antimony) GetName() string {
	return "Antimony"
}

// GetSimbol returns antimony simbol
func (Antimony) GetSimbol() string {
	return "Sb"
}

// GetAtomicNumber returns the antimony atomic number
func (Antimony) GetAtomicNumber() int {
	return 51
}

// GetAtomicWeight returns the antimony atomic weight
func (Antimony) GetAtomicWeight() float32 {
	return 121.76
}
