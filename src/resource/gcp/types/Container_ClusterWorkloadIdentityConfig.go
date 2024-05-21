package types

type Container_ClusterWorkloadIdentityConfig struct {
	// The workload pool to attach all Kubernetes service accounts to.
	WorkloadPool string `json:"workloadPool,omitempty" yaml:"workloadPool,omitempty"`
}
