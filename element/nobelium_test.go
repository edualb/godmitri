package element

import "testing"

func TestNobeliumGetPeriod(t *testing.T) {
	n := Nobelium{}
	want := "7th period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Nobelium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNobeliumGetGroup(t *testing.T) {
	n := Nobelium{}
	want := "3B"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Nobelium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNobeliumGetCategory(t *testing.T) {
	n := Nobelium{}
	want := "Actinoid"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Nobelium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNobeliumGetName(t *testing.T) {
	n := Nobelium{}
	want := "Nobelium"
	got := n.GetName()
	if got != want {
		t.Errorf("Nobelium.GetName() = got %v, want %v", got, want)
	}
}

func TestNobeliumGetSimbol(t *testing.T) {
	n := Nobelium{}
	want := "No"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Nobelium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNobeliumGetAtomicNumber(t *testing.T) {
	n := Nobelium{}
	want := 102
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Nobelium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNobeliumGetAtomicWeight(t *testing.T) {
	n := Nobelium{}
	var want float32 = 259
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Nobelium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
