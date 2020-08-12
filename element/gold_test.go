package element

import "testing"

func TestGoldGetPeriod(t *testing.T) {
	g := Gold{}
	want := "6th period"
	got := g.GetPeriod()
	if got != want {
		t.Errorf("Gold.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestGoldGetGroup(t *testing.T) {
	g := Gold{}
	want := "1B"
	got := g.GetGroup()
	if got != want {
		t.Errorf("Gold.GetGroup() = got %v, want %v", got, want)
	}
}

func TestGoldGetCategory(t *testing.T) {
	g := Gold{}
	want := "Transition metal"
	got := g.GetCategory()
	if got != want {
		t.Errorf("Gold.GetCategory() = got %v, want %v", got, want)
	}
}

func TestGoldGetName(t *testing.T) {
	g := Gold{}
	want := "Gold"
	got := g.GetName()
	if got != want {
		t.Errorf("Gold.GetName() = got %v, want %v", got, want)
	}
}

func TestGoldGetSimbol(t *testing.T) {
	g := Gold{}
	want := "Au"
	got := g.GetSimbol()
	if got != want {
		t.Errorf("Gold.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestGoldGetAtomicNumber(t *testing.T) {
	g := Gold{}
	want := 79
	got := g.GetAtomicNumber()
	if got != want {
		t.Errorf("Gold.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestGoldGetAtomicWeight(t *testing.T) {
	g := Gold{}
	var want float32 = 196.97
	got := g.GetAtomicWeight()
	if got != want {
		t.Errorf("Gold.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
