package element

// Gadolinium strtuct for gadolinium element
type Gadolinium struct{}

// GetPeriod returns the gadolinium period
func (Gadolinium) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the gadolinium group
func (Gadolinium) GetGroup() string {
	var g groupType = b3
	return g.get()
}

// GetCategory returns the gadolinium category
func (Gadolinium) GetCategory() string {
	var c categoryType = lanthanoid
	return c.get()
}

// GetName returns the gadolinium name
func (Gadolinium) GetName() string {
	return "Gadolinium"
}

// GetSimbol returns gadolinium simbol
func (Gadolinium) GetSimbol() string {
	return "Gd"
}

// GetAtomicNumber returns the gadolinium atomic number
func (Gadolinium) GetAtomicNumber() int {
	return 64
}

// GetAtomicWeight returns the gadolinium atomic weight
func (Gadolinium) GetAtomicWeight() float32 {
	return 157.25
}
