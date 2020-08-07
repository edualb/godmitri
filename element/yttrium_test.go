package element

import "testing"

func TestYttriumGetPeriod(t *testing.T) {
	y := Yttrium{}
	want := "5th period"
	got := y.GetPeriod()
	if got != want {
		t.Errorf("Yttrium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestYttriumGetGroup(t *testing.T) {
	y := Yttrium{}
	want := "3B"
	got := y.GetGroup()
	if got != want {
		t.Errorf("Yttrium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestYttriumGetCategory(t *testing.T) {
	y := Yttrium{}
	want := "Transition metal"
	got := y.GetCategory()
	if got != want {
		t.Errorf("Yttrium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestYttriumGetName(t *testing.T) {
	y := Yttrium{}
	want := "Yttrium"
	got := y.GetName()
	if got != want {
		t.Errorf("Yttrium.GetName() = got %v, want %v", got, want)
	}
}

func TestYttriumGetSimbol(t *testing.T) {
	y := Yttrium{}
	want := "Y"
	got := y.GetSimbol()
	if got != want {
		t.Errorf("Yttrium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestYttriumGetAtomicNumber(t *testing.T) {
	y := Yttrium{}
	want := 39
	got := y.GetAtomicNumber()
	if got != want {
		t.Errorf("Yttrium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestYttriumGetAtomicWeight(t *testing.T) {
	y := Yttrium{}
	var want float32 = 88.906
	got := y.GetAtomicWeight()
	if got != want {
		t.Errorf("Yttrium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
