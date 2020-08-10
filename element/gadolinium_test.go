package element

import "testing"

func TestGadoliniumGetPeriod(t *testing.T) {
	g := Gadolinium{}
	want := "6th period"
	got := g.GetPeriod()
	if got != want {
		t.Errorf("Gadolinium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestGadoliniumGetGroup(t *testing.T) {
	g := Gadolinium{}
	want := "3B"
	got := g.GetGroup()
	if got != want {
		t.Errorf("Gadolinium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestGadoliniumGetCategory(t *testing.T) {
	g := Gadolinium{}
	want := "Lanthanoid"
	got := g.GetCategory()
	if got != want {
		t.Errorf("Gadolinium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestGadoliniumGetName(t *testing.T) {
	g := Gadolinium{}
	want := "Gadolinium"
	got := g.GetName()
	if got != want {
		t.Errorf("Gadolinium.GetName() = got %v, want %v", got, want)
	}
}

func TestGadoliniumGetSimbol(t *testing.T) {
	g := Gadolinium{}
	want := "Gd"
	got := g.GetSimbol()
	if got != want {
		t.Errorf("Gadolinium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestGadoliniumGetAtomicNumber(t *testing.T) {
	g := Gadolinium{}
	want := 64
	got := g.GetAtomicNumber()
	if got != want {
		t.Errorf("Gadolinium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestGadoliniumGetAtomicWeight(t *testing.T) {
	g := Gadolinium{}
	var want float32 = 157.25
	got := g.GetAtomicWeight()
	if got != want {
		t.Errorf("Gadolinium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
