package types

type Artifactregistry_RepositoryVirtualRepositoryConfigUpstreamPolicy struct {
	// The user-provided ID of the upstream policy.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Entries with a greater priority value take precedence in the pull order.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   A reference to the repository resource, for example:
	   "projects/p1/locations/us-central1/repository/repo1".
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
