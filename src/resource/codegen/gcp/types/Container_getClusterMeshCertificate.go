package types

type Container_getClusterMeshCertificate struct {
	// When enabled the GKE Workload Identity Certificates controller and node agent will be deployed in the cluster.
	EnableCertificates bool `json:"enableCertificates,omitempty" yaml:"enableCertificates,omitempty"`
}
