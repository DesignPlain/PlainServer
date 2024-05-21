package types

type Cloudrunv2_ServiceTemplateVolumeGcs struct {
	// If true, mount the GCS bucket as read-only
	ReadOnly bool `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`

	// GCS Bucket name
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
