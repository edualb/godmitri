package element

// Arsenic strtuct for arsenic element
type Arsenic struct{}

// GetPeriod returns the arsenic period
func (Arsenic) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the arsenic group
func (Arsenic) GetGroup() string {
	var g groupType = a5
	return g.get()
}

// GetCategory returns the arsenic category
func (Arsenic) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the arsenic name
func (Arsenic) GetName() string {
	return "Arsenic"
}

// GetSimbol returns arsenic simbol
func (Arsenic) GetSimbol() string {
	return "As"
}

// GetAtomicNumber returns the arsenic atomic number
func (Arsenic) GetAtomicNumber() int {
	return 33
}

// GetAtomicWeight returns the arsenic atomic weight
func (Arsenic) GetAtomicWeight() float32 {
	return 74.922
}
