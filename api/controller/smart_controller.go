package controller

import (
	"github.com/nevcodia/smarthub/domain"
	"github.com/nevcodia/smarthub/service"
	"net/url"
	"os"
)

type SmartController interface {
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

type smartController struct {
	service service.SmartService
}

func NewSmartController(service service.SmartService) SmartController {
	return &smartController{
		service: service,
	}
}

func (s smartController) StoreNames() []string {
	//TODO implement me
	panic("implement me")
}

func (s smartController) Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) GetObject(storeType domain.StorageType, params domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) Upload(storeType domain.StorageType, params domain.ObjectParams, file *os.File) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) UploadWithMetadata(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file *os.File) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) UploadAsByteArray(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file []byte, size uint, mimeType string) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) PresignUploadLink(storeType domain.StorageType, params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s smartController) Download(storeType domain.StorageType, params domain.ObjectParams) domain.DownloadFileResponse {
	//TODO implement me
	panic("implement me")
}

func (s smartController) PresignDownloadLink(storeType domain.StorageType, params domain.ObjectParams) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s smartController) PresignDownloadLinkWithDuration(storeType domain.StorageType, params domain.ObjectParams, duration uint) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s smartController) DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) bool {
	//TODO implement me
	panic("implement me")
}

func (s smartController) Delete(storeType domain.StorageType, params domain.ObjectParams) bool {
	//TODO implement me
	panic("implement me")
}

func (s smartController) Copy(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s smartController) Move(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}
