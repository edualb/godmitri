package element

import "testing"

func TestCadmiumGetPeriod(t *testing.T) {
	c := Cadmium{}
	want := "5th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Cadmium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCadmiumGetGroup(t *testing.T) {
	c := Cadmium{}
	want := "2B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Cadmium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCadmiumGetCategory(t *testing.T) {
	c := Cadmium{}
	want := "Transition metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Cadmium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCadmiumGetName(t *testing.T) {
	c := Cadmium{}
	want := "Cadmium"
	got := c.GetName()
	if got != want {
		t.Errorf("Cadmium.GetName() = got %v, want %v", got, want)
	}
}

func TestCadmiumGetSimbol(t *testing.T) {
	c := Cadmium{}
	want := "Cd"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Cadmium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCadmiumGetAtomicNumber(t *testing.T) {
	c := Cadmium{}
	want := 48
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Cadmium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCadmiumGetAtomicWeight(t *testing.T) {
	c := Cadmium{}
	var want float32 = 112.41
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Calcium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
