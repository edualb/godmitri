package element

import "testing"

func TestZincGetPeriod(t *testing.T) {
	z := Zinc{}
	want := "4th period"
	got := z.GetPeriod()
	if got != want {
		t.Errorf("Zinc.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestZincGetGroup(t *testing.T) {
	z := Zinc{}
	want := "2B"
	got := z.GetGroup()
	if got != want {
		t.Errorf("Zinc.GetGroup() = got %v, want %v", got, want)
	}
}

func TestZincGetCategory(t *testing.T) {
	z := Zinc{}
	want := "Transition metal"
	got := z.GetCategory()
	if got != want {
		t.Errorf("Zinc.GetCategory() = got %v, want %v", got, want)
	}
}

func TestZincGetName(t *testing.T) {
	z := Zinc{}
	want := "Zinc"
	got := z.GetName()
	if got != want {
		t.Errorf("Zinc.GetName() = got %v, want %v", got, want)
	}
}

func TestZincGetSimbol(t *testing.T) {
	z := Zinc{}
	want := "Zn"
	got := z.GetSimbol()
	if got != want {
		t.Errorf("Zinc.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestZincGetAtomicNumber(t *testing.T) {
	z := Zinc{}
	want := 30
	got := z.GetAtomicNumber()
	if got != want {
		t.Errorf("Zinc.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestZincGetAtomicWeight(t *testing.T) {
	z := Zinc{}
	var want float32 = 65.38
	got := z.GetAtomicWeight()
	if got != want {
		t.Errorf("Zinc.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
