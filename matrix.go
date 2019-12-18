package nn

import (
	"fmt"
	"strings"
)

type Matrix struct {
	transposed bool
	Rows       int
	Cols       int
	data       []float64
}

func (m *Matrix) Width() int {
	if m.transposed {
		return m.Rows
	}
	return m.Cols
}

func (m *Matrix) Height() int {
	if m.transposed {
		return m.Cols
	}
	return m.Rows
}

func (m *Matrix) Dimension() (int, int) {
	return m.Width(), m.Height()
}

func (m *Matrix) Set(row, col int, val float64) {
	if m.transposed {
		m.data[col*m.Cols+row] = val
	} else {
		m.data[row*m.Cols+col] = val
	}
}

func (m *Matrix) At(row, col int) float64 {
	if m.transposed {
		return m.data[col*m.Cols+row]
	}
	return m.data[row*m.Cols+col]
}

func (m *Matrix) T() Matrix {
	t := *m
	t.transposed = !m.transposed
	t.data = make([]float64, len(m.data))
	copy(t.data, m.data)
	return t
}

func (m Matrix) String() string {
	var result strings.Builder

	width, height := m.Dimension()

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			result.WriteString(fmt.Sprintf("%6.3f", m.At(row, col)))
			if col == width-1 {
				result.WriteString("\n")
			} else {
				result.WriteString(",\t")
			}
		}
	}

	return result.String()
}

func NewMatrix(rows, cols int) *Matrix {
	m := new(Matrix)

	m.Rows = rows
	m.Cols = cols
	m.data = make([]float64, rows*cols)

	return m
}
