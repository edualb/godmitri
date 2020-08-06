package element

import "testing"

func TestArsenicGetPeriod(t *testing.T) {
	a := Arsenic{}
	want := "4th period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Arsenic.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestArsenicGetGroup(t *testing.T) {
	a := Arsenic{}
	want := "5A"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Arsenic.GetGroup() = got %v, want %v", got, want)
	}
}

func TestArsenicGetCategory(t *testing.T) {
	a := Arsenic{}
	want := "Metalloid"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Arsenic.GetCategory() = got %v, want %v", got, want)
	}
}

func TestArsenicGetName(t *testing.T) {
	a := Arsenic{}
	want := "Arsenic"
	got := a.GetName()
	if got != want {
		t.Errorf("Arsenic.GetName() = got %v, want %v", got, want)
	}
}

func TestArsenicGetSimbol(t *testing.T) {
	a := Arsenic{}
	want := "As"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Arsenic.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestArsenicGetAtomicNumber(t *testing.T) {
	a := Arsenic{}
	want := 33
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Arsenic.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestArsenicGetAtomicWeight(t *testing.T) {
	a := Arsenic{}
	var want float32 = 74.922
	got := a.GetAtomicWeight()
	if got != want {
		t.Errorf("Arsenic.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
