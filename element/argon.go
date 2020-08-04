package element

// Argon strtuct for argon element
type Argon struct{}

// GetPeriod returns the argon period
func (Argon) GetPeriod() string {
	var p periodType = thirdPeriod
	return p.get()
}

// GetGroup returns the argon group
func (Argon) GetGroup() string {
	var g groupType = a8
	return g.get()
}

// GetCategory returns the argon category
func (Argon) GetCategory() string {
	var c categoryType = nobleGas
	return c.get()
}

// GetName returns the argon name
func (Argon) GetName() string {
	return "Argon"
}

// GetSimbol returns argon simbol
func (Argon) GetSimbol() string {
	return "Ar"
}

// GetAtomicNumber returns the argon atomic number
func (Argon) GetAtomicNumber() int {
	return 18
}

// GetAtomicWeight returns the argon atomic weight
func (Argon) GetAtomicWeight() float32 {
	return 39.948
}
