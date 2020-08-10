package element

import "testing"

func TestPromethiumGetPeriod(t *testing.T) {
	p := Promethium{}
	want := "6th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Promethium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPromethiumGetGroup(t *testing.T) {
	p := Promethium{}
	want := "3B"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Promethium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPromethiumGetCategory(t *testing.T) {
	p := Promethium{}
	want := "Lanthanoid"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Promethium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPromethiumGetName(t *testing.T) {
	p := Promethium{}
	want := "Promethium"
	got := p.GetName()
	if got != want {
		t.Errorf("Promethium.GetName() = got %v, want %v", got, want)
	}
}

func TestPromethiumGetSimbol(t *testing.T) {
	p := Promethium{}
	want := "Pm"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Promethium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPromethiumGetAtomicNumber(t *testing.T) {
	p := Promethium{}
	want := 61
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Promethium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPromethiumGetAtomicWeight(t *testing.T) {
	p := Promethium{}
	var want float32 = 145
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Promethium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
