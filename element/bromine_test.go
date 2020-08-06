package element

import "testing"

func TestBromineGetPeriod(t *testing.T) {
	b := Bromine{}
	want := "4th period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Boron.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBromineGetGroup(t *testing.T) {
	b := Bromine{}
	want := "7A"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Bromine.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBromineGetCategory(t *testing.T) {
	b := Bromine{}
	want := "Non-metal"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Bromine.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBromineGetName(t *testing.T) {
	b := Bromine{}
	want := "Bromine"
	got := b.GetName()
	if got != want {
		t.Errorf("Bromine.GetName() = got %v, want %v", got, want)
	}
}

func TestBromineGetSimbol(t *testing.T) {
	b := Bromine{}
	want := "Br"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Bromine.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBromineGetAtomicNumber(t *testing.T) {
	b := Bromine{}
	want := 35
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Bromine.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBromineGetAtomicWeight(t *testing.T) {
	b := Bromine{}
	var want float32 = 79.904
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Bromine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
