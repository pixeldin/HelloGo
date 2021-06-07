package mdl

// 资源信息
type Resource struct {
	Req string
	// 资源总大小
	TotalSize int64
	// 是否支持断点下载
	Range bool
	// 资源所包含的文件列表
	Files []*FileInfo
}

type FileInfo struct {
	Name string
	Path string
	Size int64
}
