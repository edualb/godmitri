package element

import "testing"

func TestProtactiniumGetPeriod(t *testing.T) {
	p := Protactinium{}
	want := "7th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Protactinium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestProtactiniumGetGroup(t *testing.T) {
	p := Protactinium{}
	want := "3B"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Protactinium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestProtactiniumGetCategory(t *testing.T) {
	p := Protactinium{}
	want := "Actinoid"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Protactinium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestProtactiniumGetName(t *testing.T) {
	p := Protactinium{}
	want := "Protactinium"
	got := p.GetName()
	if got != want {
		t.Errorf("Protactinium.GetName() = got %v, want %v", got, want)
	}
}

func TestProtactiniumGetSimbol(t *testing.T) {
	p := Protactinium{}
	want := "Pa"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Protactinium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestProtactiniumGetAtomicNumber(t *testing.T) {
	p := Protactinium{}
	want := 91
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Protactinium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestProtactiniumGetAtomicWeight(t *testing.T) {
	p := Protactinium{}
	var want float32 = 231.04
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Protactinium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
