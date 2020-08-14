package element

// Lead strtuct for lead element
type Lead struct{}

// GetPeriod returns the lead period
func (Lead) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the lead group
func (Lead) GetGroup() string {
	var g groupType = a4
	return g.get()
}

// GetCategory returns the lead category
func (Lead) GetCategory() string {
	var c categoryType = postTransitionMetal
	return c.get()
}

// GetName returns the lead name
func (Lead) GetName() string {
	return "Lead"
}

// GetSimbol returns lead simbol
func (Lead) GetSimbol() string {
	return "Pb"
}

// GetAtomicNumber returns the lead atomic number
func (Lead) GetAtomicNumber() int {
	return 82
}

// GetAtomicWeight returns the lead atomic weight
func (Lead) GetAtomicWeight() float32 {
	return 207.2
}
