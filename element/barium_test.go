package element

import "testing"

func TestBariumGetPeriod(t *testing.T) {
	b := Barium{}
	want := "6th period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Barium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBariumGetGroup(t *testing.T) {
	b := Barium{}
	want := "2A"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Barium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBariumGetCategory(t *testing.T) {
	b := Barium{}
	want := "Alkaline earth metal"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Barium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBariumGetName(t *testing.T) {
	b := Barium{}
	want := "Barium"
	got := b.GetName()
	if got != want {
		t.Errorf("Barium.GetName() = got %v, want %v", got, want)
	}
}

func TestBariumGetSimbol(t *testing.T) {
	b := Barium{}
	want := "Ba"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Barium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBariumGetAtomicNumber(t *testing.T) {
	b := Barium{}
	want := 56
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Barium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBariumGetAtomicWeight(t *testing.T) {
	b := Barium{}
	var want float32 = 137.33
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Barium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
