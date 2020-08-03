package element

import "testing"

func TestHydrogenGetPeriod(t *testing.T) {
	h := Hydrogen{}
	want := "1st period"
	got := h.GetPeriod()
	if got != want {
		t.Errorf("Hydrogen.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestHydrogenGetGroup(t *testing.T) {
	h := Hydrogen{}
	want := ""
	got := h.GetGroup()
	if got != want {
		t.Errorf("Hydrogen.GetGroup() = got %v, want %v", got, want)
	}
}

func TestHydrogenGetCategory(t *testing.T) {
	h := Hydrogen{}
	want := "non-metal"
	got := h.GetCategory()
	if got != want {
		t.Errorf("Hydrogen.GetCategory() = got %v, want %v", got, want)
	}
}

func TestHydrogenGetName(t *testing.T) {
	h := Hydrogen{}
	want := "Hydrogen"
	got := h.GetName()
	if got != want {
		t.Errorf("Hydrogen.GetName() = got %v, want %v", got, want)
	}
}

func TestHydrogenGetSimbol(t *testing.T) {
	h := Hydrogen{}
	want := "H"
	got := h.GetSimbol()
	if got != want {
		t.Errorf("Hydrogen.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestHydrogenGetAtomicNumber(t *testing.T) {
	h := Hydrogen{}
	want := 1
	got := h.GetAtomicNumber()
	if got != want {
		t.Errorf("Hydrogen.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestHydrogenGetAtomicWeight(t *testing.T) {
	h := Hydrogen{}
	var want float32 = 1.008
	got := h.GetAtomicWeight()
	if got != want {
		t.Errorf("Hydrogen.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
