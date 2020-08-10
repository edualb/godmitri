package element

import "testing"

func TestHolmiumGetPeriod(t *testing.T) {
	h := Holmium{}
	want := "6th period"
	got := h.GetPeriod()
	if got != want {
		t.Errorf("Holmium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestHolmiumGetGroup(t *testing.T) {
	h := Holmium{}
	want := "3B"
	got := h.GetGroup()
	if got != want {
		t.Errorf("Holmium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestHolmiumGetCategory(t *testing.T) {
	h := Holmium{}
	want := "Lanthanoid"
	got := h.GetCategory()
	if got != want {
		t.Errorf("Holmium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestHolmiumGetName(t *testing.T) {
	h := Holmium{}
	want := "Holmium"
	got := h.GetName()
	if got != want {
		t.Errorf("Holmium.GetName() = got %v, want %v", got, want)
	}
}

func TestHolmiumGetSimbol(t *testing.T) {
	h := Holmium{}
	want := "Ho"
	got := h.GetSimbol()
	if got != want {
		t.Errorf("Holmium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestHolmiumGetAtomicNumber(t *testing.T) {
	h := Holmium{}
	want := 67
	got := h.GetAtomicNumber()
	if got != want {
		t.Errorf("Holmium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestHolmiumGetAtomicWeight(t *testing.T) {
	h := Holmium{}
	var want float32 = 164.93
	got := h.GetAtomicWeight()
	if got != want {
		t.Errorf("Holmium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
