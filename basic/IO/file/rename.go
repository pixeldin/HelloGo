package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func RenameFile(dir string) {
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		src := dir + "\\" + fileInfoList[i].Name()
		fmt.Println("源文件:", src) //打印当前文件或目录下的文件或目录名
		fullPath := handle(dir, fileInfoList[i].Name())
		fmt.Println("重命名:", fullPath) //打印当前文件或目录下的文件或目录名

		Rename(src, fullPath)
	}

}

func Rename(file string, target string) {
	err := os.Rename(file, target) //重命名 C:\log\2013.log 文件为install.txt
	if err != nil {
		//如果重命名文件失败,则输出错误 file rename Error!
		fmt.Println("file rename Error!")
		//打印错误详细信息
		fmt.Printf("%s", err)
	} else {
		//如果文件重命名成功,则输出 file rename OK!
		fmt.Println("file rename OK!")
	}
}

func handle(basic, src string) string {
	idx := strings.Index(src, "] ")
	src = src[idx+2:]
	return basic + "\\" + src
}
