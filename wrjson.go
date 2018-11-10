package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Host struct {
	IP   string
	Name string
}

//Wjson序列化+写入文件
func Wjson(m Host, filename string) {
	b, err := json.Marshal(m) //序列化 m    //结构体每一项必须要大写
	if err != nil {
		fmt.Println("marshal failed:", err)
		return
	}
	fmt.Println("json:", string(b)) //json: {"IP":"192.168.23.92","Name":"Sky"}

	//写入到文件
	if ioutil.WriteFile(filename, b, 0644) == nil { //将[]byte写入 文件   打开文件时先清空文件
		fmt.Println("写入文件成功")
	}
}

//读出+反序列化
func Rjson(filename string) Host {
	//从文件读出
	contents, err := ioutil.ReadFile(filename) //以[]byte 读出
	if err == nil {
		fmt.Println("读文件成功")
	}

	mmm := Host{}
	err = json.Unmarshal(contents, &mmm) //反序列化到 mmm
	if err == nil {
		fmt.Println("反序列化成功")
	}
	return mmm
}

func main() {
	m := Host{Name: "Sky", IP: "192.168.23.99"}
	filename := "strjson.json"
	Wjson(m, filename)
	mmm := Rjson(filename)
	fmt.Println(mmm)      //{192.168.23.99 Sky}
	fmt.Println(mmm.Name) //Sky
}

//f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)  写和追加模式打开文件
