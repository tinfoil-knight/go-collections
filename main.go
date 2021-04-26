package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
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
func (arr *Array) Every(callback func(element string) bool) bool {
	a := arr.internal
	for _, v := range *a {
		// index, array parameters have not been implemented
		if !callback(v) {
			return false
		}
	}
	return true
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

func (arr *Array) Includes(valueToFind string, fromIndex ...int) bool {
	if len(fromIndex) > 1 {
		panic(ErrInvalidSyntax)
	}
	a := arr.internal
	s := len(*a)
	start := 0
	if len(fromIndex) == 1 {
		fI := fromIndex[0]
		if fI >= s {
			return false
		}
		if fI < 0 {
			computedIndex := s + fI
			if !(computedIndex <= -1*s) {
				start = computedIndex
			}
		} else {
			start = fI
		}
	}
	for i := start; i < s; i++ {
		if (*a)[i] == valueToFind {
			return true
		}
	}
	return false
}

func (arr *Array) Join(separator ...string) string {
	if len(separator) > 1 {
		panic(ErrInvalidSyntax)
	}
	sep := ""
	if len(separator) == 1 {
		sep = separator[0]
	}
	a := arr.internal
	return strings.Join(*a, sep)
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

// thisArg parameter has not been implemented
func (arr *Array) Some(callback func(element string) bool) bool {
	a := arr.internal
	for _, v := range *a {
		// index, array parameters have not been implemented
		if callback(v) {
			return true
		}
	}
	return false
}

// compareFn parameter has not been implemented
func (arr *Array) Sort() *Array {
	a := arr.internal
	sort.Strings(*a)
	return arr
}

func (arr *Array) Unshift(elementN ...string) int {
	a := arr.internal
	*a = append(elementN, *a...)
	return len(*a)
}

func main() {
	// Playground
	arr := []string{"1", "2", "3"}
	a := Array{&arr}
	fmt.Printf("%v\n", a)
}
