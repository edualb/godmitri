package element

import "testing"

func TestMercuryGetPeriod(t *testing.T) {
	m := Mercury{}
	want := "6th period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Mercury.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestMercuryGetGroup(t *testing.T) {
	m := Mercury{}
	want := "2B"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Mercury.GetGroup() = got %v, want %v", got, want)
	}
}

func TestMercuryGetCategory(t *testing.T) {
	m := Mercury{}
	want := "Transition metal"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Mercury.GetCategory() = got %v, want %v", got, want)
	}
}

func TestMercuryGetName(t *testing.T) {
	m := Mercury{}
	want := "Mercury"
	got := m.GetName()
	if got != want {
		t.Errorf("Mercury.GetName() = got %v, want %v", got, want)
	}
}

func TestMercuryGetSimbol(t *testing.T) {
	m := Mercury{}
	want := "Hg"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Mercury.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestMercuryGetAtomicNumber(t *testing.T) {
	m := Mercury{}
	want := 80
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Mercury.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestMercuryGetAtomicWeight(t *testing.T) {
	m := Mercury{}
	var want float32 = 200.59
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Mercury.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
