package element

import "testing"

func TestLivermoriumGetPeriod(t *testing.T) {
	l := Livermorium{}
	want := "7th period"
	got := l.GetPeriod()
	if got != want {
		t.Errorf("Livermorium.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestLivermoriumGetGroup(t *testing.T) {
	l := Livermorium{}
	want := "6A"
	got := l.GetGroup()
	if got != want {
		t.Errorf("Livermorium.GetGroup() = got %v, want %v", got, want)
	}
}

func TestLivermoriumGetCategory(t *testing.T) {
	l := Livermorium{}
	want := "Unknown"
	got := l.GetCategory()
	if got != want {
		t.Errorf("Livermorium.GetCategory() = got %v, want %v", got, want)
	}
}

func TestLivermoriumGetName(t *testing.T) {
	l := Livermorium{}
	want := "Livermorium"
	got := l.GetName()
	if got != want {
		t.Errorf("Livermorium.GetName() = got %v, want %v", got, want)
	}
}

func TestLivermoriumGetSimbol(t *testing.T) {
	l := Livermorium{}
	want := "Lv"
	got := l.GetSimbol()
	if got != want {
		t.Errorf("Livermorium.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestLivermoriumGetAtomicNumber(t *testing.T) {
	l := Livermorium{}
	want := 116
	got := l.GetAtomicNumber()
	if got != want {
		t.Errorf("Livermorium.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestLivermoriumGetAtomicWeight(t *testing.T) {
	l := Livermorium{}
	var want float32 = 293
	got := l.GetAtomicWeight()
	if got != want {
		t.Errorf("Livermorium.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
