package element

import "testing"

func TestSilverGetPeriod(t *testing.T) {
	s := Silver{}
	want := "5th period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Silver.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSilverGetGroup(t *testing.T) {
	s := Silver{}
	want := "1B"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Silver.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSilverGetCategory(t *testing.T) {
	s := Silver{}
	want := "Transition metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Silver.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSilverGetName(t *testing.T) {
	s := Silver{}
	want := "Silver"
	got := s.GetName()
	if got != want {
		t.Errorf("Silver.GetName() = got %v, want %v", got, want)
	}
}

func TestSilverGetSimbol(t *testing.T) {
	s := Silver{}
	want := "Ag"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Silver.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSilverGetAtomicNumber(t *testing.T) {
	s := Silver{}
	want := 47
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Silver.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSilverGetAtomicWeight(t *testing.T) {
	s := Silver{}
	var want float32 = 107.87
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Silver.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
