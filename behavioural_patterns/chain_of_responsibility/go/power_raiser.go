package main

import (
	"math"
)

type PowerRaiser struct {
	nextComponent component
	exponent      int
}

func (pr *PowerRaiser) handle(dict map[string]interface{}) map[string]interface{} {
	for k, v := range dict {
		dict[k] = int(math.Pow(float64(v.(int)), float64(pr.exponent)))
	}
	return pr.nextComponent.handle(dict)
}

func (pr *PowerRaiser) setNext(next component) {
	pr.nextComponent = next
}
