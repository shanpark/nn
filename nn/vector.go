package nn

import "log"

/*
Vector A struct for row or column of a matrix.
*/
type Vector struct {
	size int
	step int
	data []float64
}

/*
NewVector return the pointer of a zero valued Vector.
*/
func NewVector(size, step int) *Vector {
	if size < 0 {
		log.Panicf("'size' parameter should not be smaller than 0.")
	}
	if step < 1 {
		log.Panicf("'step' parameter should not be smaller than 1.")
	}

	return NewVectorWith(size, step, make([]float64, size*(step-1)+1))
}

/*
NewVectorWith return the pointer of a Vector initialized with parameters.
              Invalid parameters make an invalid Vector.
*/
func NewVectorWith(size, step int, data []float64) *Vector {
	if size < 0 {
		log.Panicf("'size' parameter should not be smaller than 0.")
	}
	if step < 1 {
		log.Panicf("'step' parameter should not be smaller than 1.")
	}

	v := new(Vector)
	v.size = size
	v.step = step
	v.data = data
	return v
}

/*
Get return the value at i.
*/
func (v *Vector) Get(i int) float64 {
	return v.data[i*v.step]
}

/*
Set set the value at i.
*/
func (v *Vector) Set(i int, val float64) {
	v.data[i*v.step] = val
}

/*
Size get the size of the vector
*/
func (v *Vector) Size() int {
	return v.size
}

/*
SumOfMul todo
*/
func (v *Vector) SumOfMul(n *Vector) float64 {
	size := v.Size()
	if size != n.Size() {
		log.Panicf("The size of the vector is not correct. [%d != %d]", v.Size(), n.Size())
	}

	res := 0.0
	for inx := 0; inx < size; inx++ {
		res += v.Get(inx) * n.Get(inx)
	}

	return res
}
