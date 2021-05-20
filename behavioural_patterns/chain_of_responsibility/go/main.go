package main

import "fmt"

func main() {
	c := map[string]interface{}{"a": 2, "b": "seven", "c": 23, "d": 42, "e": 76, "f": 124}
	ct := ChainTerminator{}
	vl2 := ValueLimiter{nextComponent: &ct, lowerBound: 1, upperBound: 1000}
	pr := PowerRaiser{nextComponent: &vl2, exponent: 2}
	vl1 := ValueLimiter{nextComponent: &pr, lowerBound: 1, upperBound: 100}
	nvf := NumericValueFinder{nextComponent: &vl1}
	result := nvf.handle(c)
	fmt.Println("result = ", result)
}
