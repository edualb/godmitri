package element

import "testing"

func TestMoscoviumGetPeriod(t *testing.T) {
	m := Moscovium{}
	want := "7th period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Moscovium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestMoscoviumGetGroup(t *testing.T) {
	m := Moscovium{}
	want := "5A"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Moscovium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestMoscoviumGetCategory(t *testing.T) {
	m := Moscovium{}
	want := "Unknown"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Moscovium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestMoscoviumGetName(t *testing.T) {
	m := Moscovium{}
	want := "Moscovium"
	got := m.GetName()
	if got != want {
		t.Errorf("Moscovium.GetName() = got %v, want %v", got, want)
	}
}

func TestMoscoviumGetSimbol(t *testing.T) {
	m := Moscovium{}
	want := "Mc"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Moscovium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestMoscoviumGetAtomicNumber(t *testing.T) {
	m := Moscovium{}
	want := 115
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Moscovium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestMoscoviumGetAtomicWeight(t *testing.T) {
	m := Moscovium{}
	var want float32 = 290
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Moscovium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
