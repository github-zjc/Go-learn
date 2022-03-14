package main

import (
	"fmt"
	"reflect"
)

func main() {
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	fmt.Printf("is float canSet: %v\tcanAddr %v\n",v.CanSet(),v.CanAddr())

	vp := reflect.ValueOf(&fVar)
	fmt.Printf("is float canSet: %v\tcanAddr %v\n",vp.Elem().CanSet(),vp.Elem().CanAddr())

	vp.Elem().SetFloat(6.666)
	fmt.Println(fVar)
}
