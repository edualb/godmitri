package element

// Krypton strtuct for krypton element
type Krypton struct{}

// GetPeriod returns the krypton period
func (Krypton) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the krypton group
func (Krypton) GetGroup() string {
	var g groupType = a8
	return g.get()
}

// GetCategory returns the krypton category
func (Krypton) GetCategory() string {
	var c categoryType = nobleGas
	return c.get()
}

// GetName returns the krypton name
func (Krypton) GetName() string {
	return "Krypton"
}

// GetSimbol returns krypton simbol
func (Krypton) GetSimbol() string {
	return "Kr"
}

// GetAtomicNumber returns the krypton atomic number
func (Krypton) GetAtomicNumber() int {
	return 36
}

// GetAtomicWeight returns the krypton atomic weight
func (Krypton) GetAtomicWeight() float32 {
	return 83.798
}
