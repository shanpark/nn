package nn

import (
	"fmt"
	"log"
	"strings"
)

type Matrix struct {
	transposed bool
	row        int
	col        int
	data       []float64
}

func NewMatrix(row, col int) *Matrix {
	m := new(Matrix)

	m.row = row
	m.col = col
	m.data = make([]float64, row*col)

	return m
}

func (m *Matrix) Col() int {
	if m.transposed {
		return m.row
	}
	return m.col
}

func (m *Matrix) Row() int {
	if m.transposed {
		return m.col
	}
	return m.row
}

func (m *Matrix) Dimension() (int, int) {
	return m.Row(), m.Col()
}

func (m *Matrix) Set(row, col int, val float64) {
	if m.transposed {
		m.data[col*m.col+row] = val
	} else {
		m.data[row*m.col+col] = val
	}
}

func (m *Matrix) Get(row, col int) float64 {
	if m.transposed {
		return m.data[col*m.col+row]
	}
	return m.data[row*m.col+col]
}

func (m *Matrix) Duplicate() *Matrix {
	d := NewMatrix(m.row, m.col)
	*d = *m
	d.data = make([]float64, len(m.data))
	copy(d.data, m.data)
	return d
}

func (m *Matrix) T() *Matrix {
	t := m.Duplicate()
	t.transposed = !t.transposed
	return t
}

func (m *Matrix) Dot(n *Matrix) *Matrix {
	if m.Col() != n.Row() {
		log.Panicf("The size of the matrix is not correct. [%dx%d]", n.Row(), n.Col())
	}

	row, col := m.Row(), n.Col()
	r := NewMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {

		}
	}

	return r
}

func (m *Matrix) String() string {
	var result strings.Builder

	row, col := m.Dimension()

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			result.WriteString(fmt.Sprintf("%6.3f", m.Get(r, c)))
			if c == col-1 {
				result.WriteString("\n")
			} else {
				result.WriteString(",\t")
			}
		}
	}

	return result.String()
}
