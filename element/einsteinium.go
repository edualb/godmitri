package element

// Einsteinium strtuct for einsteinium element
type Einsteinium struct{}

// GetPeriod returns the einsteinium period
func (Einsteinium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the einsteinium group
func (Einsteinium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the einsteinium category
func (Einsteinium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the einsteinium name
func (Einsteinium) GetName() string {
	return "Einsteinium"
}

// GetSimbol returns einsteinium simbol
func (Einsteinium) GetSimbol() string {
	return "Es"
}

// GetAtomicNumber returns the einsteinium atomic number
func (Einsteinium) GetAtomicNumber() int {
	return 99
}

// GetAtomicWeight returns the einsteinium atomic weight
func (Einsteinium) GetAtomicWeight() float32 {
	return 252
}
