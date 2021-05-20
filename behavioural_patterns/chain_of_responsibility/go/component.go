package main

type component interface {
	handle(map[string]interface{}) map[string]interface{}
	setNext(component)
}
