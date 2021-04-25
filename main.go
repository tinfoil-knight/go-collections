package main

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrInvalidSyntax = errors.New("invalid syntax")
)

type Array struct {
	internal *[]string
}

func (a Array) String() string {
	return fmt.Sprintf("Array{%v}", *a.internal)
}

// thisArg parameter has not been implemented
func (arr *Array) FindIndex(callback func(element string) bool) int {
	a := arr.internal
	for i, v := range *a {
		// index, array parameters have not been implemented
		if callback(v) {
			return i
		}
	}
	return -1
}

func (arr *Array) Pop() string {
	a := arr.internal
	s := len(*a)
	if s == 0 {
		return ""
	}
	r := (*a)[s-1]
	*a = (*a)[:s-1]
	return r
}

// https://stackoverflow.com/arr/34816623/12531621
func (arr *Array) Reverse() *Array {
	a := arr.internal
	s := len(*a)
	l := s - 1
	for i := 0; i < s/2; i++ {
		(*a)[i], (*a)[l-i] = (*a)[l-i], (*a)[i]
	}
	return arr
}

func (arr *Array) Shift() string {
	a := arr.internal
	s := len(*a)
	if s == 0 {
		return ""
	}
	r := (*a)[0]
	*a = (*a)[1:]
	return r
}

// compareFn parameter has not been implemented
func (arr *Array) Sort() *Array {
	a := arr.internal
	sort.Strings(*a)
	return arr
}

func main() {
	// Playground
	arr := []string{"1", "2", "3"}
	a := Array{&arr}
	fmt.Printf("%v\n", a)
}
