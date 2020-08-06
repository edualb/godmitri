package element

// Iron strtuct for iron element
type Iron struct{}

// GetPeriod returns the iron period
func (Iron) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the iron group
func (Iron) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the iron category
func (Iron) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the iron name
func (Iron) GetName() string {
	return "Iron"
}

// GetSimbol returns iron simbol
func (Iron) GetSimbol() string {
	return "Fe"
}

// GetAtomicNumber returns the iron atomic number
func (Iron) GetAtomicNumber() int {
	return 26
}

// GetAtomicWeight returns the iron atomic weight
func (Iron) GetAtomicWeight() float32 {
	return 55.845
}
