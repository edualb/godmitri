package element

import "testing"

func TestPlatinumGetPeriod(t *testing.T) {
	p := Platinum{}
	want := "6th period"
	got := p.GetPeriod()
	if got != want {
		t.Errorf("Platinum.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestPlatinumGetGroup(t *testing.T) {
	p := Platinum{}
	want := "8B"
	got := p.GetGroup()
	if got != want {
		t.Errorf("Platinum.GetGroup() = got %v, want %v", got, want)
	}
}

func TestPlatinumGetCategory(t *testing.T) {
	p := Platinum{}
	want := "Transition metal"
	got := p.GetCategory()
	if got != want {
		t.Errorf("Platinum.GetCategory() = got %v, want %v", got, want)
	}
}

func TestPlatinumGetName(t *testing.T) {
	p := Platinum{}
	want := "Platinum"
	got := p.GetName()
	if got != want {
		t.Errorf("Platinum.GetName() = got %v, want %v", got, want)
	}
}

func TestPlatinumGetSimbol(t *testing.T) {
	p := Platinum{}
	want := "Pt"
	got := p.GetSimbol()
	if got != want {
		t.Errorf("Platinum.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestPlatinumGetAtomicNumber(t *testing.T) {
	p := Platinum{}
	want := 78
	got := p.GetAtomicNumber()
	if got != want {
		t.Errorf("Platinum.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestPlatinumGetAtomicWeight(t *testing.T) {
	p := Platinum{}
	var want float32 = 195.08
	got := p.GetAtomicWeight()
	if got != want {
		t.Errorf("Platinum.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
