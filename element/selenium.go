package element

// Selenium strtuct for selenium element
type Selenium struct{}

// GetPeriod returns the selenium period
func (Selenium) GetPeriod() string {
	var p periodType = fourthPeriod
	return p.get()
}

// GetGroup returns the selenium group
func (Selenium) GetGroup() string {
	var g groupType = a6
	return g.get()
}

// GetCategory returns the selenium category
func (Selenium) GetCategory() string {
	var c categoryType = nonMetal
	return c.get()
}

// GetName returns the selenium name
func (Selenium) GetName() string {
	return "Selenium"
}

// GetSimbol returns selenium simbol
func (Selenium) GetSimbol() string {
	return "Se"
}

// GetAtomicNumber returns the selenium atomic number
func (Selenium) GetAtomicNumber() int {
	return 34
}

// GetAtomicWeight returns the selenium atomic weight
func (Selenium) GetAtomicWeight() float32 {
	return 78.971
}
