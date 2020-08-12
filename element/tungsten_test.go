package element

import "testing"

func TestTungstenGetPeriod(t *testing.T) {
	tu := Tungsten{}
	want := "6th period"
	got := tu.GetPeriod()
	if got != want {
		t.Errorf("Tungsten.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTungstenGetGroup(t *testing.T) {
	tu := Tungsten{}
	want := "6B"
	got := tu.GetGroup()
	if got != want {
		t.Errorf("Tungsten.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTungstenGetCategory(t *testing.T) {
	tu := Tungsten{}
	want := "Transition metal"
	got := tu.GetCategory()
	if got != want {
		t.Errorf("Tungsten.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTungstenGetName(t *testing.T) {
	tu := Tungsten{}
	want := "Tungsten"
	got := tu.GetName()
	if got != want {
		t.Errorf("Tungsten.GetName() = got %v, want %v", got, want)
	}
}

func TestTungstenGetSimbol(t *testing.T) {
	tu := Tungsten{}
	want := "W"
	got := tu.GetSimbol()
	if got != want {
		t.Errorf("Tungsten.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTungstenGetAtomicNumber(t *testing.T) {
	ta := Tungsten{}
	want := 74
	got := ta.GetAtomicNumber()
	if got != want {
		t.Errorf("Tungsten.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTungstenGetAtomicWeight(t *testing.T) {
	ta := Tungsten{}
	var want float32 = 183.84
	got := ta.GetAtomicWeight()
	if got != want {
		t.Errorf("Tungsten.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
