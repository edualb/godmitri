package element

import "testing"

func TestTechnetiumGetPeriod(t *testing.T) {
	te := Technetium{}
	want := "5th period"
	got := te.GetPeriod()
	if got != want {
		t.Errorf("Technetium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTechnetiumGetGroup(t *testing.T) {
	te := Technetium{}
	want := "7B"
	got := te.GetGroup()
	if got != want {
		t.Errorf("Technetium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTechnetiumGetCategory(t *testing.T) {
	te := Technetium{}
	want := "Transition metal"
	got := te.GetCategory()
	if got != want {
		t.Errorf("Technetium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTechnetiumGetName(t *testing.T) {
	te := Technetium{}
	want := "Technetium"
	got := te.GetName()
	if got != want {
		t.Errorf("Technetium.GetName() = got %v, want %v", got, want)
	}
}

func TestTechnetiumGetSimbol(t *testing.T) {
	te := Technetium{}
	want := "Tc"
	got := te.GetSimbol()
	if got != want {
		t.Errorf("Technetium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTechnetiumGetAtomicNumber(t *testing.T) {
	te := Technetium{}
	want := 43
	got := te.GetAtomicNumber()
	if got != want {
		t.Errorf("Technetium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTechnetiumGetAtomicWeight(t *testing.T) {
	te := Technetium{}
	var want float32 = 98
	got := te.GetAtomicWeight()
	if got != want {
		t.Errorf("Technetium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
