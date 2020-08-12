package element

// Platinum strtuct for platinum element
type Platinum struct{}

// GetPeriod returns the platinum period
func (Platinum) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the platinum group
func (Platinum) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the platinum category
func (Platinum) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the platinum name
func (Platinum) GetName() string {
	return "Platinum"
}

// GetSimbol returns platinum simbol
func (Platinum) GetSimbol() string {
	return "Pt"
}

// GetAtomicNumber returns the platinum atomic number
func (Platinum) GetAtomicNumber() int {
	return 78
}

// GetAtomicWeight returns the platinum atomic weight
func (Platinum) GetAtomicWeight() float32 {
	return 195.08
}
