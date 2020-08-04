package element

import "testing"

func TestFluorineGetPeriod(t *testing.T) {
	f := Fluorine{}
	want := "2nd period"
	got := f.GetPeriod()
	if got != want {
		t.Errorf("Fluorine.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestFluorineGetGroup(t *testing.T) {
	f := Fluorine{}
	want := "7A"
	got := f.GetGroup()
	if got != want {
		t.Errorf("Fluorine.GetGroup() = got %v, want %v", got, want)
	}
}

func TestFluorineGetCategory(t *testing.T) {
	f := Fluorine{}
	want := "Non-metal"
	got := f.GetCategory()
	if got != want {
		t.Errorf("Fluorine.GetCategory() = got %v, want %v", got, want)
	}
}

func TestFluorineGetName(t *testing.T) {
	f := Fluorine{}
	want := "Fluorine"
	got := f.GetName()
	if got != want {
		t.Errorf("Fluorine.GetName() = got %v, want %v", got, want)
	}
}

func TestFluorineGetSimbol(t *testing.T) {
	f := Fluorine{}
	want := "F"
	got := f.GetSimbol()
	if got != want {
		t.Errorf("Fluorine.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestFluorineGetAtomicNumber(t *testing.T) {
	f := Fluorine{}
	want := 9
	got := f.GetAtomicNumber()
	if got != want {
		t.Errorf("Fluorine.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestFluorineGetAtomicWeight(t *testing.T) {
	f := Fluorine{}
	var want float32 = 18.998
	got := f.GetAtomicWeight()
	if got != want {
		t.Errorf("Fluorine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
