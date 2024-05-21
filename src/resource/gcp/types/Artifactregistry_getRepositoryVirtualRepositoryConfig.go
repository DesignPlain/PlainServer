package types

type Artifactregistry_getRepositoryVirtualRepositoryConfig struct {
	/*
	   Policies that configure the upstream artifacts distributed by the Virtual
	   Repository. Upstream policies cannot be set on a standard repository.
	*/
	UpstreamPolicies []Artifactregistry_getRepositoryVirtualRepositoryConfigUpstreamPolicy `json:"upstreamPolicies,omitempty" yaml:"upstreamPolicies,omitempty"`
}
