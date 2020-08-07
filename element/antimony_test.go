package element

import "testing"

func TestAntimonyGetPeriod(t *testing.T) {
	a := Antimony{}
	want := "5th period"
	got := a.GetPeriod()
	if got != want {
		t.Errorf("Antimony.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestAntimonyGetGroup(t *testing.T) {
	a := Antimony{}
	want := "5A"
	got := a.GetGroup()
	if got != want {
		t.Errorf("Antimony.GetGroup() = got %v, want %v", got, want)
	}
}

func TestAntimonyGetCategory(t *testing.T) {
	a := Antimony{}
	want := "Metalloid"
	got := a.GetCategory()
	if got != want {
		t.Errorf("Antimony.GetCategory() = got %v, want %v", got, want)
	}
}

func TestAntimonyGetName(t *testing.T) {
	a := Antimony{}
	want := "Antimony"
	got := a.GetName()
	if got != want {
		t.Errorf("Antimony.GetName() = got %v, want %v", got, want)
	}
}

func TestAntimonyGetSimbol(t *testing.T) {
	a := Antimony{}
	want := "Sb"
	got := a.GetSimbol()
	if got != want {
		t.Errorf("Antimony.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestAntimonyGetAtomicNumber(t *testing.T) {
	a := Antimony{}
	want := 51
	got := a.GetAtomicNumber()
	if got != want {
		t.Errorf("Antimony.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestAntimonyGetAtomicWeight(t *testing.T) {
	a := Antimony{}
	var want float32 = 121.76
	got := a.GetAtomicWeight()
	if got != want {
		t.Errorf("Antimony.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
