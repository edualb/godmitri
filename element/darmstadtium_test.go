package element

import "testing"

func TestDarmstadtiumGetPeriod(t *testing.T) {
	d := Darmstadtium{}
	want := "7th period"
	got := d.GetPeriod()
	if got != want {
		t.Errorf("Darmstadtium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestDarmstadtiumGetGroup(t *testing.T) {
	d := Darmstadtium{}
	want := "8B"
	got := d.GetGroup()
	if got != want {
		t.Errorf("Darmstadtium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestDarmstadtiumGetCategory(t *testing.T) {
	d := Darmstadtium{}
	want := "Unknown"
	got := d.GetCategory()
	if got != want {
		t.Errorf("Darmstadtium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestDarmstadtiumGetName(t *testing.T) {
	d := Darmstadtium{}
	want := "Darmstadtium"
	got := d.GetName()
	if got != want {
		t.Errorf("Darmstadtium.GetName() = got %v, want %v", got, want)
	}
}

func TestDarmstadtiumGetSimbol(t *testing.T) {
	d := Darmstadtium{}
	want := "Ds"
	got := d.GetSimbol()
	if got != want {
		t.Errorf("Darmstadtium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestDarmstadtiumGetAtomicNumber(t *testing.T) {
	d := Darmstadtium{}
	want := 110
	got := d.GetAtomicNumber()
	if got != want {
		t.Errorf("Darmstadtium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestDarmstadtiumGetAtomicWeight(t *testing.T) {
	d := Darmstadtium{}
	var want float32 = 281
	got := d.GetAtomicWeight()
	if got != want {
		t.Errorf("Darmstadtium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
