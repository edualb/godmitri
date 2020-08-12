package element

import "testing"

func TestHafniumGetPeriod(t *testing.T) {
	h := Hafnium{}
	want := "6th period"
	got := h.GetPeriod()
	if got != want {
		t.Errorf("Helium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestHafniumGetGroup(t *testing.T) {
	h := Hafnium{}
	want := "4B"
	got := h.GetGroup()
	if got != want {
		t.Errorf("Hafnium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestHafniumGetCategory(t *testing.T) {
	h := Hafnium{}
	want := "Transition metal"
	got := h.GetCategory()
	if got != want {
		t.Errorf("Hafnium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestHafniumGetName(t *testing.T) {
	h := Hafnium{}
	want := "Hafnium"
	got := h.GetName()
	if got != want {
		t.Errorf("Hafnium.GetName() = got %v, want %v", got, want)
	}
}

func TestHafniumGetSimbol(t *testing.T) {
	h := Hafnium{}
	want := "Hf"
	got := h.GetSimbol()
	if got != want {
		t.Errorf("Hafnium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestHafniumGetAtomicNumber(t *testing.T) {
	h := Hafnium{}
	want := 72
	got := h.GetAtomicNumber()
	if got != want {
		t.Errorf("Hafnium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestHafniumGetAtomicWeight(t *testing.T) {
	h := Hafnium{}
	var want float32 = 178.49
	got := h.GetAtomicWeight()
	if got != want {
		t.Errorf("Helium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
