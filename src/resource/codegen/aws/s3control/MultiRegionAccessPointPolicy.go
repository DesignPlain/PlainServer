package s3control

import types "libds/aws/types"

type MultiRegionAccessPointPolicy struct {
	// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details types.S3control_MultiRegionAccessPointPolicyDetails `json:"details,omitempty" yaml:"details,omitempty"`
}
