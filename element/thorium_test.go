package element

import "testing"

func TestThoriumGetPeriod(t *testing.T) {
	tm := Thorium{}
	want := "7th period"
	got := tm.GetPeriod()
	if got != want {
		t.Errorf("Thorium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestThoriumGetGroup(t *testing.T) {
	tm := Thorium{}
	want := "3B"
	got := tm.GetGroup()
	if got != want {
		t.Errorf("Thorium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestThoriumGetCategory(t *testing.T) {
	tm := Thorium{}
	want := "Actinoid"
	got := tm.GetCategory()
	if got != want {
		t.Errorf("Thorium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestThoriumGetName(t *testing.T) {
	tm := Thorium{}
	want := "Thorium"
	got := tm.GetName()
	if got != want {
		t.Errorf("Thorium.GetName() = got %v, want %v", got, want)
	}
}

func TestThoriumGetSimbol(t *testing.T) {
	tm := Thorium{}
	want := "Th"
	got := tm.GetSimbol()
	if got != want {
		t.Errorf("Thorium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestThoriumGetAtomicNumber(t *testing.T) {
	tm := Thorium{}
	want := 90
	got := tm.GetAtomicNumber()
	if got != want {
		t.Errorf("Thorium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestThoriumGetAtomicWeight(t *testing.T) {
	tm := Thorium{}
	var want float32 = 232.04
	got := tm.GetAtomicWeight()
	if got != want {
		t.Errorf("Thorium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
