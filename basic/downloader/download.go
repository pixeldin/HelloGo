package main

import (
	"HelloGo/basic/downloader/mdl"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

/*
	learn from https://github.com/monkeyWie/gopeed-core
*/
const (
	DOWNLOAD_URL = "http://127.0.0.1:8899/demo.zip"
	SAVE_PATH    = "D:\\测试下载\\download"
	CON          = 8
	RETRY_COUNT  = 5
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	// 是否支持断点续传/文件大小
	res, err := resolve(ctx, DOWNLOAD_URL)
	if err != nil {
		log.Error(err)
		return
	}

	// 创建文件, todo...创建目录
	file, err := touch(filepath.Join(SAVE_PATH, res.Files[0].Name), res.TotalSize)
	if err != nil {
		log.Error(err)
		return
	}

	var (
		// 切分文件块
		chunk []*mdl.Chunk
		// 切分块数
		chunkSize int
	)
	// 支持切分
	if res.Range {
		chunkSize = CON
		chunk = make([]*mdl.Chunk, chunkSize)
		partSize := res.TotalSize / int64(chunkSize)
		for i := 0; i < chunkSize; i++ {
			var (
				begin = partSize * int64(i)
				end   int64
			)
			if i == (chunkSize - 1) {
				end = res.TotalSize - 1
			} else {
				end = begin + partSize - 1
			}
			ck := mdl.NewChunk(begin, end)
			chunk[i] = ck
		}
	} else {
		chunkSize = 1
		// 单连接下载
		chunk = make([]*mdl.Chunk, chunkSize)
		chunk[0] = mdl.NewChunk(0, 0)
	}

	// 下载
	var doneCh = make(chan error, 1)
	fetch(ctx, res, file, chunk, chunkSize, doneCh)

	// wait for done
	if err := <-doneCh; err != nil {
		log.Printf("下载出错, err: %v", err)
		return
	}
	log.Println("下载完成。")
}

// resolve 解析下载链接
// 返回是否支持断点续传/文件大小
func resolve(ctx context.Context, reqURL string) (*mdl.Resource, error) {

	// todo... 根据协议构建不同下载解析器(http/ws/ftp)

	req, err := buildReq(ctx, reqURL)
	if err != nil {
		return nil, err
	}

	// 只访问一个字节，测试资源是否支持Range请求
	req.Header.Set(mdl.HttpHeaderRange, fmt.Sprintf(mdl.HttpHeaderRangeFormat, 0, 0))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	res := &mdl.Resource{
		Req:   reqURL,
		Range: false,
		Files: []*mdl.FileInfo{},
	}
	// 解析是否支持断点续传/文件大小
	if resp.StatusCode == mdl.HttpCodePartialContent {
		// 支持断点下载
		res.Range = true
		// 从Content-Range 获取大小
		conTotal := path.Base(resp.Header.Get(mdl.HttpHeaderContentRange))
		if conTotal != "" {
			parse, err := strconv.ParseInt(conTotal, 10, 64)
			if err != nil {
				return nil, err
			}
			res.TotalSize = parse
		}
	} else if resp.StatusCode == mdl.HttpCodeOK {
		// 200不支持断点下载, 从Content-Length获取大小
	} else {
		return nil, errors.New("Unknown file status")
	}
	// 解析文件名
	file := &mdl.FileInfo{
		Size: res.TotalSize,
	}
	// 从header获取文件名
	conDpsi := resp.Header.Get(mdl.HttpHeaderContentDisposition)
	if conDpsi != "" {
		_, params, _ := mime.ParseMediaType(conDpsi)
		filename := params["filename"]
		if filename != "" {
			file.Name = filename
		}
	}
	// 否则从url获取filename
	if file.Name == "" {
		file.Name = filepath.Base(DOWNLOAD_URL)
	}

	// 标注未知
	if file.Name == "" {
		file.Name = "unknown"
	}
	res.Files = append(res.Files, file)
	return res, nil
}

func buildReq(ctx context.Context, reqUrl string) (httpReq *http.Request, err error) {
	url, err := url.Parse(reqUrl)
	if err != nil {
		return
	}
	var (
		method string
		body   io.Reader
	)
	method = http.MethodGet
	if ctx != nil {
		httpReq, err = http.NewRequestWithContext(ctx, method, url.String(), body)
	} else {
		httpReq, err = http.NewRequest(method, url.String(), body)
	}
	if err != nil {
		return
	}
	return httpReq, nil
}

func touch(fileName string, size int64) (file *os.File, err error) {
	file, err = os.Create(fileName)
	if size > 0 {
		err := os.Truncate(fileName, size)
		if err != nil {
			return nil, err
		}
	}
	return
}

func fetch(ctx context.Context, res *mdl.Resource, file *os.File, chks []*mdl.Chunk, c int, doneCh chan error) {
	eg, _ := errgroup.WithContext(ctx)
	for i := 0; i < c; i++ {
		i := i
		eg.Go(func() error {
			return fetchChunk(ctx, res, file, i, chks)
		})
	}

	go func() {
		// error from errgroup
		err := eg.Wait()
		// 关闭文件
		file.Close()
		// 接收fetchChunk()内部状态
		doneCh <- err
	}()
	return
}

func fetchChunk(ctx context.Context, res *mdl.Resource, file *os.File, index int, chk []*mdl.Chunk) (err error) {
	ck := chk[index]
	req, err := buildReq(ctx, DOWNLOAD_URL)
	if err != nil {
		return err
	}
	var (
		client = http.DefaultClient
		buf    = make([]byte, 8192)
		//downloaded int64
	)
	/**************重试区间开始**************/
	// 根据是否分块下载设置header
	//	- 根据当前chunk设置文件区间到header
	//	- 判断请求返回status
	//		- 失败: 少于5次则做重试
	//		- 成功:
	//				根据offset, 把buf写入文件
	//				- 成功: return, 通知外部
	//				- 失败: 重试少于5次, 返回重试
	for i := 0; i < RETRY_COUNT; i++ {
		var resp *http.Response
		if res.Range {
			req.Header.Set(mdl.HttpHeaderRange,
				fmt.Sprintf(mdl.HttpHeaderRangeFormat, chk[index].Begin+ck.Downloaded, chk[index].End))
		} else {
			// 单连接重试没有断点续传
			ck.Downloaded = 0
		}

		// 获取字节区间
		if err := func() error {
			resp, err = client.Do(req)
			if err != nil {
				return err
			}
			if resp.StatusCode != mdl.HttpCodeOK && resp.StatusCode != mdl.HttpCodePartialContent {
				return errors.New(fmt.Sprintf("%d,%s", resp.StatusCode, resp.Status))
			}
			return nil
		}(); err != nil {
			continue
		}

		// 从body提取buf写入文件, 重置重试标识
		i = 0
		retry := false
		retry, err = func() (bool, error) {
			defer resp.Body.Close()
			for {
				n, err := resp.Body.Read(buf)
				if n > 0 {
					_, err := file.WriteAt(buf[:n], ck.Begin+ck.Downloaded)
					if err != nil {
						// 文件出错不重试
						return false, err
					}
					// 记录已下载, 用于断点续传
					ck.Downloaded += int64(n)
				}
				if err != nil {
					// err from read
					if err != io.EOF {
						return true, err
					}
					break
				}
			}
			// success exit
			return false, nil
		}()
		if !retry {
			break
		}
	}
	/**************重试区间结束**************/

	// 通知外部
	return
}
