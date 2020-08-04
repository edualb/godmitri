package element

import "testing"

func TestOxygenGetPeriod(t *testing.T) {
	o := Oxygen{}
	want := "2nd period"
	got := o.GetPeriod()
	if got != want {
		t.Errorf("Oxygen.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestOxygenGetGroup(t *testing.T) {
	o := Oxygen{}
	want := "6A"
	got := o.GetGroup()
	if got != want {
		t.Errorf("Oxygen.GetGroup() = got %v, want %v", got, want)
	}
}

func TestOxygenGetCategory(t *testing.T) {
	o := Oxygen{}
	want := "Non-metal"
	got := o.GetCategory()
	if got != want {
		t.Errorf("Oxygen.GetCategory() = got %v, want %v", got, want)
	}
}

func TestOxygenGetName(t *testing.T) {
	o := Oxygen{}
	want := "Oxygen"
	got := o.GetName()
	if got != want {
		t.Errorf("Oxygen.GetName() = got %v, want %v", got, want)
	}
}

func TestOxygenGetSimbol(t *testing.T) {
	o := Oxygen{}
	want := "O"
	got := o.GetSimbol()
	if got != want {
		t.Errorf("Oxygen.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestOxygenGetAtomicNumber(t *testing.T) {
	o := Oxygen{}
	want := 8
	got := o.GetAtomicNumber()
	if got != want {
		t.Errorf("Oxygen.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestOxygenGetAtomicWeight(t *testing.T) {
	o := Oxygen{}
	var want float32 = 15.999
	got := o.GetAtomicWeight()
	if got != want {
		t.Errorf("Oxygen.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
