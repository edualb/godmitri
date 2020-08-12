package element

import "testing"

func TestYtterbiumGetPeriod(t *testing.T) {
	y := Ytterbium{}
	want := "6th period"
	got := y.GetPeriod()
	if got != want {
		t.Errorf("YtYtterbiumtrium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestYtterbiumGetGroup(t *testing.T) {
	y := Ytterbium{}
	want := "3B"
	got := y.GetGroup()
	if got != want {
		t.Errorf("Ytterbium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestYtterbiumGetCategory(t *testing.T) {
	y := Ytterbium{}
	want := "Lanthanoid"
	got := y.GetCategory()
	if got != want {
		t.Errorf("Ytterbium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestYtterbiumGetName(t *testing.T) {
	y := Ytterbium{}
	want := "Ytterbium"
	got := y.GetName()
	if got != want {
		t.Errorf("Ytterbium.GetName() = got %v, want %v", got, want)
	}
}

func TestYtterbiumGetSimbol(t *testing.T) {
	y := Ytterbium{}
	want := "Yb"
	got := y.GetSimbol()
	if got != want {
		t.Errorf("Ytterbium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestYtterbiumGetAtomicNumber(t *testing.T) {
	y := Ytterbium{}
	want := 70
	got := y.GetAtomicNumber()
	if got != want {
		t.Errorf("Ytterbium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestYtterbiumGetAtomicWeight(t *testing.T) {
	y := Ytterbium{}
	var want float32 = 173.05
	got := y.GetAtomicWeight()
	if got != want {
		t.Errorf("Ytterbium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
