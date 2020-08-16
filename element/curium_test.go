package element

import "testing"

func TestCuriumGetPeriod(t *testing.T) {
	c := Curium{}
	want := "7th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Curium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCuriumGetGroup(t *testing.T) {
	c := Curium{}
	want := "3B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Curium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCuriumGetCategory(t *testing.T) {
	c := Curium{}
	want := "Actinoid"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Curium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCuriumGetName(t *testing.T) {
	c := Curium{}
	want := "Curium"
	got := c.GetName()
	if got != want {
		t.Errorf("Curium.GetName() = got %v, want %v", got, want)
	}
}

func TestCuriumGetSimbol(t *testing.T) {
	c := Curium{}
	want := "Cm"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Curium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCuriumGetAtomicNumber(t *testing.T) {
	c := Curium{}
	want := 96
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Curium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCuriumGetAtomicWeight(t *testing.T) {
	c := Curium{}
	var want float32 = 247
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Curium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
