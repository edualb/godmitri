package element

import "testing"

func TestNihoniumGetPeriod(t *testing.T) {
	n := Nihonium{}
	want := "7th period"
	got := n.GetPeriod()
	if got != want {
		t.Errorf("Nihonium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestNihoniumGetGroup(t *testing.T) {
	n := Nihonium{}
	want := "3A"
	got := n.GetGroup()
	if got != want {
		t.Errorf("Nihonium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestNihoniumGetCategory(t *testing.T) {
	n := Nihonium{}
	want := "Unknown"
	got := n.GetCategory()
	if got != want {
		t.Errorf("Nihonium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestNihoniumGetName(t *testing.T) {
	n := Nihonium{}
	want := "Nihonium"
	got := n.GetName()
	if got != want {
		t.Errorf("Nihonium.GetName() = got %v, want %v", got, want)
	}
}

func TestNihoniumGetSimbol(t *testing.T) {
	n := Nihonium{}
	want := "Nh"
	got := n.GetSimbol()
	if got != want {
		t.Errorf("Nihonium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestNihoniumGetAtomicNumber(t *testing.T) {
	n := Nihonium{}
	want := 113
	got := n.GetAtomicNumber()
	if got != want {
		t.Errorf("Nihonium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestNihoniumGetAtomicWeight(t *testing.T) {
	n := Nihonium{}
	var want float32 = 286
	got := n.GetAtomicWeight()
	if got != want {
		t.Errorf("Nihonium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
