package resample

import (
	"math"
)

// SampleUp Up-samples a signal, conceptually increasing its sampling rate by an integer factor
func SampleUp(in []float64, factor int) []float64 {
	if factor <= 1 {
		return in
	}
	out := make([]float64, len(in)*factor)
	for i, v := range in {
		out[i*factor] = v
	}

	return out
}

// SampleDown Down-samples a signal, conceptually decreasing its sampling rate by an integer factor
func SampleDown(in []float64, factor int) []float64 {
	if factor <= 1 {
		return in
	}
	out := make([]float64, int(math.Ceil(float64(len(in))/float64(factor))))
	for i := range out {
		out[i] = in[i*factor]
	}

	return out
}

// Interpolate Up-samples a signal by an integer factor with linear interpolation
func Interpolate(in []float64, factor int) []float64 {
	if len(in) == 0 {
		return in
	}

	out := make([]float64, int(len(in)*factor-(factor-1)))
	for i := 0; i < len(out); i++ {
		start := in[i/factor]
		var next float64
		if (i+factor)/factor >= len(in) {
			next = start
		} else {
			next = in[(i+factor)/factor]
		}
		scale := next - start
		factor := (float64(i%factor) / float64(factor))
		out[i] = start + factor*scale
	}
	return out
}
