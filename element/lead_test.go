package element

import "testing"

func TestLeadGetPeriod(t *testing.T) {
	l := Lead{}
	want := "6th period"
	got := l.GetPeriod()
	if got != want {
		t.Errorf("Lead.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestLeadGetGroup(t *testing.T) {
	l := Lead{}
	want := "4A"
	got := l.GetGroup()
	if got != want {
		t.Errorf("Lead.GetGroup() = got %v, want %v", got, want)
	}
}

func TestLeadGetCategory(t *testing.T) {
	l := Lead{}
	want := "Post-transition metal"
	got := l.GetCategory()
	if got != want {
		t.Errorf("Lead.GetCategory() = got %v, want %v", got, want)
	}
}

func TestLeadGetName(t *testing.T) {
	l := Lead{}
	want := "Lead"
	got := l.GetName()
	if got != want {
		t.Errorf("Lead.GetName() = got %v, want %v", got, want)
	}
}

func TestLeadGetSimbol(t *testing.T) {
	l := Lead{}
	want := "Pb"
	got := l.GetSimbol()
	if got != want {
		t.Errorf("Lead.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestLeadGetAtomicNumber(t *testing.T) {
	l := Lead{}
	want := 82
	got := l.GetAtomicNumber()
	if got != want {
		t.Errorf("Lead.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestLeadGetAtomicWeight(t *testing.T) {
	l := Lead{}
	var want float32 = 207.2
	got := l.GetAtomicWeight()
	if got != want {
		t.Errorf("Lead.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
