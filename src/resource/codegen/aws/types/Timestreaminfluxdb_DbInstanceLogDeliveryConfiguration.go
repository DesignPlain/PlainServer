package types

type Timestreaminfluxdb_DbInstanceLogDeliveryConfiguration struct {
	// Configuration for S3 bucket log delivery.
	S3Configuration Timestreaminfluxdb_DbInstanceLogDeliveryConfigurationS3Configuration `json:"s3Configuration,omitempty" yaml:"s3Configuration,omitempty"`
}
