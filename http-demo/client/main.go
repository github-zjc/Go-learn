package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
//长连接   适合连接次数频繁的条件下
var ( 
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	// resp,err := http.Get("http://127.0.0.1:9000/xxx?name=sb&age=18")
	// if err != nil {
	// 	fmt.Println("http.Get err=",err)
	// 	return
	// }

	//自定义一个request请求
	data := url.Values{} //url encode （url编码，，作用是转义）
	urlObj, err := url.Parse("http://127.0.0.1:9000/xxx/")
	if err != nil {
		fmt.Println("url.Parse err=", err)
		return
	}
	data.Set("name", "张静超")
	data.Set("age", "9000")
	queryStr := data.Encode() //	URl encode 之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("Get", urlObj.String(), nil)
	if err != nil {
		fmt.Println("http.NewRequest err = ", err)
		return
	}
	// resp, err := http.DefaultClient.Do(req) //发请求
	// if err != nil {
	// 	fmt.Println("http.DefaultClient.Do err=",err)
	// }

	//短连接  适合请求次数不频繁的场景
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }
	// client := http.Client{
	// 	Transport: tr,
	// }

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do err=", err)
	}

	//从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll err=", err)
	}
	fmt.Println(string(b))
}
