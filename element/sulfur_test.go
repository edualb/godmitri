package element

import "testing"

func TestSulfurGetPeriod(t *testing.T) {
	s := Sulfur{}
	want := "3rd period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Sulfur.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSulfurGetGroup(t *testing.T) {
	s := Sulfur{}
	want := "6A"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Sulfur.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSulfurGetCategory(t *testing.T) {
	s := Sulfur{}
	want := "Non-metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Sulfur.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSulfurGetName(t *testing.T) {
	s := Sulfur{}
	want := "Sulfur"
	got := s.GetName()
	if got != want {
		t.Errorf("Sulfur.GetName() = got %v, want %v", got, want)
	}
}

func TestSulfurGetSimbol(t *testing.T) {
	s := Sulfur{}
	want := "S"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Sulfur.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSulfurGetAtomicNumber(t *testing.T) {
	p := Sulfur{}
	want := 16
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Sulfur.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSulfurGetAtomicWeight(t *testing.T) {
	s := Sulfur{}
	var want float32 = 32.06
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Sulfur.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
