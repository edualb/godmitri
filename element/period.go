package element

type periodType int

const (
	firstPeriod = iota
	secondPeriod
	thirdPeriod
	fourthPeriod
	fifthPeriod
	sixthPeriod
	seventhPeriod
)

// Get gets the periodType
func (pt periodType) get() string {
	return [...]string{
		"1st period",
		"2nd period",
		"3rd period",
		"4th period",
		"5th period",
		"6th period",
		"7th period",
	}[pt]
}
