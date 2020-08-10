package element

import "testing"

func TestLanthanumGetPeriod(t *testing.T) {
	l := Lanthanum{}
	want := "6th period"
	got := l.GetPeriod()
	if got != want {
		t.Errorf("Lanthanum.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestLanthanumGetGroup(t *testing.T) {
	l := Lanthanum{}
	want := "3B"
	got := l.GetGroup()
	if got != want {
		t.Errorf("Lanthanum.GetGroup() = got %v, want %v", got, want)
	}
}

func TestLanthanumGetCategory(t *testing.T) {
	l := Lanthanum{}
	want := "Lanthanoid"
	got := l.GetCategory()
	if got != want {
		t.Errorf("Lanthanum.GetCategory() = got %v, want %v", got, want)
	}
}

func TestLanthanumGetName(t *testing.T) {
	l := Lanthanum{}
	want := "Lanthanum"
	got := l.GetName()
	if got != want {
		t.Errorf("Lanthanum.GetName() = got %v, want %v", got, want)
	}
}

func TestLanthanumGetSimbol(t *testing.T) {
	l := Lanthanum{}
	want := "La"
	got := l.GetSimbol()
	if got != want {
		t.Errorf("Lanthanum.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestLanthanumGetAtomicNumber(t *testing.T) {
	l := Lanthanum{}
	want := 57
	got := l.GetAtomicNumber()
	if got != want {
		t.Errorf("Lanthanum.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestLanthanumGetAtomicWeight(t *testing.T) {
	l := Lanthanum{}
	var want float32 = 138.91
	got := l.GetAtomicWeight()
	if got != want {
		t.Errorf("Lanthanum.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
