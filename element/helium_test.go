package element

import "testing"

func TestHeliumGetPeriod(t *testing.T) {
	h := Helium{}
	want := "1st period"
	got := h.GetPeriod()
	if got != want {
		t.Errorf("Helium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestHeliumGetGroup(t *testing.T) {
	h := Helium{}
	want := "8A"
	got := h.GetGroup()
	if got != want {
		t.Errorf("Helium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestHeliumGetCategory(t *testing.T) {
	h := Helium{}
	want := "noble gas"
	got := h.GetCategory()
	if got != want {
		t.Errorf("Helium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestHeliumGetName(t *testing.T) {
	h := Helium{}
	want := "Helium"
	got := h.GetName()
	if got != want {
		t.Errorf("Helium.GetName() = got %v, want %v", got, want)
	}
}

func TestHeliumGetSimbol(t *testing.T) {
	h := Helium{}
	want := "He"
	got := h.GetSimbol()
	if got != want {
		t.Errorf("Helium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestHeliumGetAtomicNumber(t *testing.T) {
	h := Helium{}
	want := 2
	got := h.GetAtomicNumber()
	if got != want {
		t.Errorf("Helium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestHeliumGetAtomicWeight(t *testing.T) {
	h := Helium{}
	var want float32 = 4.0026
	got := h.GetAtomicWeight()
	if got != want {
		t.Errorf("Helium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
