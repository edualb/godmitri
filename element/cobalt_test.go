package element

import "testing"

func TestCobaltGetPeriod(t *testing.T) {
	c := Cobalt{}
	want := "4th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Cobalt.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCobaltGetGroup(t *testing.T) {
	c := Cobalt{}
	want := "8B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Cobalt.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCobaltGetCategory(t *testing.T) {
	c := Cobalt{}
	want := "Transition metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Cobalt.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCobaltGetName(t *testing.T) {
	c := Cobalt{}
	want := "Cobalt"
	got := c.GetName()
	if got != want {
		t.Errorf("Cobalt.GetName() = got %v, want %v", got, want)
	}
}

func TestCobaltGetSimbol(t *testing.T) {
	c := Cobalt{}
	want := "Co"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Cobalt.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCobaltGetAtomicNumber(t *testing.T) {
	c := Cobalt{}
	want := 27
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Cobalt.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCobaltGetAtomicWeight(t *testing.T) {
	c := Cobalt{}
	var want float32 = 58.933
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Cobalt.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
