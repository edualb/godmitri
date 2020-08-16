package element

// Actinium strtuct for actinium element
type Actinium struct{}

// GetPeriod returns the actinium period
func (Actinium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the actinium group
func (Actinium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the actinium category
func (Actinium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the actinium name
func (Actinium) GetName() string {
	return "Actinium"
}

// GetSimbol returns actinium simbol
func (Actinium) GetSimbol() string {
	return "Ac"
}

// GetAtomicNumber returns the actinium atomic number
func (Actinium) GetAtomicNumber() int {
	return 89
}

// GetAtomicWeight returns the actinium atomic weight
func (Actinium) GetAtomicWeight() float32 {
	return 227
}
