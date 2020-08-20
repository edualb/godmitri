package element

import "testing"

func TestOganessonGetPeriod(t *testing.T) {
	o := Oganesson{}
	want := "7th period"
	got := o.GetPeriod()
	if got != want {
		t.Errorf("Oganesson.GetPeriod() = got %v, want %v", got, want)
	}
}

func TestOganessonGetGroup(t *testing.T) {
	o := Oganesson{}
	want := "8A"
	got := o.GetGroup()
	if got != want {
		t.Errorf("Oganesson.GetGroup() = got %v, want %v", got, want)
	}
}

func TestOganessonGetCategory(t *testing.T) {
	o := Oganesson{}
	want := "Unknown"
	got := o.GetCategory()
	if got != want {
		t.Errorf("Oganesson.GetCategory() = got %v, want %v", got, want)
	}
}

func TestOganessonGetName(t *testing.T) {
	o := Oganesson{}
	want := "Oganesson"
	got := o.GetName()
	if got != want {
		t.Errorf("Oganesson.GetName() = got %v, want %v", got, want)
	}
}

func TestOganessonGetSimbol(t *testing.T) {
	o := Oganesson{}
	want := "Og"
	got := o.GetSimbol()
	if got != want {
		t.Errorf("Oganesson.GetSimbol() = got %v, want %v", got, want)
	}
}

func TestOganessonGetAtomicNumber(t *testing.T) {
	o := Oganesson{}
	want := 118
	got := o.GetAtomicNumber()
	if got != want {
		t.Errorf("Oganesson.GetAtomicNumber() = got %v, want %v", got, want)
	}
}

func TestOganessonGetAtomicWeight(t *testing.T) {
	o := Oganesson{}
	var want float32 = 294
	got := o.GetAtomicWeight()
	if got != want {
		t.Errorf("Oganesson.GetAtomicWeight() = got %v, want %v", got, want)
	}
}
