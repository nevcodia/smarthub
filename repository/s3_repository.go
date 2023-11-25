package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nevcodia/smarthub/domain"
	"net/url"
	"os"
)

type s3Repository struct {
	client *s3.Client
}

func NewS3Repository(client *s3.Client) domain.StorageRepository {
	return &s3Repository{
		client: client,
	}
}

func (s *s3Repository) StoreNames() []string {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Objects(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) ObjectsWithMetadata(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) GetObject(params domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Upload(params domain.ObjectParams, metadata map[string]string, file *os.File) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignUploadLink(params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Download(params domain.ObjectParams) domain.DownloadFileResponse {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignDownloadLink(params domain.ObjectParams) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignDownloadLinkWithDuration(params domain.ObjectParams, duration uint) url.URL {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) DeleteAll(storeName string, pathPrefix string) bool {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Delete(params domain.ObjectParams) bool {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Copy(current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) []domain.StorageObject {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Move(current domain.ObjectParams, destination domain.ObjectParams) domain.StorageObject {
	//TODO implement me
	panic("implement me")
}
