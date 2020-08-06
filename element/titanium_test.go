package element

import "testing"

func TestTitaniumGetPeriod(t *testing.T) {
	tt := Titanium{}
	want := "4th period"
	got := tt.GetPeriod()
	if got != want {
		t.Errorf("Titanium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestTitaniumGetGroup(t *testing.T) {
	tt := Titanium{}
	want := "4B"
	got := tt.GetGroup()
	if got != want {
		t.Errorf("Titanium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestTitaniumGetCategory(t *testing.T) {
	tt := Titanium{}
	want := "Transition metal"
	got := tt.GetCategory()
	if got != want {
		t.Errorf("Titanium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestTitaniumGetName(t *testing.T) {
	tt := Titanium{}
	want := "Titanium"
	got := tt.GetName()
	if got != want {
		t.Errorf("Titanium.GetName() = got %v, want %v", got, want)
	}
}

func TestTitaniumGetSimbol(t *testing.T) {
	tt := Titanium{}
	want := "Ti"
	got := tt.GetSimbol()
	if got != want {
		t.Errorf("Titanium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestTitaniumGetAtomicNumber(t *testing.T) {
	tt := Titanium{}
	want := 22
	got := tt.GetAtomicNumber()
	if got != want {
		t.Errorf("Titanium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestTitaniumGetAtomicWeight(t *testing.T) {
	tt := Titanium{}
	var want float32 = 47.867
	got := tt.GetAtomicWeight()
	if got != want {
		t.Errorf("Titanium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
