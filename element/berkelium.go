package element

// Berkelium strtuct for berkelium element
type Berkelium struct{}

// GetPeriod returns the berkelium period
func (Berkelium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the berkelium group
func (Berkelium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the berkelium category
func (Berkelium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the berkelium name
func (Berkelium) GetName() string {
	return "Berkelium"
}

// GetSimbol returns berkelium simbol
func (Berkelium) GetSimbol() string {
	return "Bk"
}

// GetAtomicNumber returns the berkelium atomic number
func (Berkelium) GetAtomicNumber() int {
	return 97
}

// GetAtomicWeight returns the berkelium atomic weight
func (Berkelium) GetAtomicWeight() float32 {
	return 247
}
