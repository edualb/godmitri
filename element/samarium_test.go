package element

import "testing"

func TestSamariumGetPeriod(t *testing.T) {
	s := Samarium{}
	want := "6th period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Samarium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSamariumGetGroup(t *testing.T) {
	s := Samarium{}
	want := "3B"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Samarium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSamariumGetCategory(t *testing.T) {
	s := Samarium{}
	want := "Lanthanoid"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Samarium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSamariumGetName(t *testing.T) {
	s := Samarium{}
	want := "Samarium"
	got := s.GetName()
	if got != want {
		t.Errorf("Samarium.GetName() = got %v, want %v", got, want)
	}
}

func TestSamariumGetSimbol(t *testing.T) {
	s := Samarium{}
	want := "Sm"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Samarium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSamariumGetAtomicNumber(t *testing.T) {
	s := Samarium{}
	want := 62
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Samarium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSamariumGetAtomicWeight(t *testing.T) {
	s := Samarium{}
	var want float32 = 150.36
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Samarium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
