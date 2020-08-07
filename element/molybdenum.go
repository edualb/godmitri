package element

// Molybdenum strtuct for molybdenum element
type Molybdenum struct{}

// GetPeriod returns the molybdenum period
func (Molybdenum) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the molybdenum group
func (Molybdenum) GetGroup() string {
	var g groupType = b6
	return g.get()
}

// GetCategory returns the molybdenum category
func (Molybdenum) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the molybdenum name
func (Molybdenum) GetName() string {
	return "Molybdenum"
}

// GetSimbol returns molybdenum simbol
func (Molybdenum) GetSimbol() string {
	return "Mo"
}

// GetAtomicNumber returns the molybdenum atomic number
func (Molybdenum) GetAtomicNumber() int {
	return 42
}

// GetAtomicWeight returns the molybdenum atomic weight
func (Molybdenum) GetAtomicWeight() float32 {
	return 95.95
}
