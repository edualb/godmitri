package element

import "testing"

func TestPraseodymiumGetPeriod(t *testing.T) {
	p := Praseodymium{}
	want := "6th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Praseodymium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPraseodymiumGetGroup(t *testing.T) {
	p := Praseodymium{}
	want := "3B"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Praseodymium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPraseodymiumGetCategory(t *testing.T) {
	p := Praseodymium{}
	want := "Lanthanoid"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Praseodymium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPraseodymiumGetName(t *testing.T) {
	p := Praseodymium{}
	want := "Praseodymium"
	got := p.GetName()
	if got != want {
		t.Errorf("Praseodymium.GetName() = got %v, want %v", got, want)
	}
}

func TestPraseodymiumGetSimbol(t *testing.T) {
	p := Praseodymium{}
	want := "Pr"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Praseodymium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPraseodymiumGetAtomicNumber(t *testing.T) {
	p := Praseodymium{}
	want := 59
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Praseodymium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPraseodymiumGetAtomicWeight(t *testing.T) {
	p := Praseodymium{}
	var want float32 = 140.91
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Praseodymium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
