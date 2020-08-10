package element

import "testing"

func TestDysprosiumGetPeriod(t *testing.T) {
	d := Dysprosium{}
	want := "6th period"
	got := d.GetPeriod()
	if got != want {
		t.Errorf("Dysprosium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestDysprosiumGetGroup(t *testing.T) {
	d := Dysprosium{}
	want := "3B"
	got := d.GetGroup()
	if got != want {
		t.Errorf("Dysprosium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestDysprosiumGetCategory(t *testing.T) {
	d := Dysprosium{}
	want := "Lanthanoid"
	got := d.GetCategory()
	if got != want {
		t.Errorf("Dysprosium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestDysprosiumGetName(t *testing.T) {
	d := Dysprosium{}
	want := "Dysprosium"
	got := d.GetName()
	if got != want {
		t.Errorf("Dysprosium.GetName() = got %v, want %v", got, want)
	}
}

func TestDysprosiumGetSimbol(t *testing.T) {
	d := Dysprosium{}
	want := "Dy"
	got := d.GetSimbol()
	if got != want {
		t.Errorf("Dysprosium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestDysprosiumGetAtomicNumber(t *testing.T) {
	d := Dysprosium{}
	want := 66
	got := d.GetAtomicNumber()
	if got != want {
		t.Errorf("Dysprosium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestDysprosiumGetAtomicWeight(t *testing.T) {
	d := Dysprosium{}
	var want float32 = 157.25
	got := d.GetAtomicWeight()
	if got != want {
		t.Errorf("Dysprosium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
