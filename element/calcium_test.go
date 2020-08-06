package element

import "testing"

func TestCalciumGetPeriod(t *testing.T) {
	c := Calcium{}
	want := "4th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Calcium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCalciumGetGroup(t *testing.T) {
	c := Calcium{}
	want := "2A"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Calcium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCalciumGetCategory(t *testing.T) {
	c := Calcium{}
	want := "Alkaline earth metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Calcium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCalciumGetName(t *testing.T) {
	c := Calcium{}
	want := "Calcium"
	got := c.GetName()
	if got != want {
		t.Errorf("Calcium.GetName() = got %v, want %v", got, want)
	}
}

func TestCalciumGetSimbol(t *testing.T) {
	c := Calcium{}
	want := "Ca"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Calcium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCalciumGetAtomicNumber(t *testing.T) {
	c := Calcium{}
	want := 20
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Calcium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCalciumGetAtomicWeight(t *testing.T) {
	c := Calcium{}
	var want float32 = 40.078
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Calcium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
