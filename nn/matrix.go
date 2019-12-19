package nn

import (
	"fmt"
	"log"
	"math"
	"strings"
)

/*
Matrix todo
*/
type Matrix struct {
	transposed bool
	row        int
	col        int
	data       []float64
}

/*
NewMatrix todo
*/
func NewMatrix(row, col int) *Matrix {
	m := new(Matrix)

	m.row = row
	m.col = col
	m.data = make([]float64, row*col)

	return m
}

/*
NewMatrixWith todo
*/
func NewMatrixWith(row, col int, data []float64) *Matrix {
	m := new(Matrix)

	m.row = row
	m.col = col
	if len(data) >= row*col {
		m.data = data
	} else {
		m.data = make([]float64, row*col)
		copy(m.data, data)
	}

	return m
}

/*
Col todo
*/
func (m *Matrix) Col() int {
	if m.transposed {
		return m.row
	}
	return m.col
}

/*
Row todo
*/
func (m *Matrix) Row() int {
	if m.transposed {
		return m.col
	}
	return m.row
}

/*
Dimension todo
*/
func (m *Matrix) Dimension() (int, int) {
	return m.Row(), m.Col()
}

/*
Set todo
*/
func (m *Matrix) Set(row, col int, val float64) {
	if m.transposed {
		m.data[col*m.col+row] = val
	} else {
		m.data[row*m.col+col] = val
	}
}

/*
Get todo
*/
func (m *Matrix) Get(row, col int) float64 {
	if m.transposed {
		return m.data[col*m.col+row]
	}
	return m.data[row*m.col+col]
}

/*
Duplicate todo
*/
func (m *Matrix) Duplicate() *Matrix {
	d := NewMatrix(m.row, m.col)
	*d = *m
	d.data = make([]float64, len(m.data))
	copy(d.data, m.data)
	return d
}

/*
T todo
*/
func (m *Matrix) T() *Matrix {
	t := m.Duplicate()
	t.transposed = !t.transposed
	return t
}

/*
GetRowVector todo
*/
func (m *Matrix) GetRowVector(r int) *Vector {
	return NewVectorWith(m.Col(), 1, m.data[r*m.Col():])
}

/*
GetColVector todo
*/
func (m *Matrix) GetColVector(c int) *Vector {
	return NewVectorWith(m.Row(), m.Col(), m.data[c:])
}

/*
Dot todo
*/
func (m *Matrix) Dot(n *Matrix) *Matrix {
	if m.Col() != n.Row() {
		log.Panicf("The size of the matrix is not correct. [%dx%d]", n.Row(), n.Col())
	}

	row, col := m.Row(), n.Col()
	res := NewMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			val := m.GetRowVector(r).SumOfMul(n.GetColVector(c))
			res.Set(r, c, val)
		}
	}

	return res
}

/*
Add todo
*/
func (m *Matrix) Add(n *Matrix) *Matrix {
	row, col := m.Dimension()
	row2, col2 := n.Dimension()
	if (row == row2) && (col == col2) {
		res := NewMatrix(row, col)
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				val := m.Get(r, c) + n.Get(r, c)
				res.Set(r, c, val)
			}
		}
		return res
	} else if (row == row2) && (1 == col2) {
		res := NewMatrix(row, col)
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				val := m.Get(r, c) + n.Get(r, 0)
				res.Set(r, c, val)
			}
		}
		return res
	} else if (1 == row2) && (col == col2) {
		res := NewMatrix(row, col)
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				val := m.Get(r, c) + n.Get(0, c)
				res.Set(r, c, val)
			}
		}
		return res
	} else if (row == row2) && (col == 1) {
		res := NewMatrix(row2, col2)
		for r := 0; r < row2; r++ {
			for c := 0; c < col2; c++ {
				val := m.Get(r, 0) + n.Get(r, c)
				res.Set(r, c, val)
			}
		}
		return res
	} else if (row == 1) && (col == col2) {
		res := NewMatrix(row2, col2)
		for r := 0; r < row2; r++ {
			for c := 0; c < col2; c++ {
				val := m.Get(0, c) + n.Get(r, c)
				res.Set(r, c, val)
			}
		}
		return res
	} else {
		log.Panicf("The 'size' of the matrix is not correct.")
		return nil
	}
}

/*
Mul element-wise multiplication.
*/
func (m *Matrix) Mul(n *Matrix) *Matrix {
	row, col := m.Dimension()
	row2, col2 := n.Dimension()
	if (row == row2) && (col == col2) {
		res := NewMatrix(row, col)
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				val := m.Get(r, c) * n.Get(r, c)
				res.Set(r, c, val)
			}
		}
		return res
	} else {
		log.Panicf("The 'size' of the matrix is not correct.")
		return nil
	}
}

/*
Sum todo
*/
func (m *Matrix) Sum() float64 {
	row, col := m.Dimension()
	res := 0.0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			res += m.Get(r, c)
		}
	}
	return res
}

/*
ColSum todo
*/
func (m *Matrix) ColSum() *Matrix {
	row, col := m.Dimension()
	res := NewMatrix(1, m.Col())
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			res.Set(0, c, res.Get(0, c)+m.Get(r, c))
		}
	}
	return res
}

/*
Max return the biggest element value.
*/
func (m *Matrix) Max() float64 {
	row, col := m.Dimension()
	res := m.Get(0, 0)
	var val float64
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			val = m.Get(r, c)
			if val > res {
				res = val
			}
		}
	}
	return res
}

/*
AddScalar todo
*/
func (m *Matrix) AddScalar(scalar float64) *Matrix {
	row, col := m.Dimension()
	res := NewMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			res.Set(r, c, m.Get(r, c)+scalar)
		}
	}
	return res
}

/*
MulScalar todo
*/
func (m *Matrix) MulScalar(scalar float64) *Matrix {
	row, col := m.Dimension()
	res := NewMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			res.Set(r, c, m.Get(r, c)*scalar)
		}
	}
	return res
}

/*
Exp todo
*/
func (m *Matrix) Exp() *Matrix {
	row, col := m.Dimension()
	res := NewMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			res.Set(r, c, math.Exp(m.Get(r, c)))
		}
	}
	return res
}

/*
Log todo
*/
func (m *Matrix) Log() *Matrix {
	row, col := m.Dimension()
	res := NewMatrix(row, col)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			res.Set(r, c, math.Log(m.Get(r, c)))
		}
	}
	return res
}

/*
Softmax todo
*/
func (m *Matrix) Softmax() *Matrix {
	a := m.AddScalar(-m.Max()) // prevent overflow
	exp := a.Exp()
	sum := exp.Sum()
	y := exp.MulScalar(1.0 / sum)
	return y
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
