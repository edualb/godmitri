package element

import "testing"

func TestIridiumGetPeriod(t *testing.T) {
	i := Iridium{}
	want := "6th period"
	got := i.GetPeriod()
	if got != want {
		t.Errorf("Iridium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestIridiumGetGroup(t *testing.T) {
	i := Iridium{}
	want := "8B"
	got := i.GetGroup()
	if got != want {
		t.Errorf("Iridium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestIridiumGetCategory(t *testing.T) {
	i := Iridium{}
	want := "Transition metal"
	got := i.GetCategory()
	if got != want {
		t.Errorf("Iridium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestIridiumGetName(t *testing.T) {
	i := Iridium{}
	want := "Iridium"
	got := i.GetName()
	if got != want {
		t.Errorf("Iridium.GetName() = got %v, want %v", got, want)
	}
}

func TestIridiumGetSimbol(t *testing.T) {
	i := Iridium{}
	want := "Ir"
	got := i.GetSimbol()
	if got != want {
		t.Errorf("Iridium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestIridiumGetAtomicNumber(t *testing.T) {
	i := Iridium{}
	want := 77
	got := i.GetAtomicNumber()
	if got != want {
		t.Errorf("Iridium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestIridiumGetAtomicWeight(t *testing.T) {
	i := Iridium{}
	var want float32 = 192.22
	got := i.GetAtomicWeight()
	if got != want {
		t.Errorf("Iridium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
