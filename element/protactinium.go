package element

// Protactinium strtuct for protactinium element
type Protactinium struct{}

// GetPeriod returns the protactinium period
func (Protactinium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the protactinium group
func (Protactinium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the protactinium category
func (Protactinium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the protactinium name
func (Protactinium) GetName() string {
	return "Protactinium"
}

// GetSimbol returns protactinium simbol
func (Protactinium) GetSimbol() string {
	return "Pa"
}

// GetAtomicNumber returns the protactinium atomic number
func (Protactinium) GetAtomicNumber() int {
	return 91
}

// GetAtomicWeight returns the protactinium atomic weight
func (Protactinium) GetAtomicWeight() float32 {
	return 231.04
}
