package element

import "testing"

func TestStrontiumGetPeriod(t *testing.T) {
	s := Strontium{}
	want := "5th period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Strontium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestStrontiumGetGroup(t *testing.T) {
	s := Strontium{}
	want := "2A"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Strontium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestStrontiumGetCategory(t *testing.T) {
	s := Strontium{}
	want := "Alkaline earth metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Strontium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestStrontiumGetName(t *testing.T) {
	s := Strontium{}
	want := "Strontium"
	got := s.GetName()
	if got != want {
		t.Errorf("Strontium.GetName() = got %v, want %v", got, want)
	}
}

func TestStrontiumGetSimbol(t *testing.T) {
	s := Strontium{}
	want := "Sr"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Strontium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestStrontiumGetAtomicNumber(t *testing.T) {
	s := Strontium{}
	want := 38
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Strontium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestStrontiumGetAtomicWeight(t *testing.T) {
	s := Strontium{}
	var want float32 = 87.62
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Strontium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
