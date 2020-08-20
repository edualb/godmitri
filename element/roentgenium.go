package element

// Roentgenium strtuct for roentgenium element
type Roentgenium struct{}

// GetPeriod returns the roentgenium period
func (Roentgenium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the roentgenium group
func (Roentgenium) GetGroup() string {
	var g groupType = b1
	return g.get()
}

// GetCategory returns the roentgenium category
func (Roentgenium) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the roentgenium name
func (Roentgenium) GetName() string {
	return "Roentgenium"
}

// GetSimbol returns roentgenium simbol
func (Roentgenium) GetSimbol() string {
	return "Rg"
}

// GetAtomicNumber returns the roentgenium atomic number
func (Roentgenium) GetAtomicNumber() int {
	return 111
}

// GetAtomicWeight returns the roentgenium atomic weight
func (Roentgenium) GetAtomicWeight() float32 {
	return 282
}
