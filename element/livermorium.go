package element

// Livermorium strtuct for livermorium element
type Livermorium struct{}

// GetPeriod returns the livermorium period
func (Livermorium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the livermorium group
func (Livermorium) GetGroup() string {
	var g groupType = a6
	return g.get()
}

// GetCategory returns the livermorium category
func (Livermorium) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the livermorium name
func (Livermorium) GetName() string {
	return "Livermorium"
}

// GetSimbol returns livermorium simbol
func (Livermorium) GetSimbol() string {
	return "Lv"
}

// GetAtomicNumber returns the livermorium atomic number
func (Livermorium) GetAtomicNumber() int {
	return 116
}

// GetAtomicWeight returns the livermorium atomic weight
func (Livermorium) GetAtomicWeight() float32 {
	return 293
}
