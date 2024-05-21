package types

type Artifactregistry_getRepositoryCleanupPolicy struct {
	// Policy action. Possible values: ["DELETE", "KEEP"]
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// Policy condition for matching versions.
	Conditions []Artifactregistry_getRepositoryCleanupPolicyCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	//
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   Policy condition for retaining a minimum number of versions. May only be
	   specified with a Keep action.
	*/
	MostRecentVersions []Artifactregistry_getRepositoryCleanupPolicyMostRecentVersion `json:"mostRecentVersions,omitempty" yaml:"mostRecentVersions,omitempty"`
}
