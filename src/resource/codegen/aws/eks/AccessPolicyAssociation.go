package eks

import types "libds/aws/types"

type AccessPolicyAssociation struct {
	// The configuration block to determine the scope of the access. See `access_scope` Block below.
	AccessScope types.Eks_AccessPolicyAssociationAccessScope `json:"accessScope,omitempty" yaml:"accessScope,omitempty"`

	// Name of the EKS Cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// The ARN of the access policy that you're associating.
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`

	// The IAM Principal ARN which requires Authentication access to the EKS cluster.
	PrincipalArn string `json:"principalArn,omitempty" yaml:"principalArn,omitempty"`
}
