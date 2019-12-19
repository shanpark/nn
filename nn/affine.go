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
 (N, o) 크기의 Matrix를 입력으로 받아 (N, i) 크기의 Matrix를 반환한다.
*/
func (layer *AffineLayer) Backword(dy *Matrix) *Matrix {
	dx := dy.Dot(layer.w.T())      // (N, o)(o, i) = (N, i)
	layer.dw = layer.x.T().Dot(dy) // (i, N)(N, o) = (i, o)
	layer.db = dy.ColSum()         // ??? 이걸 batchSize(N)으로 나누어 주어야 하는 거 아닌가?

	return dx
}
