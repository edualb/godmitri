package element

import "testing"

func TestCaliforniumGetPeriod(t *testing.T) {
	c := Californium{}
	want := "7th period"
	got := c.GetPeriod()
	if got != want {
		t.Errorf("Californium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestCaliforniumGetGroup(t *testing.T) {
	c := Californium{}
	want := "3B"
	got := c.GetGroup()
	if got != want {
		t.Errorf("Californium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestCaliforniumGetCategory(t *testing.T) {
	c := Californium{}
	want := "Actinoid"
	got := c.GetCategory()
	if got != want {
		t.Errorf("Californium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestCaliforniumGetName(t *testing.T) {
	c := Californium{}
	want := "Californium"
	got := c.GetName()
	if got != want {
		t.Errorf("Californium.GetName() = got %v, want %v", got, want)
	}
}

func TestCaliforniumGetSimbol(t *testing.T) {
	c := Californium{}
	want := "Cf"
	got := c.GetSimbol()
	if got != want {
		t.Errorf("Californium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestCaliforniumGetAtomicNumber(t *testing.T) {
	c := Californium{}
	want := 98
	got := c.GetAtomicNumber()
	if got != want {
		t.Errorf("Californium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestCaliforniumGetAtomicWeight(t *testing.T) {
	c := Californium{}
	var want float32 = 251
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Californium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
