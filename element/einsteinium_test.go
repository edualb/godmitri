package element

import "testing"

func TestEinsteiniumGetPeriod(t *testing.T) {
	c := Einsteinium{}
	want := "7th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Einsteinium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestEinsteiniumGetGroup(t *testing.T) {
	c := Einsteinium{}
	want := "3B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Einsteinium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestEinsteiniumGetCategory(t *testing.T) {
	c := Einsteinium{}
	want := "Actinoid"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Einsteinium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestEinsteiniumGetName(t *testing.T) {
	c := Einsteinium{}
	want := "Einsteinium"
	got := c.GetName()
	if got != want {
		t.Errorf("Einsteinium.GetName() = got %v, want %v", got, want)
	}
}

func TestEinsteiniumGetSimbol(t *testing.T) {
	c := Einsteinium{}
	want := "Es"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Einsteinium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestEinsteiniumGetAtomicNumber(t *testing.T) {
	c := Einsteinium{}
	want := 99
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Einsteinium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestEinsteiniumGetAtomicWeight(t *testing.T) {
	c := Einsteinium{}
	var want float32 = 252
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Einsteinium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
