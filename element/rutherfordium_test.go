package element

import "testing"

func TestRutherfordiumGetPeriod(t *testing.T) {
	r := Rutherfordium{}
	want := "7th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Rutherfordium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRutherfordiumGetGroup(t *testing.T) {
	r := Rutherfordium{}
	want := "4B"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Rutherfordium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRutherfordiumGetCategory(t *testing.T) {
	r := Rutherfordium{}
	want := "Transition metal"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Rutherfordium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRutherfordiumGetName(t *testing.T) {
	r := Rutherfordium{}
	want := "Rutherfordium"
	got := r.GetName()
	if got != want {
		t.Errorf("Rutherfordium.GetName() = got %v, want %v", got, want)
	}
}

func TestRutherfordiumGetSimbol(t *testing.T) {
	r := Rutherfordium{}
	want := "Rf"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Rutherfordium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRutherfordiumGetAtomicNumber(t *testing.T) {
	r := Rutherfordium{}
	want := 104
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Rutherfordium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRutherfordiumGetAtomicWeight(t *testing.T) {
	r := Rutherfordium{}
	var want float32 = 267
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Rutherfordium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
