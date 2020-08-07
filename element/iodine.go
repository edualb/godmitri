package element

// Iodine strtuct for iodine element
type Iodine struct{}

// GetPeriod returns the iodine period
func (Iodine) GetPeriod() string {
	var p periodType = fifthPeriod
	return p.get()
}

// GetGroup returns the iodine group
func (Iodine) GetGroup() string {
	var g groupType = a7
	return g.get()
}

// GetCategory returns the iodine category
func (Iodine) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the iodine name
func (Iodine) GetName() string {
	return "Iodine"
}

// GetSimbol returns iodine simbol
func (Iodine) GetSimbol() string {
	return "I"
}

// GetAtomicNumber returns the iodine atomic number
func (Iodine) GetAtomicNumber() int {
	return 53
}

// GetAtomicWeight returns the iodine atomic weight
func (Iodine) GetAtomicWeight() float32 {
	return 126.90
}
