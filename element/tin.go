package element

// Tin strtuct for tin element
type Tin struct{}

// GetPeriod returns the tin period
func (Tin) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the tin group
func (Tin) GetGroup() string {
	var g groupType = a4
	return g.get()
}

// GetCategory returns the tin category
func (Tin) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the tin name
func (Tin) GetName() string {
	return "Tin"
}

// GetSimbol returns tin simbol
func (Tin) GetSimbol() string {
	return "Sn"
}

// GetAtomicNumber returns the tin atomic number
func (Tin) GetAtomicNumber() int {
	return 50
}

// GetAtomicWeight returns the tin atomic weight
func (Tin) GetAtomicWeight() float32 {
	return 118.71
}
