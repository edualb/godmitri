package element

import "testing"

func TestKryptonGetPeriod(t *testing.T) {
	k := Krypton{}
	want := "4th period"
	got := k.GetPeriod()
	if got != want {
		t.Errorf("Krypton.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestKryptonGetGroup(t *testing.T) {
	k := Krypton{}
	want := "8A"
	got := k.GetGroup()
	if got != want {
		t.Errorf("Krypton.GetGroup() = got %v, want %v", got, want)
	}
}

func TestKryptonGetCategory(t *testing.T) {
	k := Krypton{}
	want := "Noble gas"
	got := k.GetCategory()
	if got != want {
		t.Errorf("Krypton.GetCategory() = got %v, want %v", got, want)
	}
}

func TestKryptonGetName(t *testing.T) {
	k := Krypton{}
	want := "Krypton"
	got := k.GetName()
	if got != want {
		t.Errorf("Krypton.GetName() = got %v, want %v", got, want)
	}
}

func TestKryptonGetSimbol(t *testing.T) {
	k := Krypton{}
	want := "Kr"
	got := k.GetSimbol()
	if got != want {
		t.Errorf("Krypton.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestKryptonGetAtomicNumber(t *testing.T) {
	k := Krypton{}
	want := 36
	got := k.GetAtomicNumber()
	if got != want {
		t.Errorf("Krypton.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestKryptonGetAtomicWeight(t *testing.T) {
	k := Krypton{}
	var want float32 = 83.798
	got := k.GetAtomicWeight()
	if got != want {
		t.Errorf("Krypton.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
