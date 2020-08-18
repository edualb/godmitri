package element

// Lawrencium strtuct for lawrencium element
type Lawrencium struct{}

// GetPeriod returns the lawrencium period
func (Lawrencium) GetPeriod() string {
	var p periodType = seventhPeriod
	return p.get()
}

// GetGroup returns the lawrencium group
func (Lawrencium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the lawrencium category
func (Lawrencium) GetCategory() string {
	var c categoryType = actinoid
	return c.get()
}

// GetName returns the lawrencium name
func (Lawrencium) GetName() string {
	return "Lawrencium"
}

// GetSimbol returns lawrencium simbol
func (Lawrencium) GetSimbol() string {
	return "Lr"
}

// GetAtomicNumber returns the lawrencium atomic number
func (Lawrencium) GetAtomicNumber() int {
	return 103
}

// GetAtomicWeight returns the lawrencium atomic weight
func (Lawrencium) GetAtomicWeight() float32 {
	return 266
}
