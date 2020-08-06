package element

import "testing"

func TestScandiumGetPeriod(t *testing.T) {
	s := Scandium{}
	want := "4th period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Scandium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestScandiumGetGroup(t *testing.T) {
	s := Scandium{}
	want := "3B"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Scandium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestScandiumGetCategory(t *testing.T) {
	s := Scandium{}
	want := "Transition metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Scandium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestScandiumGetName(t *testing.T) {
	s := Scandium{}
	want := "Scandium"
	got := s.GetName()
	if got != want {
		t.Errorf("Scandium.GetName() = got %v, want %v", got, want)
	}
}

func TestScandiumGetSimbol(t *testing.T) {
	s := Scandium{}
	want := "Sc"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Scandium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestScandiumGetAtomicNumber(t *testing.T) {
	s := Scandium{}
	want := 21
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Scandium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestScandiumGetAtomicWeight(t *testing.T) {
	s := Scandium{}
	var want float32 = 44.956
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Scandium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
