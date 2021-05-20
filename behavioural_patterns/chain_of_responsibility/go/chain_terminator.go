package main

type ChainTerminator struct {
	nextComponent component
}

func (ct *ChainTerminator) handle(dict map[string]interface{}) map[string]interface{} {
	return dict
}

func (ct *ChainTerminator) setNext(next component) {
	ct.nextComponent = next
}
