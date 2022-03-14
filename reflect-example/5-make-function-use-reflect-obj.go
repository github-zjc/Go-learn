package main

import (
	"fmt"
	"reflect"
	"time"
)

func TimeSleep() {
	time.Sleep(time.Second)
	fmt.Println("hello zjc")
}

func MakeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		fmt.Println("expect a Func")
		return nil
	}
	//包装这个函数
	warpper := reflect.MakeFunc(tf, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		result := vf.Call(args)
		end := time.Now()
		fmt.Printf("the function takes %v\n", end.Sub(start))
		return result
	})

	return warpper.Interface()

}

//利用reflect创建一个加法的函数
func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("diffrent Type")
		return nil
	}
	k := a.Kind()
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())} //value.int返回的是int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())} //value.uint返回的是uint64
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())} //value.float返回的是float64
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func MakeSum(in interface{}) {
	fn := reflect.ValueOf(in).Elem() //返回一个可以修改的value
	if fn.Kind() != reflect.Func {
		fmt.Println("expect a function")
		return
	}
	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)

}

//使用reflect编写一个切片倒序的函数

func InvertSlice(agrs []reflect.Value) (result []reflect.Value) {
	inSlice, n := agrs[0], agrs[0].Len() //拿到切片
	fmt.Println(inSlice, n)
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}

}

// func Bind(p interface{}) {
// 	invert := reflect.ValueOf(p).Elem()

// 	invert.Set(reflect.MakeFunc(invert.Type(), InvertSlice))

// }

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	invert := reflect.ValueOf(p).Elem()

	invert.Set(reflect.MakeFunc(invert.Type(), f))

}

func sub(a, b int) int {
	return a - b
}
func main() {
	//由例子可见，反射可以增加底层函数的功能
	//使用reflect包装某个函数，添加一个运行时间功能
	makefunc := MakeFunc(TimeSleep).(func())
	makefunc()

	makesub := MakeFunc(sub).(func(int, int) int)
	fmt.Println(makesub(5, 4))

	//将包装TimeSleep的函数功能给一个函数
	var makefunc2 func()
	makefunc2 = MakeFunc(TimeSleep).(func())
	makefunc2()

	var makesub2 func(int, int) int
	makesub2 = MakeFunc(sub).(func(int, int) int)
	fmt.Println(makesub2(5, 4))

	//测试sum功能是否成功
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	MakeSum(&intSum)
	MakeSum(&floatSum)
	MakeSum(&stringSum)

	fmt.Println(intSum(1, 2))
	fmt.Println(floatSum(3.3, 3.3))
	fmt.Println(stringSum("zj", "c"))

	//倒置切片
	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{1, 2, 3, 4, 5}))

	var invertSlice func([][]int) [][]int
	Bind(&invertSlice, InvertSlice)
	fmt.Println(invertSlice([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}))

	//总结要是使用refect初始化一个函数需要
	//1：先写一个func([]reflect.Value) []reflect.Value 类型的函数实现想要实现的功能,
	//fn func([]reflect.Value) []reflect.Value
	//2：写一个函数用来接收main函数声明的函数f ，传入函数的声明函数f一定是传&f，
	//然后通过reflect.ValueOf(p).Elem().Set(reflect.MakeFunc(reflect.ValueOf(p).Elem().Type(), fn))
	//将fn函数的功能赋给了f，此时f就可以传入参数了

}
