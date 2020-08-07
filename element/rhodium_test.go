package element

import "testing"

func TestRhodiumGetPeriod(t *testing.T) {
	r := Rhodium{}
	want := "5th period"
	got := r.GetPeriod()
	if got != want {
		t.Errorf("Rhodium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestRhodiumGetGroup(t *testing.T) {
	r := Rhodium{}
	want := "8B"
	got := r.GetGroup()
	if got != want {
		t.Errorf("Rhodium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestRhodiumGetCategory(t *testing.T) {
	r := Rhodium{}
	want := "Transition metal"
	got := r.GetCategory()
	if got != want {
		t.Errorf("Rhodium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestRhodiumGetName(t *testing.T) {
	r := Rhodium{}
	want := "Rhodium"
	got := r.GetName()
	if got != want {
		t.Errorf("Rhodium.GetName() = got %v, want %v", got, want)
	}
}

func TestRhodiumGetSimbol(t *testing.T) {
	r := Rhodium{}
	want := "Rh"
	got := r.GetSimbol()
	if got != want {
		t.Errorf("Rhodium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestRhodiumGetAtomicNumber(t *testing.T) {
	r := Rhodium{}
	want := 45
	got := r.GetAtomicNumber()
	if got != want {
		t.Errorf("Rhodium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestRhodiumGetAtomicWeight(t *testing.T) {
	r := Rhodium{}
	var want float32 = 102.91
	got := r.GetAtomicWeight()
	if got != want {
		t.Errorf("Rhodium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
