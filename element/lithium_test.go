package element

import "testing"

func TestLithiumGetPeriod(t *testing.T) {
	l := Lithium{}
	want := "2nd period"
	got := l.GetPeriod()
	if got != want {
		t.Errorf("Lithium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestLithiumGetGroup(t *testing.T) {
	l := Lithium{}
	want := "1A"
	got := l.GetGroup()
	if got != want {
		t.Errorf("Lithium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestLithiumGetCategory(t *testing.T) {
	l := Lithium{}
	want := "Alkali metal"
	got := l.GetCategory()
	if got != want {
		t.Errorf("Lithium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestLithiumGetName(t *testing.T) {
	l := Lithium{}
	want := "Lithium"
	got := l.GetName()
	if got != want {
		t.Errorf("Lithium.GetName() = got %v, want %v", got, want)
	}
}

func TestLithiumGetSimbol(t *testing.T) {
	l := Lithium{}
	want := "Li"
	got := l.GetSimbol()
	if got != want {
		t.Errorf("Lithium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestLithiumGetAtomicNumber(t *testing.T) {
	l := Lithium{}
	want := 3
	got := l.GetAtomicNumber()
	if got != want {
		t.Errorf("Lithium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestLithiumGetAtomicWeight(t *testing.T) {
	l := Lithium{}
	var want float32 = 6.94
	got := l.GetAtomicWeight()
	if got != want {
		t.Errorf("Lithium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
