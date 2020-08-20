package element

// Oganesson strtuct for oganesson element
type Oganesson struct{}

// GetPeriod returns the oganesson period
func (Oganesson) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the oganesson group
func (Oganesson) GetGroup() string {
	var g groupType = a8
	return g.get()
}

// GetCategory returns the oganesson category
func (Oganesson) GetCategory() string {
	var c categoryType = unknown
	return c.get()
}

// GetName returns the oganesson name
func (Oganesson) GetName() string {
	return "Oganesson"
}

// GetSimbol returns oganesson simbol
func (Oganesson) GetSimbol() string {
	return "Og"
}

// GetAtomicNumber returns the oganesson atomic number
func (Oganesson) GetAtomicNumber() int {
	return 118
}

// GetAtomicWeight returns the oganesson atomic weight
func (Oganesson) GetAtomicWeight() float32 {
	return 294
}
