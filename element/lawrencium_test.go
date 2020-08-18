package element

import "testing"

func TestLawrenciumGetPeriod(t *testing.T) {
	l := Lawrencium{}
	want := "7th period"
	got := l.GetPeriod()
	if got != want {
		t.Errorf("Lawrencium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestLawrenciumGetGroup(t *testing.T) {
	l := Lawrencium{}
	want := "3B"
	got := l.GetGroup()
	if got != want {
		t.Errorf("Lawrencium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestLawrenciumGetCategory(t *testing.T) {
	l := Lawrencium{}
	want := "Actinoid"
	got := l.GetCategory()
	if got != want {
		t.Errorf("Lawrencium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestLawrenciumGetName(t *testing.T) {
	l := Lawrencium{}
	want := "Lawrencium"
	got := l.GetName()
	if got != want {
		t.Errorf("Lawrencium.GetName() = got %v, want %v", got, want)
	}
}

func TestLawrenciumGetSimbol(t *testing.T) {
	l := Lawrencium{}
	want := "Lr"
	got := l.GetSimbol()
	if got != want {
		t.Errorf("Lawrencium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestLawrenciumGetAtomicNumber(t *testing.T) {
	l := Lawrencium{}
	want := 103
	got := l.GetAtomicNumber()
	if got != want {
		t.Errorf("Lawrencium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestLawrenciumGetAtomicWeight(t *testing.T) {
	l := Lawrencium{}
	var want float32 = 266
	got := l.GetAtomicWeight()
	if got != want {
		t.Errorf("Lawrencium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
