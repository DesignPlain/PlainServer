package eks

type PodIdentityAssociation struct {
	// The name of the cluster to create the association in.
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`

	// The name of the Kubernetes namespace inside the cluster to create the association in. The service account and the pods that use the service account must be in this namespace.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to associate with the service account. The EKS Pod Identity agent manages credentials to assume this role for applications in the containers in the pods that use this service account.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	/*
	   The name of the Kubernetes service account inside the cluster to associate the IAM credentials with.

	   The following arguments are optional:
	*/
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
