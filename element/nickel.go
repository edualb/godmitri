package element

// Nickel strtuct for nickel element
type Nickel struct{}

// GetPeriod returns the nickel period
func (Nickel) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the nickel group
func (Nickel) GetGroup() string {
	var g groupType = b8
	return g.get()
}

// GetCategory returns the nickel category
func (Nickel) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the nickel name
func (Nickel) GetName() string {
	return "Nickel"
}

// GetSimbol returns nickel simbol
func (Nickel) GetSimbol() string {
	return "Ni"
}

// GetAtomicNumber returns the nickel atomic number
func (Nickel) GetAtomicNumber() int {
	return 28
}

// GetAtomicWeight returns the nickel atomic weight
func (Nickel) GetAtomicWeight() float32 {
	return 58.693
}
