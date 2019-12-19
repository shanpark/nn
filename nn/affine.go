package nn

type AffineLayer struct {
	w  *Matrix // (i, o)
	b  *Matrix // (1, o)
	x  *Matrix // (N, i)
	dw *Matrix //
	db *Matrix //
}

func NewAffine(w *Matrix, b *Matrix) *AffineLayer {
	layer := new(AffineLayer)
	layer.w = w
	layer.b = b
	return layer
}

/*
Forward 순전파 구현.
 (N, i) 크기의 Matrix를 입력으로 받아서 (N, o) 크기의 Matrix를 반환한다.
*/
func (layer *AffineLayer) Forward(x *Matrix) *Matrix {
	layer.x = x
	out := x.Dot(layer.w).Add(layer.b)
	return out
}

/*
Backword 역전파 구현.
 (N, o) 크기의 Matrix를 입력으로 받아 (1, o) 크기의 Matrix를 반환한다.
*/
func (layer *AffineLayer) Backword(dy *Matrix) *Matrix {
	dx := dy.Dot(layer.w.T())
	layer.dw = layer.x.T().Dot(dy)
	layer.db = dy.ColSum()

	return dx
}
