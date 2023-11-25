package domain

type StorageType string

const (
	S3         StorageType = "s3"
	FTP        StorageType = "ftp"
	SHAREPOINT StorageType = "sharepoint"
)

func (s StorageType) String() string {
	return string(s)
}

func StorageTypeFromValue(v string) StorageType {
	switch v {
	case "s3":
		return S3
	case "ftp":
		return FTP
	case "sharepoint":
		return SHAREPOINT
	default:
		return "unknown"
	}
}
