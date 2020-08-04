package element

import "testing"

func TestNeonGetPeriod(t *testing.T) {
	n := Neon{}
	want := "2nd period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Neon.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNeonGetGroup(t *testing.T) {
	n := Neon{}
	want := "8A"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Neon.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNeonGetCategory(t *testing.T) {
	n := Neon{}
	want := "Noble gas"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Neon.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNeonGetName(t *testing.T) {
	n := Neon{}
	want := "Neon"
	got := n.GetName()
	if got != want {
		t.Errorf("Neon.GetName() = got %v, want %v", got, want)
	}
}

func TestNeonGetSimbol(t *testing.T) {
	n := Neon{}
	want := "Ne"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Neon.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNeonGetAtomicNumber(t *testing.T) {
	n := Neon{}
	want := 10
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Neon.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNeonGetAtomicWeight(t *testing.T) {
	n := Neon{}
	var want float32 = 20.180
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Neon.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
