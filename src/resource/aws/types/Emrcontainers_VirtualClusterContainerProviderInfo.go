package types

type Emrcontainers_VirtualClusterContainerProviderInfo struct {
	// Nested list containing EKS-specific information about the cluster where the EMR Containers cluster is running
	EksInfo Emrcontainers_VirtualClusterContainerProviderInfoEksInfo `json:"eksInfo,omitempty" yaml:"eksInfo,omitempty"`
}
