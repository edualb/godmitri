package element

import "testing"

func TestPotassiumGetPeriod(t *testing.T) {
	p := Potassium{}
	want := "4th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Potassium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPotassiumGetGroup(t *testing.T) {
	p := Potassium{}
	want := "1A"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Potassium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPotassiumGetCategory(t *testing.T) {
	p := Potassium{}
	want := "Alkali metal"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Potassium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPotassiumGetName(t *testing.T) {
	p := Potassium{}
	want := "Potassium"
	got := p.GetName()
	if got != want {
		t.Errorf("Potassium.GetName() = got %v, want %v", got, want)
	}
}

func TestPotassiumGetSimbol(t *testing.T) {
	p := Potassium{}
	want := "K"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Potassium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPotassiumGetAtomicNumber(t *testing.T) {
	p := Potassium{}
	want := 19
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Potassium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPotassiumGetAtomicWeight(t *testing.T) {
	p := Potassium{}
	var want float32 = 39.098
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Potassium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
