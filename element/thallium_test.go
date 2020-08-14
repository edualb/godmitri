package element

import "testing"

func TestThalliumGetPeriod(t *testing.T) {
	th := Thallium{}
	want := "6th period"
	got := th.GetPeriod()
	if got != want {
		t.Errorf("Thallium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestThalliumGetGroup(t *testing.T) {
	th := Thallium{}
	want := "3A"
	got := th.GetGroup()
	if got != want {
		t.Errorf("Thallium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestThalliumGetCategory(t *testing.T) {
	th := Thallium{}
	want := "Post-transition metal"
	got := th.GetCategory()
	if got != want {
		t.Errorf("Thallium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestThalliumGetName(t *testing.T) {
	th := Thallium{}
	want := "Thallium"
	got := th.GetName()
	if got != want {
		t.Errorf("Thallium.GetName() = got %v, want %v", got, want)
	}
}

func TestThalliumGetSimbol(t *testing.T) {
	th := Thallium{}
	want := "Tl"
	got := th.GetSimbol()
	if got != want {
		t.Errorf("Thallium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestThalliumGetAtomicNumber(t *testing.T) {
	th := Thallium{}
	want := 81
	got := th.GetAtomicNumber()
	if got != want {
		t.Errorf("Thallium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestThalliumGetAtomicWeight(t *testing.T) {
	th := Thallium{}
	var want float32 = 204.38
	got := th.GetAtomicWeight()
	if got != want {
		t.Errorf("Thallium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
