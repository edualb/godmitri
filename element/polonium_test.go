package element

import "testing"

func TestPoloniumGetPeriod(t *testing.T) {
	p := Polonium{}
	want := "6th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Polonium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPoloniumGetGroup(t *testing.T) {
	p := Polonium{}
	want := "6A"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Polonium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPoloniumGetCategory(t *testing.T) {
	p := Polonium{}
	want := "Post-transition metal"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Polonium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPoloniumGetName(t *testing.T) {
	p := Polonium{}
	want := "Polonium"
	got := p.GetName()
	if got != want {
		t.Errorf("Polonium.GetName() = got %v, want %v", got, want)
	}
}

func TestPoloniumGetSimbol(t *testing.T) {
	p := Polonium{}
	want := "Po"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Polonium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPoloniumGetAtomicNumber(t *testing.T) {
	p := Polonium{}
	want := 84
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Polonium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPoloniumGetAtomicWeight(t *testing.T) {
	p := Polonium{}
	var want float32 = 209
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Polonium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
