package element

import "testing"

func TestChlorineGetPeriod(t *testing.T) {
	c := Chlorine{}
	want := "3rd period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Chlorine.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestChlorineGetGroup(t *testing.T) {
	c := Chlorine{}
	want := "7A"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Chlorine.GetGroup() = got %v, want %v", got, want)
	}
}

func TestChlorineGetCategory(t *testing.T) {
	c := Chlorine{}
	want := "Non-metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Chlorine.GetCategory() = got %v, want %v", got, want)
	}
}

func TestChlorineGetName(t *testing.T) {
	c := Chlorine{}
	want := "Chlorine"
	got := c.GetName()
	if got != want {
		t.Errorf("Chlorine.GetName() = got %v, want %v", got, want)
	}
}

func TestChlorineGetSimbol(t *testing.T) {
	c := Chlorine{}
	want := "Cl"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Chlorine.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestChlorineGetAtomicNumber(t *testing.T) {
	c := Chlorine{}
	want := 17
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Chlorine.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestChlorineGetAtomicWeight(t *testing.T) {
	c := Chlorine{}
	var want float32 = 35.45
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Chlorine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
