package element

import "testing"

func TestTelluriumGetPeriod(t *testing.T) {
	te := Tellurium{}
	want := "5th period"
	got := te.GetPeriod()
	if got != want {
		t.Errorf("Tellurium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTelluriumGetGroup(t *testing.T) {
	te := Tellurium{}
	want := "6A"
	got := te.GetGroup()
	if got != want {
		t.Errorf("Tellurium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTelluriumGetCategory(t *testing.T) {
	te := Tellurium{}
	want := "Metalloid"
	got := te.GetCategory()
	if got != want {
		t.Errorf("Tellurium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTelluriumGetName(t *testing.T) {
	te := Tellurium{}
	want := "Tellurium"
	got := te.GetName()
	if got != want {
		t.Errorf("Tellurium.GetName() = got %v, want %v", got, want)
	}
}

func TestTelluriumGetSimbol(t *testing.T) {
	te := Tellurium{}
	want := "Te"
	got := te.GetSimbol()
	if got != want {
		t.Errorf("Tellurium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTelluriumGetAtomicNumber(t *testing.T) {
	te := Tellurium{}
	want := 52
	got := te.GetAtomicNumber()
	if got != want {
		t.Errorf("Tellurium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTelluriumGetAtomicWeight(t *testing.T) {
	te := Tellurium{}
	var want float32 = 127.60
	got := te.GetAtomicWeight()
	if got != want {
		t.Errorf("Tellurium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
