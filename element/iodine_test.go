package element

import "testing"

func TestIodineGetPeriod(t *testing.T) {
	i := Iodine{}
	want := "5th period"
	got := i.GetPeriod()
	if got != want {
		t.Errorf("Iodine.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestIodineGetGroup(t *testing.T) {
	i := Iodine{}
	want := "7A"
	got := i.GetGroup()
	if got != want {
		t.Errorf("Iodine.GetGroup() = got %v, want %v", got, want)
	}
}

func TestIodineGetCategory(t *testing.T) {
	i := Iodine{}
	want := "Non-metal"
	got := i.GetCategory()
	if got != want {
		t.Errorf("Iodine.GetCategory() = got %v, want %v", got, want)
	}
}

func TestIodineGetName(t *testing.T) {
	i := Iodine{}
	want := "Iodine"
	got := i.GetName()
	if got != want {
		t.Errorf("Iodine.GetName() = got %v, want %v", got, want)
	}
}

func TestIodineGetSimbol(t *testing.T) {
	i := Iodine{}
	want := "I"
	got := i.GetSimbol()
	if got != want {
		t.Errorf("Iodine.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestIodineGetAtomicNumber(t *testing.T) {
	i := Iodine{}
	want := 53
	got := i.GetAtomicNumber()
	if got != want {
		t.Errorf("Iodine.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestIodineGetAtomicWeight(t *testing.T) {
	i := Iodine{}
	var want float32 = 126.90
	got := i.GetAtomicWeight()
	if got != want {
		t.Errorf("Iodine.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
