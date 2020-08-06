package element

import "testing"

func TestSeleniumGetPeriod(t *testing.T) {
	s := Selenium{}
	want := "4th period"
	got := s.GetPeriod()
	if got != want {
		t.Errorf("Selenium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestSeleniumGetGroup(t *testing.T) {
	s := Selenium{}
	want := "6A"
	got := s.GetGroup()
	if got != want {
		t.Errorf("Selenium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestSeleniumGetCategory(t *testing.T) {
	s := Selenium{}
	want := "Non-metal"
	got := s.GetCategory()
	if got != want {
		t.Errorf("Selenium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestSeleniumGetName(t *testing.T) {
	s := Selenium{}
	want := "Selenium"
	got := s.GetName()
	if got != want {
		t.Errorf("Selenium.GetName() = got %v, want %v", got, want)
	}
}

func TestSeleniumGetSimbol(t *testing.T) {
	s := Selenium{}
	want := "Se"
	got := s.GetSimbol()
	if got != want {
		t.Errorf("Selenium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestSeleniumGetAtomicNumber(t *testing.T) {
	s := Selenium{}
	want := 34
	got := s.GetAtomicNumber()
	if got != want {
		t.Errorf("Selenium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestSeleniumGetAtomicWeight(t *testing.T) {
	s := Selenium{}
	var want float32 = 78.971
	got := s.GetAtomicWeight()
	if got != want {
		t.Errorf("Selenium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
