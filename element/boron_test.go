package element

import "testing"

func TestBoronGetPeriod(t *testing.T) {
	b := Boron{}
	want := "2nd period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Boron.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBoronGetGroup(t *testing.T) {
	b := Boron{}
	want := "3A"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Boron.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBoronGetCategory(t *testing.T) {
	b := Boron{}
	want := "Metalloid"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Boron.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBoronGetName(t *testing.T) {
	b := Boron{}
	want := "Boron"
	got := b.GetName()
	if got != want {
		t.Errorf("Boron.GetName() = got %v, want %v", got, want)
	}
}

func TestBoronGetSimbol(t *testing.T) {
	b := Boron{}
	want := "B"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Boron.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBoronGetAtomicNumber(t *testing.T) {
	b := Boron{}
	want := 5
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Boron.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBoronGetAtomicWeight(t *testing.T) {
	b := Boron{}
	var want float32 = 10.81
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Boron.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
