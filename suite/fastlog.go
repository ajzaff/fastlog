package suite

import "math"

// FastLog2 implements fast approximate Log2 using the MSE constants.
// Does no range checking.
func FastLog2(x, c0, c1 float32) float32 {
	tmp := math.Float32bits(x)
	expb := uint64(tmp) >> 23
	tmp = (tmp & 0x7fffff) | (0x7f << 23)
	out := math.Float32frombits(tmp) - 1
	return out*(c0+c1*out) - 127 + float32(expb)
}
