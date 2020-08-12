package element

import "testing"

func TestLutetiumGetPeriod(t *testing.T) {
	l := Lutetium{}
	want := "6th period"
	got := l.GetPeriod()
	if got != want {
		t.Errorf("Lutetium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestLutetiumGetGroup(t *testing.T) {
	l := Lutetium{}
	want := "3B"
	got := l.GetGroup()
	if got != want {
		t.Errorf("Lutetium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestLutetiumGetCategory(t *testing.T) {
	l := Lutetium{}
	want := "Lanthanoid"
	got := l.GetCategory()
	if got != want {
		t.Errorf("Lutetium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestLutetiumGetName(t *testing.T) {
	l := Lutetium{}
	want := "Lutetium"
	got := l.GetName()
	if got != want {
		t.Errorf("Lutetium.GetName() = got %v, want %v", got, want)
	}
}

func TestLutetiumGetSimbol(t *testing.T) {
	l := Lutetium{}
	want := "Lu"
	got := l.GetSimbol()
	if got != want {
		t.Errorf("Lithium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestLutetiumGetAtomicNumber(t *testing.T) {
	l := Lutetium{}
	want := 71
	got := l.GetAtomicNumber()
	if got != want {
		t.Errorf("Lutetium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestLutetiumGetAtomicWeight(t *testing.T) {
	l := Lutetium{}
	var want float32 = 174.97
	got := l.GetAtomicWeight()
	if got != want {
		t.Errorf("Lutetium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
