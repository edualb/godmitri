package element

import "testing"

func TestNeodymiumGetPeriod(t *testing.T) {
	n := Neodymium{}
	want := "6th period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Neodymium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNeodymiumGetGroup(t *testing.T) {
	n := Neodymium{}
	want := "3B"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Neodymium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNeodymiumGetCategory(t *testing.T) {
	n := Neodymium{}
	want := "Lanthanoid"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Neodymium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNeodymiumGetName(t *testing.T) {
	n := Neodymium{}
	want := "Neodymium"
	got := n.GetName()
	if got != want {
		t.Errorf("Neodymium.GetName() = got %v, want %v", got, want)
	}
}

func TestNeodymiumGetSimbol(t *testing.T) {
	n := Neodymium{}
	want := "Nd"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Neodymium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNeodymiumGetAtomicNumber(t *testing.T) {
	n := Neodymium{}
	want := 60
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Neodymium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNeodymiumGetAtomicWeight(t *testing.T) {
	n := Neodymium{}
	var want float32 = 144.24
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Neodymium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
