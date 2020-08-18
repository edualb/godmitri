package element

// Seaborgium strtuct for seaborgium element
type Seaborgium struct{}

// GetPeriod returns the seaborgium period
func (Seaborgium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the seaborgium group
func (Seaborgium) GetGroup() string {
	var g groupType = b6
	return g.get()
}

// GetCategory returns the seaborgium category
func (Seaborgium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the seaborgium name
func (Seaborgium) GetName() string {
	return "Seaborgium"
}

// GetSimbol returns seaborgium simbol
func (Seaborgium) GetSimbol() string {
	return "Sg"
}

// GetAtomicNumber returns the seaborgium atomic number
func (Seaborgium) GetAtomicNumber() int {
	return 106
}

// GetAtomicWeight returns the seaborgium atomic weight
func (Seaborgium) GetAtomicWeight() float32 {
	return 269
}
