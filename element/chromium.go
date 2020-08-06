package element

// Chromium strtuct for chromium element
type Chromium struct{}

// GetPeriod returns the chromium period
func (Chromium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the chromium group
func (Chromium) GetGroup() string {
	var g groupType = b6
	return g.get()
}

// GetCategory returns the chromium category
func (Chromium) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the chromium name
func (Chromium) GetName() string {
	return "Chromium"
}

// GetSimbol returns chromium simbol
func (Chromium) GetSimbol() string {
	return "Cr"
}

// GetAtomicNumber returns the chromium atomic number
func (Chromium) GetAtomicNumber() int {
	return 24
}

// GetAtomicWeight returns the chromium atomic weight
func (Chromium) GetAtomicWeight() float32 {
	return 51.996
}
