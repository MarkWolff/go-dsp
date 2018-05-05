package resample

import "testing"

func TestSampleUp(t *testing.T) {
	vec1 := []float64{1, 2, 3, 4}
	vec1Factor := 3
	vec1Exp := []float64{1, 0, 0, 2, 0, 0, 3, 0, 0, 4, 0, 0} // 3x
	vec1Result := SampleUp(vec1, vec1Factor)

	vec2 := []float64{1}
	vec2Factor := 2
	vec2Exp := []float64{1, 0} // 2x
	vec2Result := SampleUp(vec2, vec2Factor)

	vec3 := []float64{}
	vec3Factor := 10
	vec3Exp := []float64{}
	vec3Result := SampleUp(vec3, vec3Factor)

	vec4 := []float64{}
	vec4Exp := []float64{}
	vec4Factor := -1
	vec4Result := SampleUp(vec4, vec4Factor)

	if !testEq(vec1Result, vec1Exp) {
		t.Fatalf("Failed SampleDown #1\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec1,
			vec1Result,
			vec1Exp)
	}
	if !testEq(vec2Result, vec2Exp) {
		t.Fatalf("Failed SampleDown #2\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec2,
			SampleDown(vec2, vec2Factor),
			vec2Exp)
	}
	if !testEq(vec3Result, vec3Exp) {
		t.Fatalf("Failed SampleDown #3\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec3,
			SampleDown(vec3, vec3Factor),
			vec3Exp)
	}

	if !testEq(vec4Result, vec4Exp) {
		t.Fatalf("Failed SampleDown #4\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec4,
			SampleDown(vec4, vec4Factor),
			vec4Exp)
	}

}

func TestSampleDown(t *testing.T) {
	vec1 := []float64{1, 0, 0, 2, 0, 0, 3, 0, 0, 4} // 3x
	vec1Exp := []float64{1, 2, 3, 4}
	vec1Factor := 3
	vec1Result := SampleDown(vec1, vec1Factor)

	vec2 := []float64{1}
	vec2Exp := []float64{1}
	vec2Factor := 2
	vec2Result := SampleDown(vec2, vec2Factor)

	vec3 := []float64{}
	vec3Exp := []float64{}
	vec3Factor := 111
	vec3Result := SampleDown(vec3, vec3Factor)

	vec4 := []float64{}
	vec4Exp := []float64{}
	vec4Factor := 0
	vec4Result := SampleDown(vec4, vec4Factor)

	if !testEq(vec1Result, vec1Exp) {
		t.Fatalf("Failed SampleDown #1\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec1,
			vec1Result,
			vec1Exp)
	}
	if !testEq(vec2Result, vec2Exp) {
		t.Fatalf("Failed SampleDown #2\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec2,
			SampleDown(vec2, vec2Factor),
			vec2Exp)
	}
	if !testEq(vec3Result, vec3Exp) {
		t.Fatalf("Failed SampleDown #3\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec3,
			SampleDown(vec3, vec3Factor),
			vec3Exp)
	}

	if !testEq(vec4Result, vec4Exp) {
		t.Fatalf("Failed SampleDown #4\nOriginal: \t%v\nResult: \t%v\nExpect: \t%v\n",
			vec4,
			SampleDown(vec4, vec4Factor),
			vec4Exp)
	}

}

func testEq(a, b []float64) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
