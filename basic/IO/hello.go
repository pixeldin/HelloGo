package main

import (
	"HelloGo/basic/IO/file"
	"fmt"
)

func main() {
	//newReader := bufio.NewReader(os.Stdin)

	//fmt.Println("Please input your name:")
	//readString, e := newReader.ReadString('\n')
	//if e != nil {
	//	fmt.Print("Err: %s\n", e)
	//} else {
	//	input := readString[:len(readString)-1]
	//	fmt.Printf("Hello, %s!\n", input)
	//}

	fmt.Println("============ Read file ============")
	file.ReadFileAndWriteWithBuffer("dd.txt", 10)

	//fp, _ := tools.GetCurrentDirectory()
	//fmt.Println(fp)

}
