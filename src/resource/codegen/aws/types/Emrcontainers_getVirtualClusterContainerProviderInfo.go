package types

type Emrcontainers_getVirtualClusterContainerProviderInfo struct {
	// Nested list containing EKS-specific information about the cluster where the EMR Containers cluster is running
	EksInfos []Emrcontainers_getVirtualClusterContainerProviderInfoEksInfo `json:"eksInfos,omitempty" yaml:"eksInfos,omitempty"`
}
