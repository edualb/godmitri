package element

// Titanium strtuct for titanium element
type Titanium struct{}

// GetPeriod returns the titanium period
func (Titanium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the titanium group
func (Titanium) GetGroup() string {
	var g groupType = b4
	return g.get()
}

// GetCategory returns the titanium category
func (Titanium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the titanium name
func (Titanium) GetName() string {
	return "Titanium"
}

// GetSimbol returns titanium simbol
func (Titanium) GetSimbol() string {
	return "Ti"
}

// GetAtomicNumber returns the titanium atomic number
func (Titanium) GetAtomicNumber() int {
	return 22
}

// GetAtomicWeight returns the titanium atomic weight
func (Titanium) GetAtomicWeight() float32 {
	return 47.867
}
