package element

import "testing"

func TestTinGetPeriod(t *testing.T) {
	tin := Tin{}
	want := "5th period"
	got := tin.GetPeriod()
	if got != want {
		t.Errorf("Tin.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTinGetGroup(t *testing.T) {
	tin := Tin{}
	want := "4A"
	got := tin.GetGroup()
	if got != want {
		t.Errorf("Tin.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTinGetCategory(t *testing.T) {
	tin := Tin{}
	want := "Post-transition metal"
	got := tin.GetCategory()
	if got != want {
		t.Errorf("Tin.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTinGetName(t *testing.T) {
	tin := Tin{}
	want := "Tin"
	got := tin.GetName()
	if got != want {
		t.Errorf("Tin.GetName() = got %v, want %v", got, want)
	}
}

func TestTinGetSimbol(t *testing.T) {
	tin := Tin{}
	want := "Sn"
	got := tin.GetSimbol()
	if got != want {
		t.Errorf("Tin.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTinGetAtomicNumber(t *testing.T) {
	tin := Tin{}
	want := 50
	got := tin.GetAtomicNumber()
	if got != want {
		t.Errorf("Tin.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTinGetAtomicWeight(t *testing.T) {
	tin := Tin{}
	var want float32 = 118.71
	got := tin.GetAtomicWeight()
	if got != want {
		t.Errorf("Tin.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
