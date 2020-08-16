package element

import "testing"

func TestRadiumGetPeriod(t *testing.T) {
	r := Radium{}
	want := "7th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Radium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRadiumGetGroup(t *testing.T) {
	r := Radium{}
	want := "2A"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Radium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRadiumGetCategory(t *testing.T) {
	r := Radium{}
	want := "Alkaline earth metal"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Radium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRadiumGetName(t *testing.T) {
	r := Radium{}
	want := "Radium"
	got := r.GetName()
	if got != want {
		t.Errorf("Radium.GetName() = got %v, want %v", got, want)
	}
}

func TestRadiumGetSimbol(t *testing.T) {
	r := Radium{}
	want := "Ra"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Radium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRadiumGetAtomicNumber(t *testing.T) {
	r := Radium{}
	want := 88
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Radium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRadiumGetAtomicWeight(t *testing.T) {
	r := Radium{}
	var want float32 = 226
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Radium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
