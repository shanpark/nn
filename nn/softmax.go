package nn

type SoftMaxWithLossLayer struct {
	loss float64
	y    *Matrix // (1, o)
	t    *Matrix // (1, o)
}

func (layer *SoftMaxWithLossLayer) Forward(x, t *Matrix) float64 {
	layer.t = t
	layer.y = Softmax(x)
	layer.loss = CrossEntropyError(layer.y, layer.t)
	return layer.loss
}

func (layer *SoftMaxWithLossLayer) Backword(dy *Matrix) *Matrix {

}
