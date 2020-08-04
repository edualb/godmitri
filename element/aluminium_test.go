package element

import "testing"

func TestAluminiumGetPeriod(t *testing.T) {
	a := Aluminium{}
	want := "3rd period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Aluminium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestAluminiumGetGroup(t *testing.T) {
	a := Aluminium{}
	want := "3A"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Aluminium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestAluminiumGetCategory(t *testing.T) {
	a := Aluminium{}
	want := "Post-transition metal"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Aluminium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestAluminiumGetName(t *testing.T) {
	a := Aluminium{}
	want := "Aluminium"
	got := a.GetName()
	if got != want {
		t.Errorf("Aluminium.GetName() = got %v, want %v", got, want)
	}
}

func TestAluminiumGetSimbol(t *testing.T) {
	a := Aluminium{}
	want := "Al"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Aluminium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestAluminiumGetAtomicNumber(t *testing.T) {
	a := Aluminium{}
	want := 13
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Aluminium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestAluminiumGetAtomicWeight(t *testing.T) {
	a := Aluminium{}
	var want float32 = 26.982
	got := a.GetAtomicWeight()
	if got != want {
		t.Errorf("Aluminium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
