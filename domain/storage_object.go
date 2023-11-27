package domain

import (
	"io"
	"mime/multipart"
)

type StorageObject struct {
	StoreName    string            `json:"store_name"`
	Key          string            `json:"key"`
	LastModified int64             `json:"last_modified,omitempty"`
	ETag         string            `json:"etag,omitempty"`
	Size         int64             `json:"size,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type StorageRepository interface {
	StoreNames() ([]string, error)
	Objects(storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]StorageObject, error)
	ObjectsWithMetadata(storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]StorageObject, error)
	GetObject(params *ObjectParams) (StorageObject, error)
	Upload(params *ObjectParams, metadata map[string]string, file io.Reader) (StorageObject, error)
	UploadMultiPart(params *ObjectParams, metadata map[string]string, fileHeader *multipart.FileHeader) (StorageObject, error)
	PresignUploadLink(params *ObjectParams, mimeType string, metadata map[string]string, exp uint) (string, error)
	Download(params *ObjectParams) (DownloadFileResponse, error)
	PresignDownloadLink(params *ObjectParams) (string, error)
	PresignDownloadLinkWithExpTime(params *ObjectParams, exp uint) (string, error)
	DeleteAll(storeName string, pathPrefix string) (bool, error)
	Delete(params *ObjectParams) (bool, error)
	Copy(current *ObjectParams, destination *ObjectParams) (StorageObject, error)
	CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]StorageObject, error)
	Move(current *ObjectParams, destination *ObjectParams) (StorageObject, error)
}
