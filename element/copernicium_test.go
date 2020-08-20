package element

import "testing"

func TestCoperniciumGetPeriod(t *testing.T) {
	c := Copernicium{}
	want := "7th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Copernicium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCoperniciumGetGroup(t *testing.T) {
	c := Copernicium{}
	want := "2B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Copernicium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCoperniciumGetCategory(t *testing.T) {
	c := Copernicium{}
	want := "Transition metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Copernicium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCoperniciumGetName(t *testing.T) {
	c := Copernicium{}
	want := "Copernicium"
	got := c.GetName()
	if got != want {
		t.Errorf("Copernicium.GetName() = got %v, want %v", got, want)
	}
}

func TestCoperniciumGetSimbol(t *testing.T) {
	c := Copernicium{}
	want := "Cn"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Copernicium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCoperniciumGetAtomicNumber(t *testing.T) {
	c := Copernicium{}
	want := 112
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Copernicium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCoperniciumGetAtomicWeight(t *testing.T) {
	c := Copernicium{}
	var want float32 = 285
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Copernicium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
