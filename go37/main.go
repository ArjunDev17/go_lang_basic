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
	byteData, err := json.Marshal(&wrapData)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println("Data is :", byteData)

	fmt.Println("-----------------------------")
	fmt.Println("Data is :", string(byteData))

	//if you want to conver marshel data into once again unmarshel
	err1 := json.Unmarshal(byteData, &wrapData)
	//here we are sending bytedata directly no need to convert data
	//because it is already iny byte array form
	if err != nil {
		fmt.Println("err :", err1)
	}
	fmt.Println("Data :", wrapData)
}
