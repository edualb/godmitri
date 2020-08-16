package element

import "testing"

func TestAmericiumGetPeriod(t *testing.T) {
	a := Americium{}
	want := "7th period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Americium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestAmericiumGetGroup(t *testing.T) {
	a := Americium{}
	want := "3B"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Americium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestAmericiumGetCategory(t *testing.T) {
	a := Americium{}
	want := "Actinoid"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Americium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestAmericiumGetName(t *testing.T) {
	a := Americium{}
	want := "Americium"
	got := a.GetName()
	if got != want {
		t.Errorf("Americium.GetName() = got %v, want %v", got, want)
	}
}

func TestAmericiumGetSimbol(t *testing.T) {
	a := Americium{}
	want := "Am"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Americium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestAmericiumGetAtomicNumber(t *testing.T) {
	a := Americium{}
	want := 95
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Americium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestAmericiumGetAtomicWeight(t *testing.T) {
	a := Americium{}
	var want float32 = 243
	got := a.GetAtomicWeight()
	if got != want {
		t.Errorf("Americium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
