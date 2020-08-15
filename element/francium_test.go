package element

import "testing"

func TestFranciumGetPeriod(t *testing.T) {
	f := Francium{}
	want := "7th period"
	got := f.GetPeriod()
	if got != want {
		t.Errorf("Francium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestFranciumGetGroup(t *testing.T) {
	f := Francium{}
	want := "1A"
	got := f.GetGroup()
	if got != want {
		t.Errorf("Francium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestFranciumGetCategory(t *testing.T) {
	f := Francium{}
	want := "Alkali metal"
	got := f.GetCategory()
	if got != want {
		t.Errorf("Francium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestFranciumGetName(t *testing.T) {
	f := Francium{}
	want := "Francium"
	got := f.GetName()
	if got != want {
		t.Errorf("Francium.GetName() = got %v, want %v", got, want)
	}
}

func TestFranciumGetSimbol(t *testing.T) {
	f := Francium{}
	want := "Fr"
	got := f.GetSimbol()
	if got != want {
		t.Errorf("Francium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestFranciumGetAtomicNumber(t *testing.T) {
	f := Francium{}
	want := 87
	got := f.GetAtomicNumber()
	if got != want {
		t.Errorf("Francium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestFranciumGetAtomicWeight(t *testing.T) {
	f := Francium{}
	var want float32 = 223
	got := f.GetAtomicWeight()
	if got != want {
		t.Errorf("Francium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
