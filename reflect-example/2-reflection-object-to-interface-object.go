package main

import (
	"fmt"
	"reflect"
)

func main() {
	//first example showing convert reflect.vlaue to float
	floatVar := 3.14
	fmt.Printf("flaotVar %T\n", floatVar)
	v := reflect.ValueOf(floatVar)
	newFloat := v.Interface().(float64)
	fmt.Println(newFloat + 1.2)
	fmt.Printf("floatVar 地址%p\t newfloat 地址%p\t\n", &floatVar, &newFloat)

	//second example showing convert reflect.vlaue to slice
	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)
	v = reflect.Append(v, reflect.ValueOf(2))
	newSlice := v.Interface().([]int)
	newSlice = append(newSlice, 4)

	fmt.Println(newSlice)
	fmt.Printf("SliceVar 指向的地址%p\t newSlice 指向的地址%p\t\n", sliceVar, newSlice) //他们指向的底层数组不是一个
	fmt.Printf("SliceVar 地址%p\t newSlice 地址%p\t\n", &sliceVar, &newSlice)

	s1 := sliceVar[:]
	fmt.Printf("SliceVar 指向的地址%p\t s1 指向的地址%p\t\n", sliceVar, s1) //指向同一个底层数组
	fmt.Printf("SliceVar 地址%p\t s1 地址%p\t\n", &sliceVar, &s1)

}
