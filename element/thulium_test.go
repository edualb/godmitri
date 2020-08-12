package element

import "testing"

func TestThuliumGetPeriod(t *testing.T) {
	tm := Thulium{}
	want := "6th period"
	got := tm.GetPeriod()
	if got != want {
		t.Errorf("Thulium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestThuliumGetGroup(t *testing.T) {
	tm := Thulium{}
	want := "3B"
	got := tm.GetGroup()
	if got != want {
		t.Errorf("Thulium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestThuliumGetCategory(t *testing.T) {
	tm := Thulium{}
	want := "Lanthanoid"
	got := tm.GetCategory()
	if got != want {
		t.Errorf("Thulium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestThuliumGetName(t *testing.T) {
	tm := Thulium{}
	want := "Thulium"
	got := tm.GetName()
	if got != want {
		t.Errorf("Thulium.GetName() = got %v, want %v", got, want)
	}
}

func TestThuliumGetSimbol(t *testing.T) {
	tm := Thulium{}
	want := "Tm"
	got := tm.GetSimbol()
	if got != want {
		t.Errorf("Thulium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestThuliumGetAtomicNumber(t *testing.T) {
	tm := Thulium{}
	want := 69
	got := tm.GetAtomicNumber()
	if got != want {
		t.Errorf("Thulium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestThuliumGetAtomicWeight(t *testing.T) {
	tm := Thulium{}
	var want float32 = 168.93
	got := tm.GetAtomicWeight()
	if got != want {
		t.Errorf("Thulium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
