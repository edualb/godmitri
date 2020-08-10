package element

// Holmium strtuct for holmium element
type Holmium struct{}

// GetPeriod returns the holmium period
func (Holmium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the holmium group
func (Holmium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the holmium category
func (Holmium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the holmium name
func (Holmium) GetName() string {
	return "Holmium"
}

// GetSimbol returns holmium simbol
func (Holmium) GetSimbol() string {
	return "Ho"
}

// GetAtomicNumber returns the holmium atomic number
func (Holmium) GetAtomicNumber() int {
	return 67
}

// GetAtomicWeight returns the holmium atomic weight
func (Holmium) GetAtomicWeight() float32 {
	return 164.93
}
