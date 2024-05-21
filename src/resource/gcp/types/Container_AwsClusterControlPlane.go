package types

type Container_AwsClusterControlPlane struct {
	// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Proxy configuration for outbound HTTP(S) traffic.
	ProxyConfig Container_AwsClusterControlPlaneProxyConfig `json:"proxyConfig,omitempty" yaml:"proxyConfig,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each control plane replica. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
	RootVolume Container_AwsClusterControlPlaneRootVolume `json:"rootVolume,omitempty" yaml:"rootVolume,omitempty"`

	// Optional. The IDs of additional security groups to add to control plane replicas. The Anthos Multi-Cloud API will automatically create and manage security groups with the minimum rules needed for a functioning cluster.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling .
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// The list of subnets where control plane replicas will run. A replica will be provisioned on each subnet and up to three values can be provided. Each subnet must be in a different AWS Availability Zone (AZ).
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Authentication configuration for management of AWS resources.
	AwsServicesAuthentication Container_AwsClusterControlPlaneAwsServicesAuthentication `json:"awsServicesAuthentication,omitempty" yaml:"awsServicesAuthentication,omitempty"`

	// The name of the AWS IAM instance pofile to assign to each control plane replica.
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" yaml:"iamInstanceProfile,omitempty"`

	// Details of placement information for an instance.
	InstancePlacement Container_AwsClusterControlPlaneInstancePlacement `json:"instancePlacement,omitempty" yaml:"instancePlacement,omitempty"`

	// Optional. SSH configuration for how to access the underlying control plane machines.
	SshConfig Container_AwsClusterControlPlaneSshConfig `json:"sshConfig,omitempty" yaml:"sshConfig,omitempty"`

	// The ARN of the AWS KMS key used to encrypt cluster configuration.
	ConfigEncryption Container_AwsClusterControlPlaneConfigEncryption `json:"configEncryption,omitempty" yaml:"configEncryption,omitempty"`

	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	DatabaseEncryption Container_AwsClusterControlPlaneDatabaseEncryption `json:"databaseEncryption,omitempty" yaml:"databaseEncryption,omitempty"`

	// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 8 GiB with the GP2 volume type.
	MainVolume Container_AwsClusterControlPlaneMainVolume `json:"mainVolume,omitempty" yaml:"mainVolume,omitempty"`

	// Optional. A set of AWS resource tags to propagate to all underlying managed AWS resources. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
