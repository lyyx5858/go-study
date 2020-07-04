package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"

	//"io"
	"net/http"
	//"os"
)

func main() {
	res, err := http.Get("http://www.baidu.com") // 请求会把返回的内容，res的body中

	if err != nil {
		fmt.Println(err)
		panic("请求出错")
	}

	/**从res读取内容**/
	defer res.Body.Close() // 读取结束后关闭流，类似文件操作完后，需要关闭一样

	fmt.Println("响应状态:", res.StatusCode,"\n") // 响应状态
	fmt.Println("响应头", res.Header, "\n")       // 响应头

	/*读取内容*/
	reqBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(reqBody))

}