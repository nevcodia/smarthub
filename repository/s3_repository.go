package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nevcodia/smarthub/domain"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type s3Repository struct {
	client        *s3.Client
	presignClient *s3.PresignClient
}

func NewS3Repository(client *s3.Client) domain.StorageRepository {
	presignClient := s3.NewPresignClient(client)
	return &s3Repository{
		client:        client,
		presignClient: presignClient,
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
		log.Printf("Couldn't get objects from %v. Here's why: %v\n", storeName, err)
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
		log.Printf("Couldn't get objects from %v. Here's why: %v\n", storeName, err)
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
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", params.StoreName, params.Key, err)
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

func (s *s3Repository) PresignUploadLink(params *domain.ObjectParams, mimeType string, metadata map[string]string, exp uint) (string, error) {
	request, err := s.presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(params.StoreName),
		Key:         aws.String(params.Key),
		Metadata:    metadata,
		ContentType: &mimeType,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(exp * uint(time.Millisecond))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n",
			params.StoreName, params.Key, err)
		return "", err
	}
	return request.URL, err
}

func (s *s3Repository) Download(params *domain.ObjectParams) (domain.DownloadFileResponse, error) {
	result, err := s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(params.StoreName),
		Key:    aws.String(params.Key),
	})
	filename := path.Base(params.Key)
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", params.StoreName, params.Key, err)
		return domain.DownloadFileResponse{}, err
	}
	defer result.Body.Close()
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", filename, err)
		return domain.DownloadFileResponse{}, err
	}
	defer file.Close()
	_, err = io.Copy(file, result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", params.Key, err)
	}
	return domain.DownloadFileResponse{
		Filename:    filename,
		Type:        *result.ContentType,
		Disposition: "inline;filename=" + filename,
	}, err
}

func (s *s3Repository) PresignDownloadLink(params *domain.ObjectParams) (string, error) {
	return s.PresignDownloadLinkWithExpTime(params, 15*uint(time.Minute)) //Default time 15 minute
}

func (s *s3Repository) PresignDownloadLinkWithExpTime(params *domain.ObjectParams, exp uint) (string, error) {
	request, err := s.presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(params.StoreName),
		Key:    aws.String(params.Key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(exp * uint(time.Millisecond))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			params.StoreName, params.Key, err)
		return "", err
	}
	return request.URL, err
}

func (s *s3Repository) DeleteAll(storeName string, pathPrefix string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *s3Repository) Delete(params *domain.ObjectParams) (bool, error) {
	response, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(params.StoreName),
		Key:    aws.String(params.Key),
	})
	if err != nil {
		log.Printf("Couldn't delete %v:%v. Here's why: %v\n",
			params.StoreName, params.Key, err)
		return false, err
	}
	return *response.DeleteMarker, nil
}

func (s *s3Repository) Copy(current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error) {
	_, err := s.client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		Bucket:     aws.String(destination.StoreName),
		Key:        aws.String(destination.Key),
		CopySource: aws.String(fmt.Sprintf("%v/%v", current.StoreName, strings.TrimLeft(current.Key, "/"))),
	})
	if err != nil {
		log.Printf("Couldn't copy object from %v:%v to %v:%v. Here's why: %v\n",
			current.StoreName, current.Key, destination.StoreName, destination.Key, err)
		return domain.StorageObject{}, err
	}
	return domain.StorageObject{
		StoreName:    destination.StoreName,
		Key:          destination.Key,
		LastModified: time.Now().UnixMilli(),
	}, nil
}

func (s *s3Repository) CopyAll(sourceStoreName string, sourcePath string, targetStoreName string, targetPath string) ([]domain.StorageObject, error) {
	panic("implement me")
}

func (s *s3Repository) Move(current *domain.ObjectParams, destination *domain.ObjectParams) (domain.StorageObject, error) {
	_, err := s.Copy(current, destination)
	_, err = s.Delete(current)
	if err != nil {
		log.Printf("Couldn't move object from %v:%v to %v:%v. Here's why: %v\n",
			current.StoreName, current.Key, destination.StoreName, destination.Key, err)
		return domain.StorageObject{}, err
	}
	return domain.StorageObject{
		StoreName:    destination.StoreName,
		Key:          destination.Key,
		LastModified: time.Now().UnixMilli(),
	}, nil
}
