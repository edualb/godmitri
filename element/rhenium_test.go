package element

import "testing"

func TestRheniumGetPeriod(t *testing.T) {
	r := Rhenium{}
	want := "6th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Rhenium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRheniumGetGroup(t *testing.T) {
	r := Rhenium{}
	want := "7B"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Rhenium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRheniumGetCategory(t *testing.T) {
	r := Rhenium{}
	want := "Transition metal"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Rhenium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRheniumGetName(t *testing.T) {
	r := Rhenium{}
	want := "Rhenium"
	got := r.GetName()
	if got != want {
		t.Errorf("Rhenium.GetName() = got %v, want %v", got, want)
	}
}

func TestRheniumGetSimbol(t *testing.T) {
	r := Rhenium{}
	want := "Re"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Rhenium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRheniumGetAtomicNumber(t *testing.T) {
	r := Rhenium{}
	want := 75
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Rhenium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRheniumGetAtomicWeight(t *testing.T) {
	r := Rhenium{}
	var want float32 = 186.21
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Rhenium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
