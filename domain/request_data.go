package domain

type PresignUploadRequest struct {
	StoreName      string            `json:"store_name"`
	Key            string            `json:"key"`
	MimeType       string            `json:"mime_type"`
	Metadata       map[string]string `json:"metadata"`
	ExpirationTime uint              `json:"exp"`
}

type ObjectMovementRequest struct {
	CurrentStoreName     string `json:"current_store_name"`
	CurrentKey           string `json:"current_key"`
	DestinationStoreName string `json:"destination_store_name"`
	DestinationKey       string `json:"destination_key"`
}
