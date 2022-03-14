package main

import (
	"fmt"
	"reflect"
)

func printMeta(in interface{}) {
	t := reflect.TypeOf(in)
	k := t.Kind()
	n := t.Name()
	v := reflect.ValueOf(in)
	fmt.Printf("Type =%s\tType.Kind = %s\tType.Name= %s\tVlaue = %s\n", t, k, n, v)
}

type f func(int,int)int
func main() {
	var num int
	str := "hello"
	type book struct {
		name  string
		pages int
	}
	b := book{
		name:  "zbl",
		pages: 200,
	}
	var sub f = func(a,b int) int {
		return a+b
	}

	slice := make([]int,5)
	type s []int
	slice2 := make(s,6)
	ch := make(chan int,1)
	printMeta(num)
	printMeta(str)
	printMeta(b)
	printMeta(sub)
	printMeta(slice)
	printMeta(slice2)
	printMeta(ch)
}
