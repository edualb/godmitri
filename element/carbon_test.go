package element

import "testing"

func TestCarbonGetPeriod(t *testing.T) {
	c := Carbon{}
	want := "2nd period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Carbon.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCarbonGetGroup(t *testing.T) {
	c := Carbon{}
	want := "4A"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Carbon.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCarbonGetCategory(t *testing.T) {
	c := Carbon{}
	want := "Non-metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Carbon.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCarbonGetName(t *testing.T) {
	c := Carbon{}
	want := "Carbon"
	got := c.GetName()
	if got != want {
		t.Errorf("Carbon.GetName() = got %v, want %v", got, want)
	}
}

func TestCarbonGetSimbol(t *testing.T) {
	c := Carbon{}
	want := "C"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Carbon.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCarbonGetAtomicNumber(t *testing.T) {
	c := Carbon{}
	want := 6
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Carbon.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCarbonGetAtomicWeight(t *testing.T) {
	c := Carbon{}
	var want float32 = 12.011
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Carbon.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
