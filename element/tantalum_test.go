package element

import "testing"

func TestTantalumGetPeriod(t *testing.T) {
	ta := Tantalum{}
	want := "6th period"
	got := ta.GetPeriod()
	if got != want {
		t.Errorf("Tantalum.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTantalumGetGroup(t *testing.T) {
	ta := Tantalum{}
	want := "5B"
	got := ta.GetGroup()
	if got != want {
		t.Errorf("Tantalum.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTantalumGetCategory(t *testing.T) {
	ta := Tantalum{}
	want := "Transition metal"
	got := ta.GetCategory()
	if got != want {
		t.Errorf("Tantalum.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTantalumGetName(t *testing.T) {
	ta := Tantalum{}
	want := "Tantalum"
	got := ta.GetName()
	if got != want {
		t.Errorf("Tantalum.GetName() = got %v, want %v", got, want)
	}
}

func TestTantalumGetSimbol(t *testing.T) {
	ta := Tantalum{}
	want := "Ta"
	got := ta.GetSimbol()
	if got != want {
		t.Errorf("Tantalum.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTantalumGetAtomicNumber(t *testing.T) {
	ta := Tantalum{}
	want := 73
	got := ta.GetAtomicNumber()
	if got != want {
		t.Errorf("Tantalum.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTantalumGetAtomicWeight(t *testing.T) {
	ta := Tantalum{}
	var want float32 = 180.95
	got := ta.GetAtomicWeight()
	if got != want {
		t.Errorf("Tantalum.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
