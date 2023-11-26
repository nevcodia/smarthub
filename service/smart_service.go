package service

import (
	"errors"
	"fmt"
	"github.com/nevcodia/smarthub/domain"
	"io"
	"mime/multipart"
	"net/url"
)

type SmartService interface {
	StoreNames(storeType domain.StorageType) ([]string, error)
	Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]domain.StorageObject, error)
	ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]domain.StorageObject, error)
	GetObject(storeType domain.StorageType, params *domain.ObjectParams) (domain.StorageObject, error)
	UploadMultiPart(storeType domain.StorageType, params *domain.ObjectParams, metadata map[string]string, fileHeader *multipart.FileHeader) (domain.StorageObject, error)
	Upload(storeType domain.StorageType, params *domain.ObjectParams, metadata map[string]string, file io.Reader) (domain.StorageObject, error)
	PresignUploadLink(storeType domain.StorageType, params *domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error)
	Download(storeType domain.StorageType, params *domain.ObjectParams) (domain.DownloadFileResponse, error)
	PresignDownloadLink(storeType domain.StorageType, params *domain.ObjectParams) (url.URL, error)
	PresignDownloadLinkWithDuration(storeType domain.StorageType, params *domain.ObjectParams, duration uint) (url.URL, error)
	DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) (bool, error)
	Delete(storeType domain.StorageType, params *domain.ObjectParams) (bool, error)
	Copy(storeType domain.StorageType, current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error)
	CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error)
	Move(storeType domain.StorageType, current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error)
}

type smartService struct {
	repos map[domain.StorageType]domain.StorageRepository
}

func NewSmartService(repos map[domain.StorageType]domain.StorageRepository) SmartService {
	return &smartService{
		repos: repos,
	}
}

func (s *smartService) StoreNames(storeType domain.StorageType) ([]string, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return nil, err
	}
	return repository.StoreNames()
}

func (s *smartService) Objects(storeType domain.StorageType, storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return nil, err
	}
	return repository.Objects(storeName, maxObjectsPerPage, requestedPage, prefix)
}

func (s *smartService) ObjectsWithMetadata(storeType domain.StorageType, storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return nil, err
	}
	return repository.ObjectsWithMetadata(storeName, maxObjectsPerPage, requestedPage, prefix)
}

func (s *smartService) GetObject(storeType domain.StorageType, params *domain.ObjectParams) (domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return domain.StorageObject{}, err
	}
	return repository.GetObject(params)
}

func (s *smartService) UploadMultiPart(storeType domain.StorageType, params *domain.ObjectParams, metadata map[string]string, fileHeader *multipart.FileHeader) (domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return domain.StorageObject{}, err
	}
	return repository.UploadMultiPart(params, metadata, fileHeader)
}

func (s *smartService) Upload(storeType domain.StorageType, params *domain.ObjectParams, metadata map[string]string, file io.Reader) (domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return domain.StorageObject{}, err
	}
	return repository.Upload(params, metadata, file)
}

func (s *smartService) PresignUploadLink(storeType domain.StorageType, params *domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return url.URL{}, err
	}
	return repository.PresignUploadLink(params, mimeType, metadata, duration)
}

func (s *smartService) Download(storeType domain.StorageType, params *domain.ObjectParams) (domain.DownloadFileResponse, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return domain.DownloadFileResponse{}, err
	}
	return repository.Download(params)
}

func (s *smartService) PresignDownloadLink(storeType domain.StorageType, params *domain.ObjectParams) (url.URL, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return url.URL{}, err
	}
	return repository.PresignDownloadLink(params)
}

func (s *smartService) PresignDownloadLinkWithDuration(storeType domain.StorageType, params *domain.ObjectParams, duration uint) (url.URL, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return url.URL{}, err
	}
	return repository.PresignDownloadLinkWithDuration(params, duration)
}

func (s *smartService) DeleteAll(storeType domain.StorageType, storeName string, pathPrefix string) (bool, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return false, err
	}
	return repository.DeleteAll(storeName, pathPrefix)
}

func (s *smartService) Delete(storeType domain.StorageType, params *domain.ObjectParams) (bool, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return false, err
	}
	return repository.Delete(params)
}

func (s *smartService) Copy(storeType domain.StorageType, current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return domain.StorageObject{}, err
	}
	return repository.Copy(current, destination)
}

func (s *smartService) CopyAll(storeType domain.StorageType, sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return []domain.StorageObject{}, err
	}
	return repository.CopyAll(sourceStoreName, sourcePath, targetStoreName, targetPath)
}

func (s *smartService) Move(storeType domain.StorageType, current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error) {
	repository, err := s.GetRepository(storeType)
	if err != nil {
		return domain.StorageObject{}, err
	}
	return repository.Move(current, destination)
}

func (s *smartService) GetRepository(storeType domain.StorageType) (domain.StorageRepository, error) {
	repository := s.repos[storeType]
	if repository == nil {
		return nil, errors.New(fmt.Sprintf("%v is not supported", storeType.String()))
	}
	return repository, nil
}
