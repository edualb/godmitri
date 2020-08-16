package element

// Neptunium strtuct for neptunium element
type Neptunium struct{}

// GetPeriod returns the neptunium period
func (Neptunium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the neptunium group
func (Neptunium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the neptunium category
func (Neptunium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the neptunium name
func (Neptunium) GetName() string {
	return "Neptunium"
}

// GetSimbol returns neptunium simbol
func (Neptunium) GetSimbol() string {
	return "Np"
}

// GetAtomicNumber returns the neptunium atomic number
func (Neptunium) GetAtomicNumber() int {
	return 93
}

// GetAtomicWeight returns the neptunium atomic weight
func (Neptunium) GetAtomicWeight() float32 {
	return 237
}
