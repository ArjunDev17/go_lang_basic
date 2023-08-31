package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string
	Designation string
	Phone       int64
}

func UsingStruct() {
	var wrapData User
	data := `{"Name":"StructArj","Designation":"Dev","Pjone":9451915742}`
	err := json.Unmarshal([]byte(data), &wrapData)
	if err != nil {
		fmt.Println("err :", err)
	}
	fmt.Println("StuctData :", wrapData)
	//with the help of srtuct by key we can easily access the data
	//example
	fmt.Println("Name :", wrapData.Name)
}
func usingOkSyntax() {
	var wrapData map[string]string
	data := `{"name":"Ram","designation":"God"}`
	err := json.Unmarshal([]byte(data), &wrapData)
	if err != nil {
		fmt.Println("err:", err)
	}
	if wrapData["name"] == "ram" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	// val, ok := wrapData["xyz"]
	// if ok {
	// 	fmt.Println(val)
	// } else {
	// 	fmt.Println("Not Present key")
	// }

}
func main() {
	fmt.Println("Jai Shree Ram")
	var wrapData map[string]string
	data := `{"name":"Ram","designation":"God"}`
	err := json.Unmarshal([]byte(data), &wrapData)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Data :", wrapData)
	fmt.Println("-----------------------------")
	TakingAlltypeValue()
	fmt.Println("-----------------------------")
	UsingStruct()
	fmt.Println("-----------------------------")
	usingOkSyntax()
}
func TakingAlltypeValue() {
	var wrapData map[string]interface{}
	//we putted here interface which is able to take all types data
	data := `{"name":"Ram","designation":"Creator","Phone num":9451915742}`
	err := json.Unmarshal([]byte(data), &wrapData)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Data :", wrapData)
	//in map by using key printing the value
	fmt.Println("Nmae :", wrapData["Name"])

}
