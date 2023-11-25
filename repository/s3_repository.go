package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nevcodia/smarthub/domain"
	"log"
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

func (s *s3Repository) StoreNames() ([]string, error) {
	buckets, err := s.client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Println("Couldn't list buckets for your account. Here's why: %v\n", err)
		return nil, err
	}
	var bucketNames []string
	for _, bucket := range buckets.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}
	return bucketNames, nil
}

func (s *s3Repository) Objects(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) ObjectsWithMetadata(storeName string, maxObjectsPerPage uint, requestedPage uint, prefix string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) GetObject(params domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Upload(params domain.ObjectParams, metadata map[string]string, file *os.File) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignUploadLink(params domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Download(params domain.ObjectParams) (domain.DownloadFileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignDownloadLink(params domain.ObjectParams) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignDownloadLinkWithDuration(params domain.ObjectParams, duration uint) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) DeleteAll(storeName string, pathPrefix string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Delete(params domain.ObjectParams) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Copy(current domain.ObjectParams, destination domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Move(current domain.ObjectParams, destination domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}
