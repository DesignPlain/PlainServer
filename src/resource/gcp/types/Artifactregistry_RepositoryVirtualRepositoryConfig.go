package types

type Artifactregistry_RepositoryVirtualRepositoryConfig struct {
	/*
	   Policies that configure the upstream artifacts distributed by the Virtual
	   Repository. Upstream policies cannot be set on a standard repository.
	   Structure is documented below.
	*/
	UpstreamPolicies []Artifactregistry_RepositoryVirtualRepositoryConfigUpstreamPolicy `json:"upstreamPolicies,omitempty" yaml:"upstreamPolicies,omitempty"`
}
