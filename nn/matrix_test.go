package nn

import (
	"fmt"
	"math"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(5, 10)
	if m == nil {
		t.Errorf("NewMatrix() return nil.")
	}
}

func TestCol(t *testing.T) {
	m := NewMatrix(1, 13)
	if m.Col() != 13 {
		t.Errorf("m.Col() does not work.")
	}
}

func TestRow(t *testing.T) {
	m := NewMatrix(12, 1)
	if m.Row() != 12 {
		t.Errorf("m.Row() does not work.")
	}
}

func TestDimension(t *testing.T) {
	m := NewMatrix(12, 13)
	if row, col := m.Dimension(); (row != 12) || (col != 13) {
		t.Errorf("m.Dimension() does not work.")
	}
}

func TestGetSet(t *testing.T) {
	m := NewMatrix(2, 3)
	if m.Set(1, 2, 12.345); m.Get(1, 2) != 12.345 {
		t.Errorf("m.Set() or m.At() does not work.")
	}
}

func TestDuplicate(t *testing.T) {
	m := NewMatrix(2, 3)
	m.Set(1, 2, 12.345)

	d := m.Duplicate()
	if d.Get(1, 2) != 12.345 {
		t.Errorf("m.Duplicate() does not work.")
	}

	d.Set(1, 2, 11.234)
	if m.Get(1, 2) == 11.234 {
		t.Errorf("m.Duplicate() returns a shallow copy.")
	}
}

func TestT(t *testing.T) {
	m := NewMatrix(2, 3)
	m.Set(1, 2, 12.345)
	tr := m.T()
	if tr.Col() != 2 {
		t.Errorf("m.T() does not work.")
	} else if tr.Row() != 3 {
		t.Errorf("m.T() does not work.")
	}

	if tr.Get(2, 1) != 12.345 {
		t.Errorf("m.T() does not work.")
	}

	tr.Set(2, 1, 11.234)
	if m.Get(1, 2) == 11.234 {
		t.Errorf("m.T() returns a shallow copy.")
	}
}

func TestDot(t *testing.T) {
	m := NewMatrixWith(2, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	n := NewMatrixWith(3, 2, []float64{6.0, 5.0, 4.0, 3.0, 2.0, 1.0})

	res := m.Dot(n)
	if (res.Get(0, 0) != 20) || (res.Get(0, 1) != 14) || (res.Get(1, 0) != 56) || (res.Get(1, 1) != 41) {
		t.Errorf("m.Dot() does not work.")
	}
}

func TestAdd(t *testing.T) {
	m := NewMatrixWith(2, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	n := NewMatrixWith(2, 3, []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0})
	n2 := NewMatrixWith(1, 3, []float64{1.0, 1.0, 1.0})
	n3 := NewMatrixWith(2, 1, []float64{1.0, 1.0})

	res := m.Add(n)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 3.0) || (res.Get(0, 2) != 4.0) ||
		(res.Get(1, 0) != 5.0) || (res.Get(1, 1) != 6.0) || (res.Get(1, 2) != 7.0) {
		t.Errorf("m.Add() does not work.")
	}

	res = m.Add(n2)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 3.0) || (res.Get(0, 2) != 4.0) ||
		(res.Get(1, 0) != 5.0) || (res.Get(1, 1) != 6.0) || (res.Get(1, 2) != 7.0) {
		t.Errorf("m.Add() does not work.")
	}

	res = m.Add(n3)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 3.0) || (res.Get(0, 2) != 4.0) ||
		(res.Get(1, 0) != 5.0) || (res.Get(1, 1) != 6.0) || (res.Get(1, 2) != 7.0) {
		t.Errorf("m.Add() does not work.")
	}

	res = n2.Add(m)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 3.0) || (res.Get(0, 2) != 4.0) ||
		(res.Get(1, 0) != 5.0) || (res.Get(1, 1) != 6.0) || (res.Get(1, 2) != 7.0) {
		t.Errorf("m.Add() does not work.")
	}

	res = n3.Add(m)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 3.0) || (res.Get(0, 2) != 4.0) ||
		(res.Get(1, 0) != 5.0) || (res.Get(1, 1) != 6.0) || (res.Get(1, 2) != 7.0) {
		t.Errorf("m.Add() does not work.")
	}
}

func TestMul(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})
	n := NewMatrixWith(2, 2, []float64{2.0, 2.0, 2.0, 2.0})

	res := m.Mul(n)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 4.0) ||
		(res.Get(1, 0) != 6.0) || (res.Get(1, 1) != 8.0) {
		t.Errorf("m.Mul() does not work.")
	}
}

func TestSum(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})

	res := m.Sum()
	if res != 10.0 {
		t.Errorf("m.Sum() does not work.")
	}
}

func TestColSum(t *testing.T) {
	m := NewMatrixWith(2, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})

	res := m.ColSum()
	if (res.Get(0, 0) != 5) || (res.Get(0, 1) != 7) || (res.Get(0, 2) != 9) {
		t.Errorf("m.ColSum() does not work.")
	}
}

func TestMax(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})

	res := m.Max()
	if res != 4.0 {
		t.Errorf("m.Max() does not work.")
	}
}

func TestAddScalar(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})

	res := m.AddScalar(1.0)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 3.0) ||
		(res.Get(1, 0) != 4.0) || (res.Get(1, 1) != 5.0) {
		t.Errorf("m.AddScalar() does not work.")
	}
}

func TestMulScalar(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})

	res := m.MulScalar(2.0)
	if (res.Get(0, 0) != 2.0) || (res.Get(0, 1) != 4.0) ||
		(res.Get(1, 0) != 6.0) || (res.Get(1, 1) != 8.0) {
		t.Errorf("m.MulScalar() does not work.")
	}
}

func TestExp(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})

	res := m.Exp()
	if (res.Get(0, 0) != math.Exp(1.0)) || (res.Get(0, 1) != math.Exp(2.0)) ||
		(res.Get(1, 0) != math.Exp(3.0)) || (res.Get(1, 1) != math.Exp(4.0)) {
		t.Errorf("m.Exp() does not work.")
	}
}

func TestLog(t *testing.T) {
	m := NewMatrixWith(2, 2, []float64{1.0, 2.0, 3.0, 4.0})

	res := m.Log()
	if (res.Get(0, 0) != math.Log(1.0)) || (res.Get(0, 1) != math.Log(2.0)) ||
		(res.Get(1, 0) != math.Log(3.0)) || (res.Get(1, 1) != math.Log(4.0)) {
		t.Errorf("m.Log() does not work.")
	}
}

func TestString(t *testing.T) {
	m := NewMatrix(2, 3)
	m.Set(0, 0, 1)
	m.Set(0, 1, 2)
	m.Set(0, 2, 3)
	m.Set(1, 0, 4)
	m.Set(1, 1, 5)
	m.Set(1, 2, 6)
	fmt.Printf("%v", m)
}
