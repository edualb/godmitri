package element

import "testing"

func TestMagnesiumGetPeriod(t *testing.T) {
	m := Magnesium{}
	want := "3rd period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Magnesium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestMagnesiumGetGroup(t *testing.T) {
	m := Magnesium{}
	want := "2A"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Magnesium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestMagnesiumGetCategory(t *testing.T) {
	m := Magnesium{}
	want := "Alkaline earth metal"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Magnesium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestMagnesiumGetName(t *testing.T) {
	m := Magnesium{}
	want := "Magnesium"
	got := m.GetName()
	if got != want {
		t.Errorf("Magnesium.GetName() = got %v, want %v", got, want)
	}
}

func TestMagnesiumGetSimbol(t *testing.T) {
	m := Magnesium{}
	want := "Mg"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Magnesium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestMagnesiumGetAtomicNumber(t *testing.T) {
	m := Magnesium{}
	want := 12
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Magnesium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestMagnesiumGetAtomicWeight(t *testing.T) {
	m := Magnesium{}
	var want float32 = 24.305
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Magnesium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
