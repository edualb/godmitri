package element

import "testing"

func TestAstatineGetPeriod(t *testing.T) {
	a := Astatine{}
	want := "6th period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Astatine.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestAstatineGetGroup(t *testing.T) {
	a := Astatine{}
	want := "7A"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Astatine.GetGroup() = got %v, want %v", got, want)
	}
}

func TestAstatineGetCategory(t *testing.T) {
	a := Astatine{}
	want := "Metalloid"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Astatine.GetCategory() = got %v, want %v", got, want)
	}
}

func TestAstatineGetName(t *testing.T) {
	a := Astatine{}
	want := "Astatine"
	got := a.GetName()
	if got != want {
		t.Errorf("Astatine.GetName() = got %v, want %v", got, want)
	}
}

func TestAstatineGetSimbol(t *testing.T) {
	a := Astatine{}
	want := "At"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Astatine.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestAstatineGetAtomicNumber(t *testing.T) {
	a := Astatine{}
	want := 85
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Astatine.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestAstatineGetAtomicWeight(t *testing.T) {
	a := Astatine{}
	var want float32 = 210
	got := a.GetAtomicWeight()
	if got != want {
		t.Errorf("Astatine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
