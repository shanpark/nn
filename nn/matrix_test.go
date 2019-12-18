package nn

import (
	"fmt"
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
		t.Errorf("m.Col() does not working.")
	}
}

func TestRow(t *testing.T) {
	m := NewMatrix(12, 1)
	if m.Row() != 12 {
		t.Errorf("m.Row() does not working.")
	}
}

func TestDimension(t *testing.T) {
	m := NewMatrix(12, 13)
	if row, col := m.Dimension(); (row != 12) || (col != 13) {
		t.Errorf("m.Dimension() does not working.")
	}
}

func TestGetSet(t *testing.T) {
	m := NewMatrix(2, 3)
	if m.Set(1, 2, 12.345); m.Get(1, 2) != 12.345 {
		t.Errorf("m.Set() or m.At() does not working.")
	}
}

func TestDuplicate(t *testing.T) {
	m := NewMatrix(2, 3)
	m.Set(1, 2, 12.345)

	d := m.Duplicate()
	if d.Get(1, 2) != 12.345 {
		t.Errorf("m.Duplicate() does not working.")
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
		t.Errorf("m.T() does not working.")
	} else if tr.Row() != 3 {
		t.Errorf("m.T() does not working.")
	}

	if tr.Get(2, 1) != 12.345 {
		t.Errorf("m.T() does not working.")
	}

	tr.Set(2, 1, 11.234)
	if m.Get(1, 2) == 11.234 {
		t.Errorf("m.T() returns a shallow copy.")
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
