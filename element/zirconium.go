package element

// Zirconium strtuct for zirconium element
type Zirconium struct{}

// GetPeriod returns the zirconium period
func (Zirconium) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the zirconium group
func (Zirconium) GetGroup() string {
	var g groupType = b4
	return g.get()
}

// GetCategory returns the zirconium category
func (Zirconium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the zirconium name
func (Zirconium) GetName() string {
	return "Zirconium"
}

// GetSimbol returns zirconium simbol
func (Zirconium) GetSimbol() string {
	return "Zr"
}

// GetAtomicNumber returns the zirconium atomic number
func (Zirconium) GetAtomicNumber() int {
	return 40
}

// GetAtomicWeight returns the zirconium atomic weight
func (Zirconium) GetAtomicWeight() float32 {
	return 91.224
}
