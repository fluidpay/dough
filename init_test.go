package dough

import (
	"strconv"
)

type testNum struct {
	String1 string
	String2 string
	String3 string
	Float1  float64
	Float2  float64
	Float3  float64
	Integer int
}

var TestLargeNums []testNum

func init() {
	// Create a bunch of test numbers
	for i := -99999; i <= 99999; i++ {
		float1 := IntToFloat(i, 1)
		float2 := IntToFloat(i, 2)
		float3 := IntToFloat(i, 3)
		str1 := strconv.FormatFloat(float1, 'f', 1, 64)
		str2 := strconv.FormatFloat(float2, 'f', 2, 64)
		str3 := strconv.FormatFloat(float3, 'f', 3, 64)

		TestLargeNums = append(TestLargeNums, testNum{str1, str2, str3, float1, float2, float3, i})
	}
}
