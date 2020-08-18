package element

import "testing"

func TestBerkeliumGetPeriod(t *testing.T) {
	b := Berkelium{}
	want := "7th period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Berkelium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBerkeliumGetGroup(t *testing.T) {
	b := Berkelium{}
	want := "3B"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Berkelium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBerkeliumGetCategory(t *testing.T) {
	b := Berkelium{}
	want := "Actinoid"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Berkelium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBerkeliumGetName(t *testing.T) {
	b := Berkelium{}
	want := "Berkelium"
	got := b.GetName()
	if got != want {
		t.Errorf("Berkelium.GetName() = got %v, want %v", got, want)
	}
}

func TestBerkeliumGetSimbol(t *testing.T) {
	b := Berkelium{}
	want := "Bk"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Berkelium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBerkeliumGetAtomicNumber(t *testing.T) {
	b := Berkelium{}
	want := 97
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Berkelium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBerkeliumGetAtomicWeight(t *testing.T) {
	b := Berkelium{}
	var want float32 = 247
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Berkelium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
