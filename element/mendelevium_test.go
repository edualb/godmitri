package element

import "testing"

func TestMendeleviumGetPeriod(t *testing.T) {
	m := Mendelevium{}
	want := "7th period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Mendelevium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestMendeleviumGetGroup(t *testing.T) {
	m := Mendelevium{}
	want := "3B"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Mendelevium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestMendeleviumGetCategory(t *testing.T) {
	m := Mendelevium{}
	want := "Actinoid"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Mendelevium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestMendeleviumGetName(t *testing.T) {
	m := Mendelevium{}
	want := "Mendelevium"
	got := m.GetName()
	if got != want {
		t.Errorf("Mendelevium.GetName() = got %v, want %v", got, want)
	}
}

func TestMendeleviumGetSimbol(t *testing.T) {
	m := Mendelevium{}
	want := "Md"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Mendelevium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestMendeleviumGetAtomicNumber(t *testing.T) {
	m := Mendelevium{}
	want := 101
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Mendelevium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestMendeleviumGetAtomicWeight(t *testing.T) {
	m := Mendelevium{}
	var want float32 = 258
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Mendelevium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
