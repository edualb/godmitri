package element

import "testing"

func TestRutheniumGetPeriod(t *testing.T) {
	r := Ruthenium{}
	want := "5th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Ruthenium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRutheniumGetGroup(t *testing.T) {
	r := Ruthenium{}
	want := "8B"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Ruthenium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRutheniumGetCategory(t *testing.T) {
	r := Ruthenium{}
	want := "Transition metal"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Ruthenium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRutheniumGetName(t *testing.T) {
	r := Ruthenium{}
	want := "Ruthenium"
	got := r.GetName()
	if got != want {
		t.Errorf("Ruthenium.GetName() = got %v, want %v", got, want)
	}
}

func TestRutheniumGetSimbol(t *testing.T) {
	r := Ruthenium{}
	want := "Ru"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Ruthenium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRutheniumGetAtomicNumber(t *testing.T) {
	r := Ruthenium{}
	want := 44
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Ruthenium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRutheniumGetAtomicWeight(t *testing.T) {
	r := Ruthenium{}
	var want float32 = 101.07
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Ruthenium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
