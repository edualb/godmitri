package element

// Tennessine strtuct for tennessine element
type Tennessine struct{}

// GetPeriod returns the tennessine period
func (Tennessine) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the tennessine group
func (Tennessine) GetGroup() string {
	var g groupType = a7
	return g.get()
}

// GetCategory returns the tennessine category
func (Tennessine) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the tennessine name
func (Tennessine) GetName() string {
	return "Tennessine"
}

// GetSimbol returns tennessine simbol
func (Tennessine) GetSimbol() string {
	return "Ts"
}

// GetAtomicNumber returns the tennessine atomic number
func (Tennessine) GetAtomicNumber() int {
	return 117
}

// GetAtomicWeight returns the tennessine atomic weight
func (Tennessine) GetAtomicWeight() float32 {
	return 294
}
