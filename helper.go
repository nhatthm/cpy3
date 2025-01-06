package cpy

//go:generate go run resources/scripts/variadic.go

// #include "Python.h"
import "C"

// togo converts a *C.PyObject to a *PyObject.
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}
