package types

type Batch_ComputeEnvironmentEksConfiguration struct {
	// The Amazon Resource Name (ARN) of the Amazon EKS cluster.
	EksClusterArn string `json:"eksClusterArn,omitempty" yaml:"eksClusterArn,omitempty"`

	// The namespace of the Amazon EKS cluster. AWS Batch manages pods in this namespace.
	KubernetesNamespace string `json:"kubernetesNamespace,omitempty" yaml:"kubernetesNamespace,omitempty"`
}
