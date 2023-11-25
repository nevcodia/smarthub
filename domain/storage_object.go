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
	StoreNames() []string
	Objects(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []StorageObject
	ObjectsWithMetadata(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []StorageObject
	GetObject(params ObjectParams) StorageObject
	Upload(params ObjectParams, metadata map[string]string, file *os.File) StorageObject
	PresignUploadLink(params ObjectParams, mimeType string, metadata map[string]string, duration uint) url.URL
	Download(params ObjectParams) DownloadFileResponse
	PresignDownloadLink(params ObjectParams) url.URL
	PresignDownloadLinkWithDuration(params ObjectParams, duration uint) url.URL
	DeleteAll(storeName string, pathPrefix string) bool
	Delete(params ObjectParams) bool
	Copy(current ObjectParams, destination ObjectParams) StorageObject
	CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) []StorageObject
	Move(current ObjectParams, destination ObjectParams) StorageObject
}
