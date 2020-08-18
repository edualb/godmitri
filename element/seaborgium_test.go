package element

import "testing"

func TestSeaborgiumGetPeriod(t *testing.T) {
	s := Seaborgium{}
	want := "7th period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Seaborgium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSeaborgiumGetGroup(t *testing.T) {
	s := Seaborgium{}
	want := "6B"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Seaborgium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSeaborgiumGetCategory(t *testing.T) {
	s := Seaborgium{}
	want := "Transition metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Seaborgium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSeaborgiumGetName(t *testing.T) {
	s := Seaborgium{}
	want := "Seaborgium"
	got := s.GetName()
	if got != want {
		t.Errorf("Seaborgium.GetName() = got %v, want %v", got, want)
	}
}

func TestSeaborgiumGetSimbol(t *testing.T) {
	s := Seaborgium{}
	want := "Sg"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Seaborgium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSeaborgiumGetAtomicNumber(t *testing.T) {
	s := Seaborgium{}
	want := 106
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Seaborgium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSeaborgiumGetAtomicWeight(t *testing.T) {
	s := Seaborgium{}
	var want float32 = 269
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Seaborgium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
