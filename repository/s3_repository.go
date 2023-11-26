package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nevcodia/smarthub/domain"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"strings"
	"time"
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

func (s *s3Repository) Objects(storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]domain.StorageObject, error) {
	prefix = strings.TrimLeft(prefix, "/")
	response, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:  aws.String(storeName),
		Prefix:  &prefix,
		MaxKeys: &maxObjectsPerPage,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var storageObjects []domain.StorageObject
	for _, content := range response.Contents {
		storageObjects = append(storageObjects, domain.StorageObject{
			StoreName:    storeName,
			Key:          *content.Key,
			LastModified: (*content.LastModified).UnixMilli(),
			ETag:         *content.ETag,
			Size:         *content.Size,
		})
	}
	if storageObjects == nil {
		storageObjects = []domain.StorageObject{}
	}
	return storageObjects, nil
}

func (s *s3Repository) ObjectsWithMetadata(storeName string, maxObjectsPerPage int32, requestedPage int32, prefix string) ([]domain.StorageObject, error) {
	prefix = strings.TrimLeft(prefix, "/")
	response, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:  aws.String(storeName),
		Prefix:  &prefix,
		MaxKeys: &maxObjectsPerPage,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var storageObjects []domain.StorageObject
	for _, content := range response.Contents {
		metadataResponse, _ := s.client.HeadObject(context.TODO(), &s3.HeadObjectInput{
			Bucket: aws.String(storeName),
			Key:    content.Key,
		})
		storageObjects = append(storageObjects, domain.StorageObject{
			StoreName:    storeName,
			Key:          *content.Key,
			LastModified: (*content.LastModified).UnixMilli(),
			ETag:         *content.ETag,
			Size:         *content.Size,
			Metadata:     metadataResponse.Metadata,
		})
	}
	if storageObjects == nil {
		storageObjects = []domain.StorageObject{}
	}
	return storageObjects, nil
}

func (s *s3Repository) GetObject(params *domain.ObjectParams) (domain.StorageObject, error) {
	response, err := s.client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(params.StoreName),
		Key:    aws.String(params.Key),
	})
	if err != nil {
		log.Println(err)
		return domain.StorageObject{}, err
	}
	return domain.StorageObject{
		StoreName:    params.StoreName,
		Key:          params.Key,
		LastModified: (*response.LastModified).UnixMilli(),
		ETag:         *response.ETag,
		Size:         *response.ContentLength,
		Metadata:     response.Metadata,
	}, nil
}

func (s *s3Repository) UploadMultiPart(params *domain.ObjectParams, metadata map[string]string, fileHeader *multipart.FileHeader) (domain.StorageObject, error) {
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Error opening file:", err)
		return domain.StorageObject{}, err
	}
	defer file.Close()

	return s.Upload(params, metadata, file)
}

func (s *s3Repository) Upload(params *domain.ObjectParams, metadata map[string]string, file io.Reader) (domain.StorageObject, error) {
	response, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:   aws.String(params.StoreName),
		Key:      aws.String(params.Key),
		Metadata: metadata,
		Body:     file,
	})
	if err != nil {
		log.Printf("Couldn't upload file %v to %v. Here's why: %v\n",
			params.Key, params.StoreName, err)
		return domain.StorageObject{}, err
	}
	return domain.StorageObject{
		StoreName:    params.StoreName,
		Key:          params.Key,
		LastModified: time.Now().UnixMilli(),
		ETag:         *response.ETag,
		Metadata:     metadata,
	}, nil
}

func (s *s3Repository) PresignUploadLink(params *domain.ObjectParams, mimeType string, metadata map[string]string, duration uint) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Download(params *domain.ObjectParams) (domain.DownloadFileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignDownloadLink(params *domain.ObjectParams) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) PresignDownloadLinkWithDuration(params *domain.ObjectParams, duration uint) (url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) DeleteAll(storeName string, pathPrefix string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Delete(params *domain.ObjectParams) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Copy(current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Move(current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error) {
	//TODO implement me
	panic("implement me")
}
