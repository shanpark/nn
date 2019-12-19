package nn

func CrossEntropyError(y, t *Matrix) float64 {
	return -(y.AddScalar(1e-9).Log().Mul(t).Sum())
}
