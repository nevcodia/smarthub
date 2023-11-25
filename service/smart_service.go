package service

import (
	"github.com/nevcodia/smarthub/domain"
	"net/url"
	"os"
)

type SmartService interface {
	StoreNames() []string
	Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject
	ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject
	GetObject(storeType domain.StorageType, params domain.ObjectParams) domain.StorageObject
	Upload(storeType domain.StorageType, params domain.ObjectParams, file *os.File) domain.StorageObject
	UploadWithMetadata(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file *os.File) domain.StorageObject
	UploadAsByteArray(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file []byte, size uint, mimeType string) domain.StorageObject
	PresignUploadLink(storeType domain.StorageType, params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) url.URL
	Download(storeType domain.StorageType, params domain.ObjectParams) domain.DownloadFileResponse
	PresignDownloadLink(storeType domain.StorageType, params domain.ObjectParams) url.URL
	PresignDownloadLinkWithDuration(storeType domain.StorageType, params domain.ObjectParams, duration uint) url.URL
	DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) bool
	Delete(storeType domain.StorageType, params domain.ObjectParams) bool
	Copy(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject
	CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) []domain.StorageObject
	Move(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject
}

type smartService struct {
	repos *map[domain.StorageType]domain.StorageRepository
}

func NewSmartService(repos *map[domain.StorageType]domain.StorageRepository) SmartService {
	return &smartService{
		repos: repos,
	}
}

func (s smartService) StoreNames() []string {
	//TODO implement me
	panic("implement me")
}

func (s smartService) Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) GetObject(storeType domain.StorageType, params domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) Upload(storeType domain.StorageType, params domain.ObjectParams, file *os.File) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) UploadWithMetadata(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file *os.File) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) UploadAsByteArray(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file []byte, size uint, mimeType string) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) PresignUploadLink(storeType domain.StorageType, params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s smartService) Download(storeType domain.StorageType, params domain.ObjectParams) domain.DownloadFileResponse {
	//TODO implement me
	panic("implement me")
}

func (s smartService) PresignDownloadLink(storeType domain.StorageType, params domain.ObjectParams) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s smartService) PresignDownloadLinkWithDuration(storeType domain.StorageType, params domain.ObjectParams, duration uint) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s smartService) DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) bool {
	//TODO implement me
	panic("implement me")
}

func (s smartService) Delete(storeType domain.StorageType, params domain.ObjectParams) bool {
	//TODO implement me
	panic("implement me")
}

func (s smartService) Copy(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartService) Move(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}
