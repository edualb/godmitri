package element

// Technetium strtuct for technetium element
type Technetium struct{}

// GetPeriod returns the technetium period
func (Technetium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the technetium group
func (Technetium) GetGroup() string {
	var g groupType = b7
	return g.get()
}

// GetCategory returns the technetium category
func (Technetium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the technetium name
func (Technetium) GetName() string {
	return "Technetium"
}

// GetSimbol returns technetium simbol
func (Technetium) GetSimbol() string {
	return "Tc"
}

// GetAtomicNumber returns the technetium atomic number
func (Technetium) GetAtomicNumber() int {
	return 43
}

// GetAtomicWeight returns the technetium atomic weight
func (Technetium) GetAtomicWeight() float32 {
	return 98
}
