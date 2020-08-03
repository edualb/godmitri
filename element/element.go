package element

// Element interface for periodic table elements
type Element interface {
	GetPeriod() string
	GetGroup() string
	GetCategory() string
	GetName() string
	GetSimbol() string
	GetAtomicNumber() int
	GetAtomicWeight() float32
}
