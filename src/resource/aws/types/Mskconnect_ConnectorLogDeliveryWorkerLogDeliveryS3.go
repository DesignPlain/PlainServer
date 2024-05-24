package types

type Mskconnect_ConnectorLogDeliveryWorkerLogDeliveryS3 struct {
	// The name of the S3 bucket that is the destination for log delivery.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Whether log delivery to Amazon CloudWatch Logs is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The S3 prefix that is the destination for log delivery.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
