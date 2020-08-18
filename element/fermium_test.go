package element

import "testing"

func TestFermiumGetPeriod(t *testing.T) {
	f := Fermium{}
	want := "7th period"
	got := f.GetPeriod()
	if got != want {
		t.Errorf("Fermium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestFermiumGetGroup(t *testing.T) {
	f := Fermium{}
	want := "3B"
	got := f.GetGroup()
	if got != want {
		t.Errorf("Fermium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestFermiumGetCategory(t *testing.T) {
	f := Fermium{}
	want := "Actinoid"
	got := f.GetCategory()
	if got != want {
		t.Errorf("Fermium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestFermiumGetName(t *testing.T) {
	f := Fermium{}
	want := "Fermium"
	got := f.GetName()
	if got != want {
		t.Errorf("Fermium.GetName() = got %v, want %v", got, want)
	}
}

func TestFermiumGetSimbol(t *testing.T) {
	f := Fermium{}
	want := "Fm"
	got := f.GetSimbol()
	if got != want {
		t.Errorf("Fermium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestFermiumGetAtomicNumber(t *testing.T) {
	f := Fermium{}
	want := 100
	got := f.GetAtomicNumber()
	if got != want {
		t.Errorf("Fermium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestFermiumGetAtomicWeight(t *testing.T) {
	f := Fermium{}
	var want float32 = 257
	got := f.GetAtomicWeight()
	if got != want {
		t.Errorf("Fermium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
