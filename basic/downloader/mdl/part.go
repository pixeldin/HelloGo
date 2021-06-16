package mdl

type Chunk struct {
	Status Status
	Begin  int64
	End    int64
	// 已下载字节数量, 后续更新, 用于断点续传
	Downloaded int64
}

func NewChunk(begin int64, end int64) *Chunk {
	return &Chunk{
		Status: DownloadStatusReady,
		Begin:  begin,
		End:    end,
	}
}

type Extra struct {
	Method string
	Header map[string]string
	Body   string
}
