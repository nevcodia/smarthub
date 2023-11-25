package service

import (
	"errors"
	"fmt"
	"github.com/nevcodia/smarthub/domain"
	"net/url"
	"os"
)

type SmartService interface {
	StoreNames(storeType domain.StorageType) ([]string, error)
	Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]domain.StorageObject, error)
	ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]domain.StorageObject, error)
	GetObject(storeType domain.StorageType, params domain.ObjectParams) (domain.StorageObject, error)
	Upload(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file *os.File) (domain.StorageObject, error)
	PresignUploadLink(storeType domain.StorageType, params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error)
	Download(storeType domain.StorageType, params domain.ObjectParams) (domain.DownloadFileResponse, error)
	PresignDownloadLink(storeType domain.StorageType, params domain.ObjectParams) (url.URL, error)
	PresignDownloadLinkWithDuration(storeType domain.StorageType, params domain.ObjectParams, duration uint) (url.URL, error)
	DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) (bool, error)
	Delete(storeType domain.StorageType, params domain.ObjectParams) (bool, error)
	Copy(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) (domain.StorageObject, error)
	CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error)
	Move(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) (domain.StorageObject, error)
}

type smartService struct {
	repos *map[domain.StorageType]domain.StorageRepository
}

func NewSmartService(repos *map[domain.StorageType]domain.StorageRepository) SmartService {
	return &smartService{
		repos: repos,
	}
}

func (s *smartService) StoreNames(storeType domain.StorageType) ([]string, error) {
	repository := (*s.repos)[storeType]
	if repository == nil {
		return nil, errors.New(fmt.Sprintf("%v is not supported", storeType.String()))
	}
	names, err := repository.StoreNames()
	if err != nil {
		return nil, err
	}
	return names, nil
}

func (s *smartService) Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) GetObject(storeType domain.StorageType, params domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) Upload(storeType domain.StorageType, params domain.ObjectParams, metadata map[string]string, file *os.File) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) PresignUploadLink(storeType domain.StorageType, params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) Download(storeType domain.StorageType, params domain.ObjectParams) (domain.DownloadFileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) PresignDownloadLink(storeType domain.StorageType, params domain.ObjectParams) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) PresignDownloadLinkWithDuration(storeType domain.StorageType, params domain.ObjectParams, duration uint) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) Delete(storeType domain.StorageType, params domain.ObjectParams) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) Copy(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *smartService) Move(storeType domain.StorageType, current domain.ObjectParams, destination domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}
