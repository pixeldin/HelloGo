package filetool

import (
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(path string) error {
	// 创建文件夹
	return os.Mkdir(path, os.ModePerm)
}

func MoveFile(src, tar string) error {
	return os.Rename(src, tar)
}

func GetFileNameFromUrl(url string) string {
	return filepath.Base(url)
}
