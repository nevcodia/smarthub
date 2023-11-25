package domain

type DownloadFileResponse struct {
	Content     []byte
	Type        string
	Disposition string
}
