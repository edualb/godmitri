package element

import "testing"

func TestOsmiumGetPeriod(t *testing.T) {
	o := Osmium{}
	want := "6th period"
	got := o.GetPeriod()
	if got != want {
		t.Errorf("Osmium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestOsmiumGetGroup(t *testing.T) {
	o := Osmium{}
	want := "8B"
	got := o.GetGroup()
	if got != want {
		t.Errorf("Osmium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestOsmiumGetCategory(t *testing.T) {
	o := Osmium{}
	want := "Transition metal"
	got := o.GetCategory()
	if got != want {
		t.Errorf("Osmium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestOsmiumGetName(t *testing.T) {
	o := Osmium{}
	want := "Osmium"
	got := o.GetName()
	if got != want {
		t.Errorf("Osmium.GetName() = got %v, want %v", got, want)
	}
}

func TestOsmiumGetSimbol(t *testing.T) {
	o := Osmium{}
	want := "Os"
	got := o.GetSimbol()
	if got != want {
		t.Errorf("Osmium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestOsmiumGetAtomicNumber(t *testing.T) {
	o := Osmium{}
	want := 76
	got := o.GetAtomicNumber()
	if got != want {
		t.Errorf("Osmium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestOsmiumGetAtomicWeight(t *testing.T) {
	o := Osmium{}
	var want float32 = 190.23
	got := o.GetAtomicWeight()
	if got != want {
		t.Errorf("Osmium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
