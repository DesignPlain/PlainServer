package types

type Artifactregistry_RepositoryCleanupPolicy struct {
	// The identifier for this object. Format specified above.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   Policy condition for retaining a minimum number of versions. May only be
	   specified with a Keep action.
	   Structure is documented below.
	*/
	MostRecentVersions Artifactregistry_RepositoryCleanupPolicyMostRecentVersions `json:"mostRecentVersions,omitempty" yaml:"mostRecentVersions,omitempty"`

	/*
	   Policy action.
	   Possible values are: `DELETE`, `KEEP`.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	/*
	   Policy condition for matching versions.
	   Structure is documented below.
	*/
	Condition Artifactregistry_RepositoryCleanupPolicyCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
