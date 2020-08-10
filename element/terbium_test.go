package element

import "testing"

func TestTerbiumGetPeriod(t *testing.T) {
	te := Terbium{}
	want := "6th period"
	got := te.GetPeriod()
	if got != want {
		t.Errorf("Terbium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTerbiumGetGroup(t *testing.T) {
	te := Terbium{}
	want := "3B"
	got := te.GetGroup()
	if got != want {
		t.Errorf("Terbium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTerbiumGetCategory(t *testing.T) {
	te := Terbium{}
	want := "Lanthanoid"
	got := te.GetCategory()
	if got != want {
		t.Errorf("Terbium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTerbiumGetName(t *testing.T) {
	te := Terbium{}
	want := "Terbium"
	got := te.GetName()
	if got != want {
		t.Errorf("Terbium.GetName() = got %v, want %v", got, want)
	}
}

func TestTerbiumGetSimbol(t *testing.T) {
	te := Terbium{}
	want := "Tb"
	got := te.GetSimbol()
	if got != want {
		t.Errorf("Terbium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTerbiumGetAtomicNumber(t *testing.T) {
	te := Terbium{}
	want := 65
	got := te.GetAtomicNumber()
	if got != want {
		t.Errorf("Terbium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTerbiumGetAtomicWeight(t *testing.T) {
	te := Terbium{}
	var want float32 = 158.93
	got := te.GetAtomicWeight()
	if got != want {
		t.Errorf("Terbium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
