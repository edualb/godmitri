package element

import "testing"

func TestCeriumGetPeriod(t *testing.T) {
	c := Cerium{}
	want := "6th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Cerium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCeriumGetGroup(t *testing.T) {
	c := Cerium{}
	want := "3B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Cerium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCeriumGetCategory(t *testing.T) {
	c := Cerium{}
	want := "Lanthanoid"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Cerium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCeriumGetName(t *testing.T) {
	c := Cerium{}
	want := "Cerium"
	got := c.GetName()
	if got != want {
		t.Errorf("Cerium.GetName() = got %v, want %v", got, want)
	}
}

func TestCeriumGetSimbol(t *testing.T) {
	c := Cerium{}
	want := "Ce"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Cerium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCeriumGetAtomicNumber(t *testing.T) {
	c := Cerium{}
	want := 58
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Cerium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCeriumGetAtomicWeight(t *testing.T) {
	c := Cerium{}
	var want float32 = 140.12
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Cerium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
