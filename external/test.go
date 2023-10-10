package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	//调用路由a

	log.Println("\n")
	log.Println("调用路由a")
	resp, err := http.Get("http://localhost:8099/a/get")
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("读取失败：", err)
		return
	}

	fmt.Println("响应结果：", string(body))

	//调用路由b
	log.Println("调用路由b")
	//接口参数为表单形式
	formValues := url.Values{}
	formValues.Set("param", "check post")
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8099/b/get", formBytesReader)
	if err != nil {
		// handle error
		log.Fatal("请求失败：", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//Do方法发送请求
	resp, err = client.Do(req)

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("读取失败：", err)
		return
	}
	fmt.Println("响应结果：", string(body))

}
