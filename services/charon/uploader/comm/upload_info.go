package comm

type UploadInfo struct {
	Url        string
	UploadType int
}

const (
	UploadTypeLocal = iota
	UploadTypeMysql
	UploadType
)
