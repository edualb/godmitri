package element

import "testing"

func TestEuropiumGetPeriod(t *testing.T) {
	c := Europium{}
	want := "6th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Europium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestEuropiumGetGroup(t *testing.T) {
	c := Europium{}
	want := "3B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Europium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestEuropiumGetCategory(t *testing.T) {
	c := Europium{}
	want := "Lanthanoid"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Europium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestEuropiumGetName(t *testing.T) {
	c := Europium{}
	want := "Europium"
	got := c.GetName()
	if got != want {
		t.Errorf("Europium.GetName() = got %v, want %v", got, want)
	}
}

func TestEuropiumGetSimbol(t *testing.T) {
	c := Europium{}
	want := "Eu"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Europium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestEuropiumGetAtomicNumber(t *testing.T) {
	c := Europium{}
	want := 63
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Europium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestEuropiumGetAtomicWeight(t *testing.T) {
	c := Europium{}
	var want float32 = 151.96
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Europium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
