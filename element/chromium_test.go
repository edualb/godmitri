package element

import "testing"

func TestChromiumGetPeriod(t *testing.T) {
	c := Chromium{}
	want := "4th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Chromium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestChromiumGetGroup(t *testing.T) {
	c := Chromium{}
	want := "6B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Chromium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestChromiumGetCategory(t *testing.T) {
	c := Chromium{}
	want := "Transition metal"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Chromium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestChromiumGetName(t *testing.T) {
	c := Chromium{}
	want := "Chromium"
	got := c.GetName()
	if got != want {
		t.Errorf("Chromium.GetName() = got %v, want %v", got, want)
	}
}

func TestChromiumGetSimbol(t *testing.T) {
	c := Chromium{}
	want := "Cr"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Chromium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestChromiumGetAtomicNumber(t *testing.T) {
	c := Chromium{}
	want := 24
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Chromium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestChromiumGetAtomicWeight(t *testing.T) {
	c := Chromium{}
	var want float32 = 51.996
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Chromium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
