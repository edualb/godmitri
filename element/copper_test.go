package element

import "testing"

func TestCopperGetPeriod(t *testing.T) {
	c := Copper{}
	want := "4th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Copper.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCopperGetGroup(t *testing.T) {
	c := Copper{}
	want := "1B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Copper.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCopperGetCategory(t *testing.T) {
	c := Copper{}
	want := "Transition metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Copper.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCopperGetName(t *testing.T) {
	c := Copper{}
	want := "Copper"
	got := c.GetName()
	if got != want {
		t.Errorf("Copper.GetName() = got %v, want %v", got, want)
	}
}

func TestCopperGetSimbol(t *testing.T) {
	c := Copper{}
	want := "Cu"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Copper.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCopperGetAtomicNumber(t *testing.T) {
	c := Copper{}
	want := 29
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Copper.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCopperGetAtomicWeight(t *testing.T) {
	c := Copper{}
	var want float32 = 63.546
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Copper.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
