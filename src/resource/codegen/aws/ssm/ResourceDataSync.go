package ssm

import types "libds/aws/types"

type ResourceDataSync struct {
	// Name for the configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Amazon S3 configuration details for the sync.
	S3Destination types.Ssm_ResourceDataSyncS3Destination `json:"s3Destination,omitempty" yaml:"s3Destination,omitempty"`
}
