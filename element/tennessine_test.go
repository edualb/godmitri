package element

import "testing"

func TestTennessineGetPeriod(t *testing.T) {
	te := Tennessine{}
	want := "7th period"
	got := te.GetPeriod()
	if got != want {
		t.Errorf("Tennessine.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTennessineGetGroup(t *testing.T) {
	te := Tennessine{}
	want := "7A"
	got := te.GetGroup()
	if got != want {
		t.Errorf("Tennessine.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTennessineGetCategory(t *testing.T) {
	te := Tennessine{}
	want := "Unknown"
	got := te.GetCategory()
	if got != want {
		t.Errorf("Tennessine.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTennessineGetName(t *testing.T) {
	te := Tennessine{}
	want := "Tennessine"
	got := te.GetName()
	if got != want {
		t.Errorf("Tennessine.GetName() = got %v, want %v", got, want)
	}
}

func TestTennessineGetSimbol(t *testing.T) {
	te := Tennessine{}
	want := "Ts"
	got := te.GetSimbol()
	if got != want {
		t.Errorf("Tennessine.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTennessineGetAtomicNumber(t *testing.T) {
	te := Tennessine{}
	want := 117
	got := te.GetAtomicNumber()
	if got != want {
		t.Errorf("Tennessine.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTennessineGetAtomicWeight(t *testing.T) {
	te := Tennessine{}
	var want float32 = 294
	got := te.GetAtomicWeight()
	if got != want {
		t.Errorf("Tennessine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
