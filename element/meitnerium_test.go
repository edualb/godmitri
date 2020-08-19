package element

import "testing"

func TestMeitneriumGetPeriod(t *testing.T) {
	m := Meitnerium{}
	want := "7th period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Meitnerium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestMeitneriumGetGroup(t *testing.T) {
	m := Meitnerium{}
	want := "8B"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Meitnerium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestMeitneriumGetCategory(t *testing.T) {
	m := Meitnerium{}
	want := "Unknown"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Meitnerium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestMeitneriumGetName(t *testing.T) {
	m := Meitnerium{}
	want := "Meitnerium"
	got := m.GetName()
	if got != want {
		t.Errorf("Meitnerium.GetName() = got %v, want %v", got, want)
	}
}

func TestMeitneriumGetSimbol(t *testing.T) {
	m := Meitnerium{}
	want := "Mt"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Meitnerium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestMeitneriumGetAtomicNumber(t *testing.T) {
	m := Meitnerium{}
	want := 109
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Meitnerium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestMeitneriumGetAtomicWeight(t *testing.T) {
	m := Meitnerium{}
	var want float32 = 278
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Meitnerium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
