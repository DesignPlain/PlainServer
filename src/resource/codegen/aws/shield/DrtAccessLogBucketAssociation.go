package shield

import types "libds/aws/types"

type DrtAccessLogBucketAssociation struct {
	// The Amazon S3 bucket that contains the logs that you want to share.
	LogBucket string `json:"logBucket,omitempty" yaml:"logBucket,omitempty"`

	// The ID of the Role Arn association used for allowing Shield DRT Access.
	RoleArnAssociationId string `json:"roleArnAssociationId,omitempty" yaml:"roleArnAssociationId,omitempty"`

	//
	Timeouts types.Shield_DrtAccessLogBucketAssociationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
