package element

import "testing"

func TestUraniumGetPeriod(t *testing.T) {
	u := Uranium{}
	want := "7th period"
	got := u.GetPeriod()
	if got != want {
		t.Errorf("Uranium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestUraniumGetGroup(t *testing.T) {
	u := Uranium{}
	want := "3B"
	got := u.GetGroup()
	if got != want {
		t.Errorf("Uranium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestUraniumGetCategory(t *testing.T) {
	u := Uranium{}
	want := "Actinoid"
	got := u.GetCategory()
	if got != want {
		t.Errorf("Uranium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestUraniumGetName(t *testing.T) {
	u := Uranium{}
	want := "Uranium"
	got := u.GetName()
	if got != want {
		t.Errorf("Uranium.GetName() = got %v, want %v", got, want)
	}
}

func TestUraniumGetSimbol(t *testing.T) {
	u := Uranium{}
	want := "U"
	got := u.GetSimbol()
	if got != want {
		t.Errorf("Uranium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestUraniumGetAtomicNumber(t *testing.T) {
	u := Uranium{}
	want := 92
	got := u.GetAtomicNumber()
	if got != want {
		t.Errorf("Uranium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestUraniumGetAtomicWeight(t *testing.T) {
	u := Uranium{}
	var want float32 = 238.03
	got := u.GetAtomicWeight()
	if got != want {
		t.Errorf("Uranium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
