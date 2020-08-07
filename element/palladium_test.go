package element

import "testing"

func TestPalladiumGetPeriod(t *testing.T) {
	p := Palladium{}
	want := "5th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Palladium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPalladiumGetGroup(t *testing.T) {
	p := Palladium{}
	want := "8B"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Palladium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPalladiumGetCategory(t *testing.T) {
	p := Palladium{}
	want := "Transition metal"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Palladium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPalladiumGetName(t *testing.T) {
	p := Palladium{}
	want := "Palladium"
	got := p.GetName()
	if got != want {
		t.Errorf("Palladium.GetName() = got %v, want %v", got, want)
	}
}

func TestPalladiumGetSimbol(t *testing.T) {
	p := Palladium{}
	want := "Pd"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Palladium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPalladiumGetAtomicNumber(t *testing.T) {
	p := Palladium{}
	want := 46
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Palladium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPalladiumGetAtomicWeight(t *testing.T) {
	p := Palladium{}
	var want float32 = 106.42
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Palladium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
