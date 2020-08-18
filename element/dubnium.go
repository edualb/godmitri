package element

// Dubnium strtuct for dubnium element
type Dubnium struct{}

// GetPeriod returns the dubnium period
func (Dubnium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the dubnium group
func (Dubnium) GetGroup() string {
	var g groupType = b5
	return g.get()
}

// GetCategory returns the dubnium category
func (Dubnium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the dubnium name
func (Dubnium) GetName() string {
	return "Dubnium"
}

// GetSimbol returns dubnium simbol
func (Dubnium) GetSimbol() string {
	return "Db"
}

// GetAtomicNumber returns the dubnium atomic number
func (Dubnium) GetAtomicNumber() int {
	return 105
}

// GetAtomicWeight returns the dubnium atomic weight
func (Dubnium) GetAtomicWeight() float32 {
	return 268
}
