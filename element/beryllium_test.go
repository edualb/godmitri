package element

import "testing"

func TestBerylliumGetPeriod(t *testing.T) {
	b := Beryllium{}
	want := "2nd period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Beryllium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBerylliumGetGroup(t *testing.T) {
	b := Beryllium{}
	want := "2A"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Beryllium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBerylliumGetCategory(t *testing.T) {
	b := Beryllium{}
	want := "Alkaline earth metal"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Beryllium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBerylliumGetName(t *testing.T) {
	b := Beryllium{}
	want := "Beryllium"
	got := b.GetName()
	if got != want {
		t.Errorf("Beryllium.GetName() = got %v, want %v", got, want)
	}
}

func TestBerylliumGetSimbol(t *testing.T) {
	b := Beryllium{}
	want := "Be"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Beryllium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBerylliumGetAtomicNumber(t *testing.T) {
	b := Beryllium{}
	want := 4
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Beryllium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBerylliumGetAtomicWeight(t *testing.T) {
	b := Beryllium{}
	var want float32 = 9.0122
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Beryllium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
