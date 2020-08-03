package element

// Helium strtuct for helium element
type Helium struct{}

// GetPeriod returns the helium period
func (Helium) GetPeriod() string {
	var p periodType = firstPeriod
	return p.get()
}

// GetGroup returns the helium group
func (Helium) GetGroup() string {
	return "8A"
}

// GetCategory returns the helium category
func (Helium) GetCategory() string {
	return "noble gas"
}

// GetName returns the helium name
func (Helium) GetName() string {
	return "Helium"
}

// GetSimbol returns helium simbol
func (Helium) GetSimbol() string {
	return "He"
}

// GetAtomicNumber returns the helium atomic number
func (Helium) GetAtomicNumber() int {
	return 2
}

// GetAtomicWeight returns the helium atomic weight
func (Helium) GetAtomicWeight() float32 {
	return 4.0026
}
