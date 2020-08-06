package element

import "testing"

func TestVanadiumGetPeriod(t *testing.T) {
	v := Vanadium{}
	want := "4th period"
	got := v.GetPeriod()
	if got != want {
		t.Errorf("Vanadium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestVanadiumGetGroup(t *testing.T) {
	v := Vanadium{}
	want := "5B"
	got := v.GetGroup()
	if got != want {
		t.Errorf("Vanadium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestVanadiumGetCategory(t *testing.T) {
	v := Vanadium{}
	want := "Transition metal"
	got := v.GetCategory()
	if got != want {
		t.Errorf("Vanadium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestVanadiumGetName(t *testing.T) {
	v := Vanadium{}
	want := "Vanadium"
	got := v.GetName()
	if got != want {
		t.Errorf("Vanadium.GetName() = got %v, want %v", got, want)
	}
}

func TestVanadiumGetSimbol(t *testing.T) {
	v := Vanadium{}
	want := "V"
	got := v.GetSimbol()
	if got != want {
		t.Errorf("Vanadium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestVanadiumGetAtomicNumber(t *testing.T) {
	v := Vanadium{}
	want := 23
	got := v.GetAtomicNumber()
	if got != want {
		t.Errorf("Vanadium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestVanadiumGetAtomicWeight(t *testing.T) {
	v := Vanadium{}
	var want float32 = 50.942
	got := v.GetAtomicWeight()
	if got != want {
		t.Errorf("Vanadium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
