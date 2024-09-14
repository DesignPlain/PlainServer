package types

type Container_AttachedClusterWorkloadIdentityConfig struct {
	/*
	   The ID of the OIDC Identity Provider (IdP) associated to
	   the Workload Identity Pool.
	*/
	IdentityProvider string `json:"identityProvider,omitempty" yaml:"identityProvider,omitempty"`

	// The OIDC issuer URL for this cluster.
	IssuerUri string `json:"issuerUri,omitempty" yaml:"issuerUri,omitempty"`

	// The Workload Identity Pool associated to the cluster.
	WorkloadPool string `json:"workloadPool,omitempty" yaml:"workloadPool,omitempty"`
}
