package element

// Boron strtuct for boron element
type Boron struct{}

// GetPeriod returns the boron period
func (Boron) GetPeriod() string {
	var p periodType = secondPeriod
	return p.get()
}

// GetGroup returns the boron group
func (Boron) GetGroup() string {
	var g groupType = a3
	return g.get()
}

// GetCategory returns the boron category
func (Boron) GetCategory() string {
	var c categoryType = metalloid
	return c.get()
}

// GetName returns the boron name
func (Boron) GetName() string {
	return "Boron"
}

// GetSimbol returns boron simbol
func (Boron) GetSimbol() string {
	return "B"
}

// GetAtomicNumber returns the boron atomic number
func (Boron) GetAtomicNumber() int {
	return 5
}

// GetAtomicWeight returns the boron atomic weight
func (Boron) GetAtomicWeight() float32 {
	return 10.81
}
