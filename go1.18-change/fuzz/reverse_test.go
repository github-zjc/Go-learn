package main

import (
	"testing"
	"unicode/utf8"
)
//go test 指令使用 f.Add(tc) 添加到测试的饲料库作为测试的内容-----在不进行模糊测试的情况下运行模糊测试，以确保种子输入通过。
//go test -fuzz=Fuzz 会进行模糊测试产生一些符合规定的测试数据进行测试 
//！！！这样知道遇到失败才会退出，并在当前目录下自动创建/testdata/fuzz/Fuzzxxx里面会有测试失败的数据内容
//go test -fuzz=Fuzz -fuzztime 30s 可以做测试时间的限制，如果没有发现错误，他会在退出前模糊测试30秒
//go test -fuzz=Fuzzxxxxxxxxxxxxxxxxxxx  可以指定测试失败的数据作为测试内容
func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
