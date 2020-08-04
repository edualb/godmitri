package element

import "testing"

func TestSiliconGetPeriod(t *testing.T) {
	s := Silicon{}
	want := "3rd period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Silicon.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSiliconGetGroup(t *testing.T) {
	s := Silicon{}
	want := "4A"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Silicon.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSiliconGetCategory(t *testing.T) {
	s := Silicon{}
	want := "Metalloid"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Silicon.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSiliconGetName(t *testing.T) {
	s := Silicon{}
	want := "Silicon"
	got := s.GetName()
	if got != want {
		t.Errorf("Silicon.GetName() = got %v, want %v", got, want)
	}
}

func TestSiliconGetSimbol(t *testing.T) {
	s := Silicon{}
	want := "Si"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Silicon.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSiliconGetAtomicNumber(t *testing.T) {
	s := Silicon{}
	want := 14
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Silicon.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSiliconGetAtomicWeight(t *testing.T) {
	s := Silicon{}
	var want float32 = 28.085
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Silicon.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
