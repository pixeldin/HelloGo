package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//所有成员属性必须大写开头
type Address struct {
	Type     string
	City     string
	Conuntry string
}

type VCard struct {
	FirstName string
	LastName  string
	Add       []*Address
	Remark    string
}

func main() {
	pa := &Address{"A", "NewYork", "USA"}
	pb := &Address{"B", "BeiJing", "CNA"}
	vc := VCard{"Jim", "Black", []*Address{pa, pb}, "N"}

	//Encode
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
	// Equals to
	//对象序列号为字符串
	js, _ := json.Marshal(vc)
	fmt.Printf("Json format: %s\n", js)
	//jf, _ := os.OpenFile("JsVcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	//defer jf.Close()
	//writer := bufio.NewWriter(jf)
	//writer.Write(js)
	//writer.Flush()

	//Todo... Decode demo
	existFile, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_RDONLY, 0666)
	defer existFile.Close()
	existFile2, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_RDONLY, 0666)
	defer existFile2.Close()
	reader := bufio.NewReader(existFile)

	//	p是指针 obj是引用对象, p和&obj都是对象的地址,
	//  p := &obj
	dec2 := json.NewDecoder(reader)//reader
	var obj map[string]interface{}
	dec2.Decode(&obj)//p
	fmt.Println(obj)

	reader = bufio.NewReader(existFile2)
	dec := json.NewDecoder(reader)
	var dvc VCard
	dec.Decode(&dvc)
	fmt.Println(dvc)

}
