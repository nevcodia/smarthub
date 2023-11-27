package domain

type DownloadFileResponse struct {
	Filename    string
	Type        string
	Disposition string
}

type ErrorResponse struct {
	Message string `json:"message"`
}
