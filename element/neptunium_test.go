package element

import "testing"

func TestNeptuniumGetPeriod(t *testing.T) {
	n := Neptunium{}
	want := "7th period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Neptunium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNeptuniumGetGroup(t *testing.T) {
	n := Neptunium{}
	want := "3B"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Neptunium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNeptuniumGetCategory(t *testing.T) {
	n := Neptunium{}
	want := "Actinoid"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Neptunium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNeptuniumGetName(t *testing.T) {
	n := Neptunium{}
	want := "Neptunium"
	got := n.GetName()
	if got != want {
		t.Errorf("Neptunium.GetName() = got %v, want %v", got, want)
	}
}

func TestNeptuniumGetSimbol(t *testing.T) {
	n := Neptunium{}
	want := "Np"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Neptunium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNeptuniumGetAtomicNumber(t *testing.T) {
	n := Neptunium{}
	want := 93
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Neptunium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNeptuniumGetAtomicWeight(t *testing.T) {
	n := Neptunium{}
	var want float32 = 237
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Neptunium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
