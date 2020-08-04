package element

type groupType int

const (
	a1 = iota
	a2
	a3
	a4
	a5
	a6
	a7
	a8
	b1
	b2
	b3
	b4
	b5
	b6
	b7
	b8
)

// Get gets the groupType
func (gt groupType) get() string {
	return [...]string{
		"1A",
		"2A",
		"3A",
		"4A",
		"5A",
		"6A",
		"7A",
		"8A",
		"1B",
		"2B",
		"3B",
		"4B",
		"5B",
		"6B",
		"7B",
		"8B",
	}[gt]
}
