package domain

type StorageType string

const (
	S3         StorageType = "s3"
	FTP        StorageType = "ftp"
	SHAREPOINT StorageType = "sharepoint"
)
