package dough

import (
	"strconv"
	"strings"
)

type testNum struct {
	String  string
	Float1  float64
	Float2  float64
	Float3  float64
	Integer int
}

var TestLargeNums []testNum

func init() {
	// Create a bunch of test numbers
	for i := -999.99; i < 999.99; i += .03 {
		str2 := strconv.FormatFloat(i, 'f', 2, 64) // 1 is the format number
		strReplace := strings.Replace(str2, ".", "", -1)
		str1 := strReplace[0:len(strReplace) - 1] // doing string manipulation to get a 10ths place string
		str1 += "."
		str1 += strReplace[len(strReplace) - 1:]
		str3 := strReplace[0:len(strReplace) - 3] // doing string manipulation to get a 1000ths place string
		str3 += "."
		str3 += strReplace[len(strReplace) - 3:]
		float1, _ := strconv.ParseFloat(str1, 64)
		float2, _ := strconv.ParseFloat(str2, 64)
		float3, _ := strconv.ParseFloat(str3, 64)
		integer, _ := strconv.Atoi(strReplace)
		TestLargeNums = append(TestLargeNums, testNum{str2, float1, float2, float3, integer})
	}
}
