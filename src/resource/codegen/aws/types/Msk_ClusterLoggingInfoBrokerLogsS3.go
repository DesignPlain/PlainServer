package types

type Msk_ClusterLoggingInfoBrokerLogsS3 struct {
	// Name of the S3 bucket to deliver logs to.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Prefix to append to the folder name.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
