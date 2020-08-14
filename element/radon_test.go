package element

import "testing"

func TestRadonGetPeriod(t *testing.T) {
	r := Radon{}
	want := "6th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Radon.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRadonGetGroup(t *testing.T) {
	r := Radon{}
	want := "8A"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Radon.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRadonGetCategory(t *testing.T) {
	r := Radon{}
	want := "Noble gas"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Radon.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRadonGetName(t *testing.T) {
	r := Radon{}
	want := "Radon"
	got := r.GetName()
	if got != want {
		t.Errorf("Radon.GetName() = got %v, want %v", got, want)
	}
}

func TestRadonGetSimbol(t *testing.T) {
	r := Radon{}
	want := "Rn"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Radon.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRadonGetAtomicNumber(t *testing.T) {
	r := Radon{}
	want := 86
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Radon.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRadonGetAtomicWeight(t *testing.T) {
	r := Radon{}
	var want float32 = 222
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Radon.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
