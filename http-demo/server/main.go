package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xxx.html")
	if err == io.EOF {
		return
	}
	if err != nil {
		w.Write([]byte("滚"))
		fmt.Println("ioutil.ReadFile err=", err)
		return
	}
	w.Write(b)

}

func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求，参数都放在URL上（query param），请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL中的查询参数
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name,age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/mmb", f1)
	http.HandleFunc("/xxx/", f2)
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe err=", err)
		return
	}

}
