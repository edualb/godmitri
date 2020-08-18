package element

import "testing"

func TestDubniumGetPeriod(t *testing.T) {
	d := Dubnium{}
	want := "7th period"
	got := d.GetPeriod()
	if got != want {
		t.Errorf("Dubnium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestDubniumGetGroup(t *testing.T) {
	d := Dubnium{}
	want := "5B"
	got := d.GetGroup()
	if got != want {
		t.Errorf("Dubnium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestDubniumGetCategory(t *testing.T) {
	d := Dubnium{}
	want := "Transition metal"
	got := d.GetCategory()
	if got != want {
		t.Errorf("Dubnium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestDubniumGetName(t *testing.T) {
	d := Dubnium{}
	want := "Dubnium"
	got := d.GetName()
	if got != want {
		t.Errorf("Dubnium.GetName() = got %v, want %v", got, want)
	}
}

func TestDubniumGetSimbol(t *testing.T) {
	d := Dubnium{}
	want := "Db"
	got := d.GetSimbol()
	if got != want {
		t.Errorf("Dubnium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestDubniumGetAtomicNumber(t *testing.T) {
	d := Dubnium{}
	want := 105
	got := d.GetAtomicNumber()
	if got != want {
		t.Errorf("Dubnium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestDubniumGetAtomicWeight(t *testing.T) {
	d := Dubnium{}
	var want float32 = 268
	got := d.GetAtomicWeight()
	if got != want {
		t.Errorf("Dubnium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
