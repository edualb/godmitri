package element

import "testing"

func TestNickelGetPeriod(t *testing.T) {
	n := Nickel{}
	want := "4th period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Nickel.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNickelGetGroup(t *testing.T) {
	n := Nickel{}
	want := "8B"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Nickel.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNickelGetCategory(t *testing.T) {
	n := Nickel{}
	want := "Transition metal"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Nickel.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNickelGetName(t *testing.T) {
	n := Nickel{}
	want := "Nickel"
	got := n.GetName()
	if got != want {
		t.Errorf("Nickel.GetName() = got %v, want %v", got, want)
	}
}

func TestNickelGetSimbol(t *testing.T) {
	n := Nickel{}
	want := "Ni"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Nickel.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNickelGetAtomicNumber(t *testing.T) {
	n := Nickel{}
	want := 28
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Nickel.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNickelGetAtomicWeight(t *testing.T) {
	n := Nickel{}
	var want float32 = 58.693
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Nickel.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
