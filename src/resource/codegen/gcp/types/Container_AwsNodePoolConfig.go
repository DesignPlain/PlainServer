package types

type Container_AwsNodePoolConfig struct {
	// The OS image type to use on node pool instances.
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	// Optional. The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// Optional. Configuration related to CloudWatch metrics collection on the Auto Scaling group of the node pool. When unspecified, metrics collection is disabled.
	AutoscalingMetricsCollection Container_AwsNodePoolConfigAutoscalingMetricsCollection `json:"autoscalingMetricsCollection,omitempty" yaml:"autoscalingMetricsCollection,omitempty"`

	// The ARN of the AWS KMS key used to encrypt node pool configuration.
	ConfigEncryption Container_AwsNodePoolConfigConfigEncryption `json:"configEncryption,omitempty" yaml:"configEncryption,omitempty"`

	// Optional. The SSH configuration.
	SshConfig Container_AwsNodePoolConfigSshConfig `json:"sshConfig,omitempty" yaml:"sshConfig,omitempty"`

	// Optional. The initial taints assigned to nodes of this node pool.
	Taints []Container_AwsNodePoolConfigTaint `json:"taints,omitempty" yaml:"taints,omitempty"`

	// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Optional. When specified, the node pool will provision Spot instances from the set of spot_config.instance_types. This field is mutually exclusive with `instance_type`
	SpotConfig Container_AwsNodePoolConfigSpotConfig `json:"spotConfig,omitempty" yaml:"spotConfig,omitempty"`

	// Optional. Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Proxy configuration for outbound HTTP(S) traffic.
	ProxyConfig Container_AwsNodePoolConfigProxyConfig `json:"proxyConfig,omitempty" yaml:"proxyConfig,omitempty"`

	// Optional. Template for the root volume provisioned for node pool nodes. Volumes will be provisioned in the availability zone assigned to the node pool subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
	RootVolume Container_AwsNodePoolConfigRootVolume `json:"rootVolume,omitempty" yaml:"rootVolume,omitempty"`

	// The name of the AWS IAM role assigned to nodes in the pool.
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" yaml:"iamInstanceProfile,omitempty"`

	// Details of placement information for an instance.
	InstancePlacement Container_AwsNodePoolConfigInstancePlacement `json:"instancePlacement,omitempty" yaml:"instancePlacement,omitempty"`

	// Optional. The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
