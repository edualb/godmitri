package element

import "testing"

func TestXenonGetPeriod(t *testing.T) {
	x := Xenon{}
	want := "5th period"
	got := x.GetPeriod()
	if got != want {
		t.Errorf("Xenon.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestXenonGetGroup(t *testing.T) {
	x := Xenon{}
	want := "8A"
	got := x.GetGroup()
	if got != want {
		t.Errorf("Xenon.GetGroup() = got %v, want %v", got, want)
	}
}

func TestXenonGetCategory(t *testing.T) {
	x := Xenon{}
	want := "Noble gas"
	got := x.GetCategory()
	if got != want {
		t.Errorf("Xenon.GetCategory() = got %v, want %v", got, want)
	}
}

func TestXenonGetName(t *testing.T) {
	x := Xenon{}
	want := "Xenon"
	got := x.GetName()
	if got != want {
		t.Errorf("Xenon.GetName() = got %v, want %v", got, want)
	}
}

func TestXenonGetSimbol(t *testing.T) {
	x := Xenon{}
	want := "Xe"
	got := x.GetSimbol()
	if got != want {
		t.Errorf("Xenon.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestXenonGetAtomicNumber(t *testing.T) {
	x := Xenon{}
	want := 54
	got := x.GetAtomicNumber()
	if got != want {
		t.Errorf("Xenon.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestXenonGetAtomicWeight(t *testing.T) {
	x := Xenon{}
	var want float32 = 131.29
	got := x.GetAtomicWeight()
	if got != want {
		t.Errorf("Xenon.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
