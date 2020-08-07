package element

import "testing"

func TestCaesiumGetPeriod(t *testing.T) {
	c := Caesium{}
	want := "6th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Caesium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCaesiumGetGroup(t *testing.T) {
	c := Caesium{}
	want := "1A"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Caesium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCaesiumGetCategory(t *testing.T) {
	c := Caesium{}
	want := "Alkali metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Caesium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCaesiumGetName(t *testing.T) {
	c := Caesium{}
	want := "Caesium"
	got := c.GetName()
	if got != want {
		t.Errorf("Caesium.GetName() = got %v, want %v", got, want)
	}
}

func TestCaesiumGetSimbol(t *testing.T) {
	c := Caesium{}
	want := "Cs"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Caesium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCaesiumGetAtomicNumber(t *testing.T) {
	c := Caesium{}
	want := 55
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Caesium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCaesiumGetAtomicWeight(t *testing.T) {
	c := Caesium{}
	var want float32 = 132.91
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Caesium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
