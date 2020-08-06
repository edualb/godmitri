package element

import "testing"

func TestGalliumGetPeriod(t *testing.T) {
	g := Gallium{}
	want := "4th period"
	got := g.GetPeriod()
	if got != want {
		t.Errorf("Gallium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestGalliumGetGroup(t *testing.T) {
	g := Gallium{}
	want := "3A"
	got := g.GetGroup()
	if got != want {
		t.Errorf("Gallium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestGalliumGetCategory(t *testing.T) {
	g := Gallium{}
	want := "Post-transition metal"
	got := g.GetCategory()
	if got != want {
		t.Errorf("Gallium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestGalliumGetName(t *testing.T) {
	g := Gallium{}
	want := "Gallium"
	got := g.GetName()
	if got != want {
		t.Errorf("Gallium.GetName() = got %v, want %v", got, want)
	}
}

func TestGalliumGetSimbol(t *testing.T) {
	g := Gallium{}
	want := "Ga"
	got := g.GetSimbol()
	if got != want {
		t.Errorf("Gallium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestGalliumGetAtomicNumber(t *testing.T) {
	g := Gallium{}
	want := 31
	got := g.GetAtomicNumber()
	if got != want {
		t.Errorf("Gallium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestGalliumGetAtomicWeight(t *testing.T) {
	g := Gallium{}
	var want float32 = 69.723
	got := g.GetAtomicWeight()
	if got != want {
		t.Errorf("Gallium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
