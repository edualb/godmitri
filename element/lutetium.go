package element

// Lutetium strtuct for lutetium element
type Lutetium struct{}

// GetPeriod returns the lutetium period
func (Lutetium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the lutetium group
func (Lutetium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the lutetium category
func (Lutetium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the lutetium name
func (Lutetium) GetName() string {
	return "Lutetium"
}

// GetSimbol returns lutetium simbol
func (Lutetium) GetSimbol() string {
	return "Lu"
}

// GetAtomicNumber returns the lutetium atomic number
func (Lutetium) GetAtomicNumber() int {
	return 71
}

// GetAtomicWeight returns the lutetium atomic weight
func (Lutetium) GetAtomicWeight() float32 {
	return 174.97
}
