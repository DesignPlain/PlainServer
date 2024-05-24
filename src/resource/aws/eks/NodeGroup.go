package eks

import types "DesignSphere_Server/src/resource/aws/types"

type NodeGroup struct {
	// Type of capacity associated with the EKS Node Group. Valid values: `ON_DEMAND`, `SPOT`. This provider will only perform drift detection if a configuration value is provided.
	CapacityType string `json:"capacityType,omitempty" yaml:"capacityType,omitempty"`

	// List of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. The provider will only perform drift detection if a configuration value is provided.
	InstanceTypes []string `json:"instanceTypes,omitempty" yaml:"instanceTypes,omitempty"`

	// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Identifiers of EC2 Subnets to associate with the EKS Node Group.

	   The following arguments are optional:
	*/
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Kubernetes version. Defaults to EKS Cluster Kubernetes version. The provider will only perform drift detection if a configuration value is provided.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. See the [AWS documentation](https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType) for valid values. This provider will only perform drift detection if a configuration value is provided.
	AmiType string `json:"amiType,omitempty" yaml:"amiType,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `node_group_name`.
	NodeGroupNamePrefix string `json:"nodeGroupNamePrefix,omitempty" yaml:"nodeGroupNamePrefix,omitempty"`

	// Name of the EKS Cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// Disk size in GiB for worker nodes. Defaults to `50` for Windows, `20` all other node groups. The provider will only perform drift detection if a configuration value is provided.
	DiskSize int `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`

	// Name of the EKS Node Group. If omitted, the provider will assign a random, unique name. Conflicts with `node_group_name_prefix`. The node group name can't be longer than 63 characters. It must start with a letter or digit, but can also include hyphens and underscores for the remaining characters.
	NodeGroupName string `json:"nodeGroupName,omitempty" yaml:"nodeGroupName,omitempty"`

	// Configuration block with remote access settings. See `remote_access` below for details. Conflicts with `launch_template`.
	RemoteAccess types.Eks_NodeGroupRemoteAccess `json:"remoteAccess,omitempty" yaml:"remoteAccess,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Kubernetes taints to be applied to the nodes in the node group. Maximum of 50 taints per node group. See taint below for details.
	Taints []types.Eks_NodeGroupTaint `json:"taints,omitempty" yaml:"taints,omitempty"`

	// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateVersion bool `json:"forceUpdateVersion,omitempty" yaml:"forceUpdateVersion,omitempty"`

	// Configuration block with Launch Template settings. See `launch_template` below for details. Conflicts with `remote_access`.
	LaunchTemplate types.Eks_NodeGroupLaunchTemplate `json:"launchTemplate,omitempty" yaml:"launchTemplate,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
	NodeRoleArn string `json:"nodeRoleArn,omitempty" yaml:"nodeRoleArn,omitempty"`

	// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
	ReleaseVersion string `json:"releaseVersion,omitempty" yaml:"releaseVersion,omitempty"`

	// Configuration block with scaling settings. See `scaling_config` below for details.
	ScalingConfig types.Eks_NodeGroupScalingConfig `json:"scalingConfig,omitempty" yaml:"scalingConfig,omitempty"`

	// Configuration block with update settings. See `update_config` below for details.
	UpdateConfig types.Eks_NodeGroupUpdateConfig `json:"updateConfig,omitempty" yaml:"updateConfig,omitempty"`
}
