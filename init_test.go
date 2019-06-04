package dough

import (
	"strconv"
	"strings"
)

type testNum struct {
	String  string
	Float   float64
	Integer int
}

var TestLargeNums []testNum

func init() {
	// Create a bunch of test numbers
	for i := -999.99; i < 999.99; i += .03 {
		str := strconv.FormatFloat(i, 'f', 2, 64) // 2 is the format number
		strReplace := strings.Replace(str, ".", "", -1)
		float, _ := strconv.ParseFloat(str, 64)
		integer, _ := strconv.Atoi(strReplace)
		TestLargeNums = append(TestLargeNums, testNum{str, float, integer})
		// fmt.Printf("%#v %#v %#v \n", str, float, integer)
	}
}
