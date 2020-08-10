package element

// Promethium strtuct for promethium element
type Promethium struct{}

// GetPeriod returns the promethium period
func (Promethium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the promethium group
func (Promethium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the promethium category
func (Promethium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the promethium name
func (Promethium) GetName() string {
	return "Promethium"
}

// GetSimbol returns promethium simbol
func (Promethium) GetSimbol() string {
	return "Pm"
}

// GetAtomicNumber returns the promethium atomic number
func (Promethium) GetAtomicNumber() int {
	return 61
}

// GetAtomicWeight returns the promethium atomic weight
func (Promethium) GetAtomicWeight() float32 {
	return 145
}
