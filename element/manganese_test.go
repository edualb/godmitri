package element

import "testing"

func TestManganeseGetPeriod(t *testing.T) {
	m := Manganese{}
	want := "4th period"
	got := m.GetPeriod()
	if got != want {
		t.Errorf("Manganese.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestManganeseGetGroup(t *testing.T) {
	m := Manganese{}
	want := "7B"
	got := m.GetGroup()
	if got != want {
		t.Errorf("Manganese.GetGroup() = got %v, want %v", got, want)
	}
}

func TestManganeseGetCategory(t *testing.T) {
	m := Manganese{}
	want := "Transition metal"
	got := m.GetCategory()
	if got != want {
		t.Errorf("Manganese.GetCategory() = got %v, want %v", got, want)
	}
}

func TestManganeseGetName(t *testing.T) {
	m := Manganese{}
	want := "Manganese"
	got := m.GetName()
	if got != want {
		t.Errorf("Manganese.GetName() = got %v, want %v", got, want)
	}
}

func TestManganeseGetSimbol(t *testing.T) {
	m := Manganese{}
	want := "Mn"
	got := m.GetSimbol()
	if got != want {
		t.Errorf("Manganese.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestManganeseGetAtomicNumber(t *testing.T) {
	m := Manganese{}
	want := 25
	got := m.GetAtomicNumber()
	if got != want {
		t.Errorf("Manganese.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestManganeseGetAtomicWeight(t *testing.T) {
	m := Manganese{}
	var want float32 = 54.938
	got := m.GetAtomicWeight()
	if got != want {
		t.Errorf("Manganese.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
