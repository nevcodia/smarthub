package domain

import (
	"net/url"
	"os"
)

type StorageObject struct {
	StoreName    string            `json:"store_name"`
	Name         string            `json:"name"`
	Path         string            `json:"path"`
	LastModified uint64            `json:"last_modified"`
	ETag         string            `json:"etag"`
	Size         uint64            `json:"size"`
	Metadata     map[string]string `json:"metadata"`
}

type StorageRepository interface {
	StoreNames() ([]string, error)
	Objects(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]StorageObject, error)
	ObjectsWithMetadata(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]StorageObject, error)
	GetObject(params ObjectParams) (StorageObject, error)
	Upload(params ObjectParams, metadata map[string]string, file *os.File) (StorageObject, error)
	PresignUploadLink(params ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error)
	Download(params ObjectParams) (DownloadFileResponse, error)
	PresignDownloadLink(params ObjectParams) (url.URL, error)
	PresignDownloadLinkWithDuration(params ObjectParams, duration uint) (url.URL, error)
	DeleteAll(storeName string, pathPrefix string) (bool, error)
	Delete(params ObjectParams) (bool, error)
	Copy(current ObjectParams, destination ObjectParams) (StorageObject, error)
	CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]StorageObject, error)
	Move(current ObjectParams, destination ObjectParams) (StorageObject, error)
}
