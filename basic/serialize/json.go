package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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
	fmt.Printf("Json format: %s", js)
	//jf, _ := os.OpenFile("JsVcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	//defer jf.Close()
	//writer := bufio.NewWriter(jf)
	//writer.Write(js)
	//writer.Flush()

	//Todo... Decode demo
	existFile, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_RDONLY, 0666)
	defer existFile.Close()
	dec := json.NewDecoder(existFile)
	var dvc Address
	dec.Decode(dvc)
	fmt.Println(dvc)
}
