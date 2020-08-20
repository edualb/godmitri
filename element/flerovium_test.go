package element

import "testing"

func TestFleroviumGetPeriod(t *testing.T) {
	f := Flerovium{}
	want := "7th period"
	got := f.GetPeriod()
	if got != want {
		t.Errorf("Flerovium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestFleroviumGetGroup(t *testing.T) {
	f := Flerovium{}
	want := "4A"
	got := f.GetGroup()
	if got != want {
		t.Errorf("Flerovium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestFleroviumGetCategory(t *testing.T) {
	f := Flerovium{}
	want := "Post-transition metal"
	got := f.GetCategory()
	if got != want {
		t.Errorf("Flerovium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestFleroviumGetName(t *testing.T) {
	f := Flerovium{}
	want := "Flerovium"
	got := f.GetName()
	if got != want {
		t.Errorf("Flerovium.GetName() = got %v, want %v", got, want)
	}
}

func TestFleroviumGetSimbol(t *testing.T) {
	f := Flerovium{}
	want := "Fl"
	got := f.GetSimbol()
	if got != want {
		t.Errorf("Flerovium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestFleroviumGetAtomicNumber(t *testing.T) {
	f := Flerovium{}
	want := 114
	got := f.GetAtomicNumber()
	if got != want {
		t.Errorf("Flerovium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestFleroviumGetAtomicWeight(t *testing.T) {
	f := Flerovium{}
	var want float32 = 289
	got := f.GetAtomicWeight()
	if got != want {
		t.Errorf("Flerovium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
