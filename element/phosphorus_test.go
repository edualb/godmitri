package element

import "testing"

func TestPhosphorusGetPeriod(t *testing.T) {
	p := Phosphorus{}
	want := "3rd period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Phosphorus.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPhosphorusGetGroup(t *testing.T) {
	p := Phosphorus{}
	want := "5A"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Phosphorus.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPhosphorusGetCategory(t *testing.T) {
	p := Phosphorus{}
	want := "Non-metal"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Phosphorus.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPhosphorusGetName(t *testing.T) {
	p := Phosphorus{}
	want := "Phosphorus"
	got := p.GetName()
	if got != want {
		t.Errorf("Phosphorus.GetName() = got %v, want %v", got, want)
	}
}

func TestPhosphorusGetSimbol(t *testing.T) {
	p := Phosphorus{}
	want := "P"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Phosphorus.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPhosphorusGetAtomicNumber(t *testing.T) {
	p := Phosphorus{}
	want := 15
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Phosphorus.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPhosphorusGetAtomicWeight(t *testing.T) {
	p := Phosphorus{}
	var want float32 = 30.974
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Phosphorus.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
