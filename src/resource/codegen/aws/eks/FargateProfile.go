package eks

import types "libds/aws/types"

type FargateProfile struct {
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors []types.Eks_FargateProfileSelector `json:"selectors,omitempty" yaml:"selectors,omitempty"`

	/*
	   Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).

	   The following arguments are optional:
	*/
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Name of the EKS Cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Name of the EKS Fargate Profile.
	FargateProfileName string `json:"fargateProfileName,omitempty" yaml:"fargateProfileName,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn string `json:"podExecutionRoleArn,omitempty" yaml:"podExecutionRoleArn,omitempty"`
}
