package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) {
	filehandler, e := os.Open(path)
	if e != nil {
		fmt.Printf("Open file err: %s", e)
		return
	}
	reader := bufio.NewReader(filehandler)

	for {
		s, e := reader.ReadString('\n')
		fmt.Printf("File data: %s", s)
		//var databyte []byte = []byte(s)
		//ioutil.WriteFile("copy_dd.txt", databyte, 0644)

		if e == io.EOF {
			return
		}
	}
}

func ReadFileAndWriteWithBuffer(path string, bufSize int) {
	//打开文件句柄
	filehandler, e := os.Open(path)
	if e != nil {
		fmt.Printf("Open file err: %s", e)
		return
	}
	defer filehandler.Close()

	//嵌套一层buffio
	reader := bufio.NewReader(filehandler)

	//Try to output into file
	//ioutil.WriteFile("copy_dd.txt", buf, 0644)
	//打开写入文件目标句柄
	//outputFile, outputError := os.OpenFile("copy_dd.txt", os.O_WRONLY|os.O_CREATE, 0666)
	//写入前truncate 文件
	outputFile, outputError := os.OpenFile("copy_dd.txt", os.O_TRUNC|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation, err: %s\n", outputError)
		return
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	for {
		buf := make([]byte, bufSize)
		//循环按照bufsize大小读取文件
		n, _ := reader.Read(buf)
		if n == 0 {
			break
		}
		if n <= bufSize {
			buf = buf[0:n]
		}
		//handle file segment
		fmt.Println("file data: " + string(buf))

		//依次写入固定目标句柄
		writer.Write(buf)
		writer.Flush()

	}
	//for {
	//	s, e := reader.ReadString('\n')
	//	fmt.Printf("File data: %s", s)
	//	//var databyte []byte = []byte(s)
	//ioutil.WriteFile("copy_dd.txt", databyte, 0644)
	//
	//	if e == io.EOF {
	//		return
	//	}
	//}

}
