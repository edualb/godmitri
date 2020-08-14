package element

import "testing"

func TestBismuthGetPeriod(t *testing.T) {
	b := Bismuth{}
	want := "6th period"
	got := b.GetPeriod()
	if got != want {
		t.Errorf("Bismuth.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestBismuthGetGroup(t *testing.T) {
	b := Bismuth{}
	want := "5A"
	got := b.GetGroup()
	if got != want {
		t.Errorf("Bismuth.GetGroup() = got %v, want %v", got, want)
	}
}

func TestBismuthGetCategory(t *testing.T) {
	b := Bismuth{}
	want := "Post-transition metal"
	got := b.GetCategory()
	if got != want {
		t.Errorf("Bismuth.GetCategory() = got %v, want %v", got, want)
	}
}

func TestBismuthGetName(t *testing.T) {
	b := Bismuth{}
	want := "Bismuth"
	got := b.GetName()
	if got != want {
		t.Errorf("Bismuth.GetName() = got %v, want %v", got, want)
	}
}

func TestBismuthGetSimbol(t *testing.T) {
	b := Bismuth{}
	want := "Bi"
	got := b.GetSimbol()
	if got != want {
		t.Errorf("Bismuth.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestBismuthGetAtomicNumber(t *testing.T) {
	b := Bismuth{}
	want := 83
	got := b.GetAtomicNumber()
	if got != want {
		t.Errorf("Bismuth.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestBismuthGetAtomicWeight(t *testing.T) {
	b := Bismuth{}
	var want float32 = 208.98
	got := b.GetAtomicWeight()
	if got != want {
		t.Errorf("Bismuth.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
