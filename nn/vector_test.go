package nn

import "testing"

func TestNewVector(t *testing.T) {
	func() {
		defer func() {
			recover()
		}()
		NewVector(-1, 3)
		t.Errorf("The 'size' of a vector should not be smaller than 0.")
	}()

	func() {
		defer func() {
			recover()
		}()
		NewVector(1, 0)
		t.Errorf("The 'step' of a vector should not be smaller than 1.")
	}()

	func() {
		defer func() {
			recover()
		}()
		NewVectorWith(-1, 3, []float64{1.0, 2.0})
		t.Errorf("The 'size' of a vector should not be smaller than 0.")
	}()

	func() {
		defer func() {
			recover()
		}()
		NewVectorWith(1, 0, []float64{1.0, 2.0})
		t.Errorf("The 'step' of a vector should not be smaller than 1.")
	}()

	v := NewVector(3, 3)
	if v == nil {
		t.Errorf("NewVector() return nil.")
	}

	data := make([]float64, 7)
	v = NewVectorWith(3, 3, data)
	if v == nil {
		t.Errorf("NewVectorWith() return nil.")
	}
}

func TestVectorGetSet(t *testing.T) {
	v := NewVector(3, 3)
	if v.Set(1, 12.345); v.Get(1) != 12.345 {
		t.Errorf("v.Set() or v.Get() does not working.")
	}
}

func TestVectorSumOfMul(t *testing.T) {
	v := NewVectorWith(2, 1, []float64{1.0, 2.0})

	k := NewVectorWith(2, 2, []float64{3.0, 0.0, 4.0})
	l := NewVectorWith(3, 1, []float64{3.0, 0.0, 4.0})

	res := v.SumOfMul(k)
	if res != 11.0 {
		t.Errorf("v.SumOfMul() does not working.")
	}

	func() {
		defer func() {
			recover() // recover!
		}()

		res = v.SumOfMul(l) // panic!
		t.Errorf("v.SumOfMul() does not working.")
	}()
}
