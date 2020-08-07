package element

import "testing"

func TestNiobiumGetPeriod(t *testing.T) {
	n := Niobium{}
	want := "5th period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Niobium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNiobiumGetGroup(t *testing.T) {
	n := Niobium{}
	want := "5B"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Niobium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNiobiumGetCategory(t *testing.T) {
	n := Niobium{}
	want := "Transition metal"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Niobium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNiobiumGetName(t *testing.T) {
	n := Niobium{}
	want := "Niobium"
	got := n.GetName()
	if got != want {
		t.Errorf("Niobium.GetName() = got %v, want %v", got, want)
	}
}

func TestNiobiumGetSimbol(t *testing.T) {
	n := Niobium{}
	want := "Nb"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Niobium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNiobiumGetAtomicNumber(t *testing.T) {
	n := Niobium{}
	want := 41
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Niobium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNiobiumGetAtomicWeight(t *testing.T) {
	n := Niobium{}
	var want float32 = 92.906
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Niobium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
