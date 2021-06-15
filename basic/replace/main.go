package main

import (
	"HelloGo/basic/replace/filetool"
	"HelloGo/basic/replace/util"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	SAVE_DIR             = "D:\\测试下载\\download"
	DOWN_LOAD_LINK       = "http://127.0.0.1:8899/demo.zip"
	CLIENT_PATH          = "C:\\JX3HDBeta\\bin\\zhcn_hd\\bin64\\"
	EXE_NAME             = "JX3ClientX64.exe"
	BK_EXE_NAME          = "JX3ClientX64_bk.exe"
	HOUR_SECONDS         = 3600
	MIN_SECONDS          = 60
	DOWNLOAD_TIMEOUT_SES = 20
	RETRY_DOWNLOAD_TIMES = 5
)

func main() {
	util.PrintInfo("========可执行包更新常驻进程========")
	for {
		initDir()

		//download()
		downloadWithRetry()

		move()

		// 6小时替换一次
		time.Sleep(6 * HOUR_SECONDS * time.Second)
	}
}

func initDir() {
	if exist, err := filetool.PathExists(SAVE_DIR); !exist {
		util.PrintInfo("创建目录: %s", SAVE_DIR)
		if err != nil {
			panic(err)
		}
		if err := filetool.CreateDir(SAVE_DIR); err != nil {
			panic(err)
		}
	}
}

// 异步下载，增加超时重试
func downloadWithRetry() {
	for i := 0; i < RETRY_DOWNLOAD_TIMES; i++ {
		errCh, ch := downloadWithCh(DOWNLOAD_TIMEOUT_SES)
		select {
		case err := <-errCh:
			util.PrintError("下载出错, 重试%d次, err: %v", i+1, err)
			continue
		case <-time.After(DOWNLOAD_TIMEOUT_SES * time.Second):
			util.PrintError("下载超时, 重试%d次。", i+1)
			continue
		case _ = <-ch:
			util.PrintInfo("下载完成。")
			return
		}
	}
}

// 同步下载
func download() error {
	util.PrintInfo("开始下载: %s", DOWN_LOAD_LINK)
	res, err := http.Get(DOWN_LOAD_LINK)
	if err != nil {
		panic(err)
		//return err
	}
	f, err := os.Create(filepath.Join(SAVE_DIR, filetool.GetFileNameFromUrl(DOWN_LOAD_LINK)))
	defer func() {
		res.Body.Close()
		f.Close()
	}()

	if err != nil {
		//return err
		panic(err)
	}
	io.Copy(f, res.Body)
	util.PrintInfo("下载完成: %s ", DOWN_LOAD_LINK)
	return nil
}

// 异步下载
func downloadWithCh(i int) (chan error, chan struct{}) {
	errCh := make(chan error)
	retCh := make(chan struct{})
	go func() {
		util.PrintInfo("开始下载: %s", DOWN_LOAD_LINK)

		var err error
		var res = new(http.Response)
		var f = new(os.File)
		//if i == 0 {
		//	time.Sleep(20 * time.Second)
		//} else if i <= 3 {
		//	errCh <- errors.New("mock err")
		//	return
		//}
		res, err = http.Get(DOWN_LOAD_LINK)
		defer res.Body.Close()
		if err != nil {
			errCh <- err
			f.Close()
		}
		f, err = os.Create(filepath.Join(SAVE_DIR, filetool.GetFileNameFromUrl(DOWN_LOAD_LINK)))
		//defer func() {
		//	if cerr := f.Close(); cerr != nil {
		//		util.PrintError("Close file err: %v", cerr)
		//	}
		//}()

		if err != nil {
			errCh <- err
			f.Close()
			//return
		}
		io.Copy(f, res.Body)
		util.PrintInfo("下载完成: %s ", DOWN_LOAD_LINK)
		//f.Sync()
		f.Close()
		retCh <- struct{}{}
		//break
	}()
	return errCh, retCh
}

func move() {
	util.PrintInfo("更新可执行exe...")
	tarFile := CLIENT_PATH + EXE_NAME
	bakFile := CLIENT_PATH + BK_EXE_NAME
	if exist, _ := filetool.PathExists(tarFile); exist {
		util.PrintInfo("备份可执行exe -> %s", bakFile)
		if err := filetool.MoveFile(CLIENT_PATH+EXE_NAME, bakFile); err != nil {
			panic(err)
		}
	}

	for i := 0; i < RETRY_DOWNLOAD_TIMES; i++ {
		if err := filetool.MoveFile(SAVE_DIR+EXE_NAME, tarFile); err != nil {
			util.PrintError("Rename err: %v, retry %d time.", err, i+1)
			time.Sleep(3 * time.Second)
			continue
			//panic(err)
		}
		break
	}
	util.PrintInfo("更新替换完成。")
}
