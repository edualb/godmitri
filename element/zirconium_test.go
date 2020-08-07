package element

import "testing"

func TestZirconiumGetPeriod(t *testing.T) {
	z := Zirconium{}
	want := "5th period"
	got := z.GetPeriod()
	if got != want {
		t.Errorf("Zirconium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestZirconiumGetGroup(t *testing.T) {
	z := Zirconium{}
	want := "4B"
	got := z.GetGroup()
	if got != want {
		t.Errorf("Zirconium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestZirconiumGetCategory(t *testing.T) {
	z := Zirconium{}
	want := "Transition metal"
	got := z.GetCategory()
	if got != want {
		t.Errorf("Zirconium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestZirconiumGetName(t *testing.T) {
	z := Zirconium{}
	want := "Zirconium"
	got := z.GetName()
	if got != want {
		t.Errorf("Zirconium.GetName() = got %v, want %v", got, want)
	}
}

func TestZirconiumGetSimbol(t *testing.T) {
	z := Zirconium{}
	want := "Zr"
	got := z.GetSimbol()
	if got != want {
		t.Errorf("Zirconium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestZirconiumGetAtomicNumber(t *testing.T) {
	z := Zirconium{}
	want := 40
	got := z.GetAtomicNumber()
	if got != want {
		t.Errorf("Zirconium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestZirconiumGetAtomicWeight(t *testing.T) {
	z := Zirconium{}
	var want float32 = 91.224
	got := z.GetAtomicWeight()
	if got != want {
		t.Errorf("Zirconium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
