package element

// Mendelevium strtuct for mendelevium element
type Mendelevium struct{}

// GetPeriod returns the mendelevium period
func (Mendelevium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the mendelevium group
func (Mendelevium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the mendelevium category
func (Mendelevium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the mendelevium name
func (Mendelevium) GetName() string {
	return "Mendelevium"
}

// GetSimbol returns mendelevium simbol
func (Mendelevium) GetSimbol() string {
	return "Md"
}

// GetAtomicNumber returns the mendelevium atomic number
func (Mendelevium) GetAtomicNumber() int {
	return 101
}

// GetAtomicWeight returns the mendelevium atomic weight
func (Mendelevium) GetAtomicWeight() float32 {
	return 258
}
