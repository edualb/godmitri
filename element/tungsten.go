package element

// Tungsten strtuct for tungsten element
type Tungsten struct{}

// GetPeriod returns the tungsten period
func (Tungsten) GetPeriod() string {
	var p periodType = sixthPeriod
	return p.get()
}

// GetGroup returns the tungsten group
func (Tungsten) GetGroup() string {
	var g groupType = b6
	return g.get()
}

// GetCategory returns the tungsten category
func (Tungsten) GetCategory() string {
	var c categoryType = transitionMetal
	return c.get()
}

// GetName returns the tungsten name
func (Tungsten) GetName() string {
	return "Tungsten"
}

// GetSimbol returns tungsten simbol
func (Tungsten) GetSimbol() string {
	return "W"
}

// GetAtomicNumber returns the tungsten atomic number
func (Tungsten) GetAtomicNumber() int {
	return 74
}

// GetAtomicWeight returns the tungsten atomic weight
func (Tungsten) GetAtomicWeight() float32 {
	return 183.84
}
