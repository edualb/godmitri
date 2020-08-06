package element

import "testing"

func TestGermaniumGetPeriod(t *testing.T) {
	g := Germanium{}
	want := "4th period"
	got := g.GetPeriod()
	if got != want {
		t.Errorf("Germanium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestGermaniumGetGroup(t *testing.T) {
	g := Germanium{}
	want := "4A"
	got := g.GetGroup()
	if got != want {
		t.Errorf("Germanium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestGermaniumGetCategory(t *testing.T) {
	g := Germanium{}
	want := "Metalloid"
	got := g.GetCategory()
	if got != want {
		t.Errorf("Germanium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestGermaniumGetName(t *testing.T) {
	g := Germanium{}
	want := "Germanium"
	got := g.GetName()
	if got != want {
		t.Errorf("Germanium.GetName() = got %v, want %v", got, want)
	}
}

func TestGermaniumGetSimbol(t *testing.T) {
	g := Germanium{}
	want := "Ge"
	got := g.GetSimbol()
	if got != want {
		t.Errorf("Germanium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestGermaniumGetAtomicNumber(t *testing.T) {
	g := Germanium{}
	want := 32
	got := g.GetAtomicNumber()
	if got != want {
		t.Errorf("Germanium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestGermaniumGetAtomicWeight(t *testing.T) {
	g := Germanium{}
	var want float32 = 72.630
	got := g.GetAtomicWeight()
	if got != want {
		t.Errorf("Germanium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
