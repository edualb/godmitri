package element

// Samarium strtuct for samarium element
type Samarium struct{}

// GetPeriod returns the samarium period
func (Samarium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the samarium group
func (Samarium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the samarium category
func (Samarium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the samarium name
func (Samarium) GetName() string {
	return "Samarium"
}

// GetSimbol returns samarium simbol
func (Samarium) GetSimbol() string {
	return "Sm"
}

// GetAtomicNumber returns the samarium atomic number
func (Samarium) GetAtomicNumber() int {
	return 62
}

// GetAtomicWeight returns the samarium atomic weight
func (Samarium) GetAtomicWeight() float32 {
	return 150.36
}
