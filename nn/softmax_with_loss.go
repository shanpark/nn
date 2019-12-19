package nn

type SoftMaxWithLossLayer struct {
	loss float64
	y    *Matrix // (1, o)
	t    *Matrix // (1, o)
}

func (layer *SoftMaxWithLossLayer) Forward(x, t *Matrix) float64 {
	layer.t = t
	layer.y = x.Softmax()
	layer.loss = CrossEntropyError(layer.y, layer.t)
	return layer.loss
}

/*
Backword 역전파 구현.
 (N, o) 크기의 Matrix를 입력으로 받아 (1, o) 크기의 Matrix를 반환한다.
*/
func (layer *SoftMaxWithLossLayer) Backword(dy *Matrix) *Matrix {
	batchSize := dy.Row()
	return layer.y.Add(layer.t.MulScalar(-1.0)).MulScalar(1.0 / float64(batchSize)) // ??? 왜 손실의 크기를 batchSize로 나누어 주는가? 이게 꼭 필요한가?
}
