package main

type ValueLimiter struct {
	nextComponent component
	lowerBound    int
	upperBound    int
}

func (vl *ValueLimiter) handle(dict map[string]interface{}) map[string]interface{} {
	for k, v := range dict {
		if v.(int) >= vl.upperBound || v.(int) < vl.lowerBound {
			delete(dict, k)
		}
	}
	return vl.nextComponent.handle(dict)
}

func (vl *ValueLimiter) setNext(next component) {
	vl.nextComponent = next
}
