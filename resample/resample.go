package resample

import "math"

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
