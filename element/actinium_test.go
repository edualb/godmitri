package element

import "testing"

func TestActiniumGetPeriod(t *testing.T) {
	a := Actinium{}
	want := "7th period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Actinium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestActiniumGetGroup(t *testing.T) {
	a := Actinium{}
	want := "3B"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Actinium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestActiniumGetCategory(t *testing.T) {
	a := Actinium{}
	want := "Actinoid"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Actinium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestActiniumGetName(t *testing.T) {
	a := Actinium{}
	want := "Actinium"
	got := a.GetName()
	if got != want {
		t.Errorf("Actinium.GetName() = got %v, want %v", got, want)
	}
}

func TestActiniumGetSimbol(t *testing.T) {
	a := Actinium{}
	want := "Ac"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Actinium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestActiniumGetAtomicNumber(t *testing.T) {
	a := Actinium{}
	want := 89
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Actinium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestActiniumGetAtomicWeight(t *testing.T) {
	a := Actinium{}
	var want float32 = 227
	got := a.GetAtomicWeight()
	if got != want {
		t.Errorf("Actinium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
