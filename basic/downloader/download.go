package main

import (
	"HelloGo/basic/downloader/mdl"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// 获取文件大小

	// 切分

	// 多协程下载
}

// resolve 解析下载链接
// 返回是否支持断点续传/文件大小
func resolve(reqURL string) (*mdl.Resource, error) {

	// todo... 根据协议构建不同下载解析器(http/ws/ftp)

	req, err := buildReq(nil, reqURL)
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
	// todo...解析断点续传/文件大小
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
	//body = ioutil.NopCloser(bytes.NewBufferString(extra.Body))
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
