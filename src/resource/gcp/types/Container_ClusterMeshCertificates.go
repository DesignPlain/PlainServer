package types

type Container_ClusterMeshCertificates struct {
	// Controls the issuance of workload mTLS certificates. It is enabled by default. Workload Identity is required, see workload_config.
	EnableCertificates bool `json:"enableCertificates,omitempty" yaml:"enableCertificates,omitempty"`
}
