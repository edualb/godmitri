package element

import "testing"

func TestSodiumGetPeriod(t *testing.T) {
	s := Sodium{}
	want := "3rd period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Sodium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSodiumGetGroup(t *testing.T) {
	s := Sodium{}
	want := "1A"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Sodium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSodiumGetCategory(t *testing.T) {
	s := Sodium{}
	want := "Alkali metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Sodium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSodiumGetName(t *testing.T) {
	s := Sodium{}
	want := "Sodium"
	got := s.GetName()
	if got != want {
		t.Errorf("Sodium.GetName() = got %v, want %v", got, want)
	}
}

func TestSodiumGetSimbol(t *testing.T) {
	s := Sodium{}
	want := "Na"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Sodium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSodiumGetAtomicNumber(t *testing.T) {
	s := Sodium{}
	want := 11
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Sodium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSodiumGetAtomicWeight(t *testing.T) {
	s := Sodium{}
	var want float32 = 22.990
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Sodium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
