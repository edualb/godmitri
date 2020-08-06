package element

// Zinc strtuct for zinc element
type Zinc struct{}

// GetPeriod returns the zinc period
func (Zinc) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the zinc group
func (Zinc) GetGroup() string {
	var g groupType = b2
	return g.get()
}

// GetCategory returns the zinc category
func (Zinc) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the zinc name
func (Zinc) GetName() string {
	return "Zinc"
}

// GetSimbol returns zinc simbol
func (Zinc) GetSimbol() string {
	return "Zn"
}

// GetAtomicNumber returns the zinc atomic number
func (Zinc) GetAtomicNumber() int {
	return 30
}

// GetAtomicWeight returns the zinc atomic weight
func (Zinc) GetAtomicWeight() float32 {
	return 65.38
}
