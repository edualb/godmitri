package element

// Lanthanum strtuct for lanthanum element
type Lanthanum struct{}

// GetPeriod returns the lanthanum period
func (Lanthanum) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the lanthanum group
func (Lanthanum) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the lanthanum category
func (Lanthanum) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the lanthanum name
func (Lanthanum) GetName() string {
	return "Lanthanum"
}

// GetSimbol returns lanthanum simbol
func (Lanthanum) GetSimbol() string {
	return "La"
}

// GetAtomicNumber returns the lanthanum atomic number
func (Lanthanum) GetAtomicNumber() int {
	return 57
}

// GetAtomicWeight returns the lanthanum atomic weight
func (Lanthanum) GetAtomicWeight() float32 {
	return 138.91
}
