package models

import "time"

type Resource struct {
	PublicKey        *string          `json:"public_key"`
	Embedded         *ResourceList    `json:"_embedded"`
	Name             *string          `json:"name"`
	Preview          *string          `json:"preview"`
	Created          *time.Time       `json:"created"`
	Modified         *time.Time       `json:"modified"`
	CustomProperties CustomProperties `json:"custom_properties"`
	PublicUrl        *string          `json:"public_url"`
	Path             *string          `json:"path"`
	OriginPath       *string          `json:"origin_path"`
	Md5              *string          `json:"md5"`
	Type             *string          `json:"type"`
	MimeType         *string          `json:"mime_type"`
	Size             *int64           `json:"size"`
}

type ResourceList struct {
	Sort      *string    `json:"sort"`
	PublicKey *string    `json:"public_key"`
	Items     []Resource `json:"items"`
	Path      *string    `json:"path"`
	Limit     *int       `json:"limit"`
	Offset    *int       `json:"offset"`
	Total     *int       `json:"total"`
}

type FileResourceList struct {
	Items  []Resource `json:"items"`
	Limit  *int       `json:"limit"`
	Offset *int       `json:"offset"`
}

type LastUploadedResourceList struct {
	Items []Resource `json:"items"`
	Limit *int       `json:"limit"`
}