package types

type S3_BucketMetricFilter struct {
	// S3 Access Point ARN for filtering (singular).
	AccessPoint string `json:"accessPoint,omitempty" yaml:"accessPoint,omitempty"`

	// Object prefix for filtering (singular).
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// Object tags for filtering (up to 10).
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
