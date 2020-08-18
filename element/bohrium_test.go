package element

import "testing"

func TestBohriumGetPeriod(t *testing.T) {
	b := Bohrium{}
	want := "7th period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Bohrium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBohriumGetGroup(t *testing.T) {
	b := Bohrium{}
	want := "7B"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Bohrium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBohriumGetCategory(t *testing.T) {
	b := Bohrium{}
	want := "Transition metal"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Bohrium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBohriumGetName(t *testing.T) {
	b := Bohrium{}
	want := "Bohrium"
	got := b.GetName()
	if got != want {
		t.Errorf("Bohrium.GetName() = got %v, want %v", got, want)
	}
}

func TestBohriumGetSimbol(t *testing.T) {
	b := Bohrium{}
	want := "Bh"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Bohrium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBohriumGetAtomicNumber(t *testing.T) {
	b := Bohrium{}
	want := 107
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Bohrium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBohriumGetAtomicWeight(t *testing.T) {
	b := Bohrium{}
	var want float32 = 270
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Bohrium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
