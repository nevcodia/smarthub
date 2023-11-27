package domain

type DownloadFileResponse struct {
	Content     []byte
	Type        string
	Disposition string
}

type ErrorResponse struct {
	Message string `json:"message"`
}
