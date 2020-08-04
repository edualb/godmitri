package element

import "testing"

func TestNitrogenGetPeriod(t *testing.T) {
	n := Nitrogen{}
	want := "2nd period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Nitrogen.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNitrogenGetGroup(t *testing.T) {
	n := Nitrogen{}
	want := "5A"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Nitrogen.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNitrogenGetCategory(t *testing.T) {
	n := Nitrogen{}
	want := "Non-metal"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Nitrogen.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNitrogenGetName(t *testing.T) {
	n := Nitrogen{}
	want := "Nitrogen"
	got := n.GetName()
	if got != want {
		t.Errorf("Nitrogen.GetName() = got %v, want %v", got, want)
	}
}

func TestNitrogenGetSimbol(t *testing.T) {
	n := Nitrogen{}
	want := "N"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Nitrogen.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNitrogenGetAtomicNumber(t *testing.T) {
	n := Nitrogen{}
	want := 7
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Nitrogen.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNitrogenGetAtomicWeight(t *testing.T) {
	n := Nitrogen{}
	var want float32 = 14.007
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Nitrogen.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
