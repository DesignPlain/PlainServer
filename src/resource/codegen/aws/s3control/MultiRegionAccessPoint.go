package s3control

import types "libds/aws/types"

type MultiRegionAccessPoint struct {
	// The AWS account ID for the owner of the buckets for which you want to create a Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// A configuration block containing details about the Multi-Region Access Point. See Details Configuration Block below for more details
	Details types.S3control_MultiRegionAccessPointDetails `json:"details,omitempty" yaml:"details,omitempty"`
}
