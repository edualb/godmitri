package element

import "testing"

func TestMolybdenumGetPeriod(t *testing.T) {
	m := Molybdenum{}
	want := "5th period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Molybdenum.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestMolybdenumGetGroup(t *testing.T) {
	m := Molybdenum{}
	want := "6B"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Molybdenum.GetGroup() = got %v, want %v", got, want)
	}
}

func TestMolybdenumGetCategory(t *testing.T) {
	m := Molybdenum{}
	want := "Transition metal"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Molybdenum.GetCategory() = got %v, want %v", got, want)
	}
}

func TestMolybdenumGetName(t *testing.T) {
	m := Molybdenum{}
	want := "Molybdenum"
	got := m.GetName()
	if got != want {
		t.Errorf("Molybdenum.GetName() = got %v, want %v", got, want)
	}
}

func TestMolybdenumGetSimbol(t *testing.T) {
	m := Molybdenum{}
	want := "Mo"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Molybdenum.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestMolybdenumGetAtomicNumber(t *testing.T) {
	m := Molybdenum{}
	want := 42
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Molybdenum.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestMolybdenumGetAtomicWeight(t *testing.T) {
	m := Molybdenum{}
	var want float32 = 95.95
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Molybdenum.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
