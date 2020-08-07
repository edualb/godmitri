package element

// Barium strtuct for barium element
type Barium struct{}

// GetPeriod returns the barium period
func (Barium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the barium group
func (Barium) GetGroup() string {
	var g groupType = a2
	return g.get()
}

// GetCategory returns the barium category
func (Barium) GetCategory() string {
	var c categoryType = alkalineEarthMetal
	return c.get()
}

// GetName returns the barium name
func (Barium) GetName() string {
	return "Barium"
}

// GetSimbol returns barium simbol
func (Barium) GetSimbol() string {
	return "Ba"
}

// GetAtomicNumber returns the barium atomic number
func (Barium) GetAtomicNumber() int {
	return 56
}

// GetAtomicWeight returns the barium atomic weight
func (Barium) GetAtomicWeight() float32 {
	return 137.33
}
