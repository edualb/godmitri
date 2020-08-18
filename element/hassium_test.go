package element

import "testing"

func TestHassiumGetPeriod(t *testing.T) {
	h := Hassium{}
	want := "7th period"
	got := h.GetPeriod()
	if got != want {
		t.Errorf("Hassium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestHassiumGetGroup(t *testing.T) {
	h := Hassium{}
	want := "8B"
	got := h.GetGroup()
	if got != want {
		t.Errorf("Hassium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestHassiumGetCategory(t *testing.T) {
	h := Hassium{}
	want := "Transition metal"
	got := h.GetCategory()
	if got != want {
		t.Errorf("Hassium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestHassiumfGetName(t *testing.T) {
	h := Hassium{}
	want := "Hassium"
	got := h.GetName()
	if got != want {
		t.Errorf("Hassium.GetName() = got %v, want %v", got, want)
	}
}

func TestHassiumGetSimbol(t *testing.T) {
	h := Hassium{}
	want := "Hs"
	got := h.GetSimbol()
	if got != want {
		t.Errorf("Hassium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestHassiumGetAtomicNumber(t *testing.T) {
	h := Hassium{}
	want := 108
	got := h.GetAtomicNumber()
	if got != want {
		t.Errorf("Hassium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestHassiumGetAtomicWeight(t *testing.T) {
	h := Hassium{}
	var want float32 = 277
	got := h.GetAtomicWeight()
	if got != want {
		t.Errorf("Hassium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
