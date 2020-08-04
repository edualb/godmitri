package element

import "testing"

func TestArgonGetPeriod(t *testing.T) {
	a := Argon{}
	want := "3rd period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Argon.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestArgonGetGroup(t *testing.T) {
	a := Argon{}
	want := "8A"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Argon.GetGroup() = got %v, want %v", got, want)
	}
}

func TestArgonGetCategory(t *testing.T) {
	a := Argon{}
	want := "Noble gas"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Argon.GetCategory() = got %v, want %v", got, want)
	}
}

func TestArgonGetName(t *testing.T) {
	a := Argon{}
	want := "Argon"
	got := a.GetName()
	if got != want {
		t.Errorf("Argon.GetName() = got %v, want %v", got, want)
	}
}

func TestArgonGetSimbol(t *testing.T) {
	a := Argon{}
	want := "Ar"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Argon.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestArgonGetAtomicNumber(t *testing.T) {
	a := Argon{}
	want := 18
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Argon.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestArgonGetAtomicWeight(t *testing.T) {
	c := Argon{}
	var want float32 = 39.948
	got := c.GetAtomicWeight()
	if got != want {
		t.Errorf("Chlorine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
