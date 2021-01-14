package main

import (
	"fmt"
	"reflect"
)

type Example struct {
	data        string
	anotherdata int
}

func (v *Example) getData() string {
	return v.data
}

func (v *Example) getAnotherData() int {
	return v.anotherdata * 5
}

type ExampleTwo struct {
	Example
}

type ExampleThree struct {
	Base Example
}

func main() {
	exampleTwo := &ExampleTwo{
		Example{"Prasad", 5},
	}
	fmt.Println(exampleTwo.getData())
	fmt.Println(reflect.TypeOf(exampleTwo))
}
