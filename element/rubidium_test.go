package element

import "testing"

func TestRubidiumGetPeriod(t *testing.T) {
	r := Rubidium{}
	want := "5th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Rubidium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRubidiumGetGroup(t *testing.T) {
	r := Rubidium{}
	want := "1A"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Rubidium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRubidiumGetCategory(t *testing.T) {
	r := Rubidium{}
	want := "Alkali metal"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Rubidium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRubidiumGetName(t *testing.T) {
	r := Rubidium{}
	want := "Rubidium"
	got := r.GetName()
	if got != want {
		t.Errorf("Rubidium.GetName() = got %v, want %v", got, want)
	}
}

func TestRubidiumGetSimbol(t *testing.T) {
	r := Rubidium{}
	want := "Rb"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Rubidium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRubidiumGetAtomicNumber(t *testing.T) {
	r := Rubidium{}
	want := 37
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Rubidium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRubidiumGetAtomicWeight(t *testing.T) {
	r := Rubidium{}
	var want float32 = 85.468
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Rubidium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
