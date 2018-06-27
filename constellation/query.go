package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
)

type (
	Message struct {
		Stauts  string `json:"status"`
		Msg    string   `json:"msg"`
		Result Result `json:"result"`
	}

	Result struct {
		Astroid   string  `json:"astroid"`
		Astroname string  `json:"astroname"`
		Year      Year    `json:"year"`
		Week      Week	   `json:"week"`
		Today     Today    `json:"today"`
		Tomorrow Tomorrow  `json:"tomorrow"`
		Month     Month  `json:"month"`
	}

	Year struct {
		Date    string  `json:"date"`
		Summary string `json:"summary"`
		Money   string `json:"money"`
		Career  string `json:"career"`
		Love    string `json:"love"`
	}

	Week struct {
		Date  string `json:"date"`
		Money  string `json:"money"`
		Career string `json:"career"`
		Love   string `json:"love"`
		Health string `json:"health"`
		Job    string `json:"job"`
	}

	Today struct {
		Date       string `json:"date"`
		Presummary string `json:"presummary"`
		Star       string `json:"star"`
		Color      string `json:"color"`
		Number     string `json:"number"`
		Summary    string `json:"summary"`
		Money      string `json:"money"`
		Career     string `json:"career"`
		Love       string `json:"love"`
		Health     string `json:"health"`
	}

	Tomorrow struct {
		Date       string `json:"date"`
		Presummary string `json:"presummary"`
		Star       string `json:"star"`
		Color      string `json:"color"`
		Number     string `json:"number"`
		Summary    string `json:"summary"`
		Money      string `json:"money"`
		Career     string `json:"career"`
		Love       string `json:"love"`
		Health     string `json:"health"`
	}

	Month struct {
		Date    string `json:"date"`
		Summary string `json:"summary"`
		Money   string `json:"money"`
		Career  string `json:"career"`
		Love    string `json:"love"`
	}
)

// todo https://mholt.github.io/json-to-go/ json 代码自动生成
func main() {
	//生成client 参数为默认
	client := &http.Client{}

	//生成要访问的url
	url := "http://jisuastro.market.alicloudapi.com/astro/fortune?astroid=3&date=2018-06-27"

	//提交请求
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "APPCODE 981573620df24ecb8a55495b02e96a48")

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(" client.Do(req) error: ", err)
	}

	//返回的状态码
	status := response.StatusCode
	if status != 200 {
		panic(errors.New("GET Fail"))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(response.Body) error ", err)
		return
	}

	var message Message

	err = json.Unmarshal([]byte(body), &message)

	fmt.Println(message)
	fmt.Println(err)

	fmt.Println("-------------------------------------------")

	m := map[string]interface{}{}
	err = json.Unmarshal([]byte(body), &m)

	fmt.Println(m["result"].(map[string]interface{})["year"].(map[string]interface{})["money"])
	fmt.Println(err)

}
