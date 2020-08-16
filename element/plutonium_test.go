package element

import "testing"

func TestPlutoniumGetPeriod(t *testing.T) {
	p := Plutonium{}
	want := "7th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Plutonium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPlutoniumGetGroup(t *testing.T) {
	p := Plutonium{}
	want := "3B"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Plutonium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPlutoniumGetCategory(t *testing.T) {
	p := Plutonium{}
	want := "Actinoid"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Plutonium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPlutoniumGetName(t *testing.T) {
	p := Plutonium{}
	want := "Plutonium"
	got := p.GetName()
	if got != want {
		t.Errorf("Plutonium.GetName() = got %v, want %v", got, want)
	}
}

func TestPlutoniumGetSimbol(t *testing.T) {
	p := Plutonium{}
	want := "Pu"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Plutonium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPlutoniumGetAtomicNumber(t *testing.T) {
	p := Plutonium{}
	want := 94
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Plutonium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPlutoniumGetAtomicWeight(t *testing.T) {
	p := Plutonium{}
	var want float32 = 244
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Plutonium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
