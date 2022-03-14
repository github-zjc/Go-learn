package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
}

func (s *student) DoHomework(num int) {
	fmt.Printf("%s is doing homework %d\n", s.name, num)
}

func (s *student) Sub(a, b int) int {
	fmt.Println(a - b)
	return a - b
}

func main() {
	//use reflect invoke the DoHomework of a student
	s := student{name: "zjc"}
	v := reflect.ValueOf(&s)
	methodV := v.MethodByName("DoHomework")
	if methodV.IsValid() {
		in := reflect.ValueOf(100)
		methodV.Call([]reflect.Value{in})
	}

	methodV = v.MethodByName("Sub")
	if methodV.IsValid() {
		in := []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(2)}
		methodV.Call(in)
	}
}
