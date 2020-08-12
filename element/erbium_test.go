package element

import "testing"

func TestErbiumGetPeriod(t *testing.T) {
	c := Erbium{}
	want := "6th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Erbium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestErbiumGetGroup(t *testing.T) {
	c := Erbium{}
	want := "3B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Erbium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestErbiumGetCategory(t *testing.T) {
	c := Erbium{}
	want := "Lanthanoid"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Erbium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestErbiumGetName(t *testing.T) {
	c := Erbium{}
	want := "Erbium"
	got := c.GetName()
	if got != want {
		t.Errorf("Erbium.GetName() = got %v, want %v", got, want)
	}
}

func TestErbiumGetSimbol(t *testing.T) {
	c := Erbium{}
	want := "Er"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Erbium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestErbiumGetAtomicNumber(t *testing.T) {
	c := Erbium{}
	want := 68
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Erbium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestErbiumGetAtomicWeight(t *testing.T) {
	c := Erbium{}
	var want float32 = 167.26
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Erbium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
