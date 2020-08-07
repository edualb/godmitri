package element

// Silver strtuct for silver element
type Silver struct{}

// GetPeriod returns the silver period
func (Silver) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the silver group
func (Silver) GetGroup() string {
	var g groupType = b1
	return g.get()
}

// GetCategory returns the silver category
func (Silver) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the silver name
func (Silver) GetName() string {
	return "Silver"
}

// GetSimbol returns silver simbol
func (Silver) GetSimbol() string {
	return "Ag"
}

// GetAtomicNumber returns the silver atomic number
func (Silver) GetAtomicNumber() int {
	return 47
}

// GetAtomicWeight returns the silver atomic weight
func (Silver) GetAtomicWeight() float32 {
	return 107.87
}
