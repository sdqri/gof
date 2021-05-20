package main

type NumericValueFinder struct {
	nextComponent component
}

func (nvf *NumericValueFinder) handle(dict map[string]interface{}) map[string]interface{} {
	for k, v := range dict {
		_, t := v.(int)
		if t == false {
			delete(dict, k)
		}
	}
	return nvf.nextComponent.handle(dict)
}

func (nvf *NumericValueFinder) setNext(next component) {
	nvf.nextComponent = next
}
