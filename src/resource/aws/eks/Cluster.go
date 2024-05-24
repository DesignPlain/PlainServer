package eks

import types "DesignSphere_Server/src/resource/aws/types"

type Cluster struct {
	//
	DefaultAddonsToRemoves []string `json:"defaultAddonsToRemoves,omitempty" yaml:"defaultAddonsToRemoves,omitempty"`

	// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
	EncryptionConfig types.Eks_ClusterEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	// Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn't available for creating Amazon EKS clusters on the AWS cloud.
	OutpostConfig types.Eks_ClusterOutpostConfig `json:"outpostConfig,omitempty" yaml:"outpostConfig,omitempty"`

	// ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.

	   The following arguments are optional:
	*/
	VpcConfig types.Eks_ClusterVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Configuration block for the access config associated with your cluster, see [Amazon EKS Access Entries](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html).
	AccessConfig types.Eks_ClusterAccessConfig `json:"accessConfig,omitempty" yaml:"accessConfig,omitempty"`

	// Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
	KubernetesNetworkConfig types.Eks_ClusterKubernetesNetworkConfig `json:"kubernetesNetworkConfig,omitempty" yaml:"kubernetesNetworkConfig,omitempty"`

	// Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]-$`).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
	EnabledClusterLogTypes []string `json:"enabledClusterLogTypes,omitempty" yaml:"enabledClusterLogTypes,omitempty"`
}
