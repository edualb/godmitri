package element

import "testing"

func TestRoentgeniumGetPeriod(t *testing.T) {
	r := Roentgenium{}
	want := "7th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Roentgenium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRoentgeniumGetGroup(t *testing.T) {
	r := Roentgenium{}
	want := "1B"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Roentgenium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRoentgeniumGetCategory(t *testing.T) {
	r := Roentgenium{}
	want := "Unknown"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Roentgenium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRoentgeniumGetName(t *testing.T) {
	r := Roentgenium{}
	want := "Roentgenium"
	got := r.GetName()
	if got != want {
		t.Errorf("Roentgenium.GetName() = got %v, want %v", got, want)
	}
}

func TestRoentgeniumGetSimbol(t *testing.T) {
	r := Roentgenium{}
	want := "Rg"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Roentgenium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRoentgeniumGetAtomicNumber(t *testing.T) {
	r := Roentgenium{}
	want := 111
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Roentgenium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRoentgeniumGetAtomicWeight(t *testing.T) {
	r := Roentgenium{}
	var want float32 = 282
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Roentgenium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
