package eks

type AccessEntry struct {
	// Defaults to STANDARD which provides the standard workflow. EC2_LINUX, EC2_WINDOWS, FARGATE_LINUX types disallow users to input a username or groups, and prevent associations.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Defaults to principal ARN if user is principal else defaults to assume-role/session-name is role is used.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`

	// Name of the EKS Cluster.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// List of string which can optionally specify the Kubernetes groups the user would belong to when creating an access entry.
	KubernetesGroups []string `json:"kubernetesGroups,omitempty" yaml:"kubernetesGroups,omitempty"`

	/*
	   The IAM Principal ARN which requires Authentication access to the EKS cluster.

	   The following arguments are optional:
	*/
	PrincipalArn string `json:"principalArn,omitempty" yaml:"principalArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
