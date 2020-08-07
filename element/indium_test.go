package element

import "testing"

func TestIndiumGetPeriod(t *testing.T) {
	i := Indium{}
	want := "5th period"
	got := i.GetPeriod()
	if got != want {
		t.Errorf("Indium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestIndiumGetGroup(t *testing.T) {
	i := Indium{}
	want := "3A"
	got := i.GetGroup()
	if got != want {
		t.Errorf("Indium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestIndiumGetCategory(t *testing.T) {
	i := Indium{}
	want := "Post-transition metal"
	got := i.GetCategory()
	if got != want {
		t.Errorf("Indium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestIndiumGetName(t *testing.T) {
	i := Indium{}
	want := "Indium"
	got := i.GetName()
	if got != want {
		t.Errorf("Indium.GetName() = got %v, want %v", got, want)
	}
}

func TestIndiumGetSimbol(t *testing.T) {
	i := Indium{}
	want := "In"
	got := i.GetSimbol()
	if got != want {
		t.Errorf("Indium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestIndiumGetAtomicNumber(t *testing.T) {
	i := Indium{}
	want := 49
	got := i.GetAtomicNumber()
	if got != want {
		t.Errorf("Indium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestIndiumGetAtomicWeight(t *testing.T) {
	i := Indium{}
	var want float32 = 114.82
	got := i.GetAtomicWeight()
	if got != want {
		t.Errorf("Indium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
