package element

import "testing"

func TestIronGetPeriod(t *testing.T) {
	i := Iron{}
	want := "4th period"
	got := i.GetPeriod()
	if got != want {
		t.Errorf("Iron.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestIronGetGroup(t *testing.T) {
	i := Iron{}
	want := "8B"
	got := i.GetGroup()
	if got != want {
		t.Errorf("Iron.GetGroup() = got %v, want %v", got, want)
	}
}

func TestIronGetCategory(t *testing.T) {
	i := Iron{}
	want := "Transition metal"
	got := i.GetCategory()
	if got != want {
		t.Errorf("Iron.GetCategory() = got %v, want %v", got, want)
	}
}

func TestIronGetName(t *testing.T) {
	i := Iron{}
	want := "Iron"
	got := i.GetName()
	if got != want {
		t.Errorf("Iron.GetName() = got %v, want %v", got, want)
	}
}

func TestIronGetSimbol(t *testing.T) {
	i := Iron{}
	want := "Fe"
	got := i.GetSimbol()
	if got != want {
		t.Errorf("Iron.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestIronGetAtomicNumber(t *testing.T) {
	i := Iron{}
	want := 26
	got := i.GetAtomicNumber()
	if got != want {
		t.Errorf("Iron.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestIronGetAtomicWeight(t *testing.T) {
	i := Iron{}
	var want float32 = 55.845
	got := i.GetAtomicWeight()
	if got != want {
		t.Errorf("Iron.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
