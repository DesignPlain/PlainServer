package types

type Artifactregistry_getRepositoryCleanupPolicyMostRecentVersion struct {
	// Minimum number of versions to keep.
	KeepCount int `json:"keepCount,omitempty" yaml:"keepCount,omitempty"`

	// Match versions by package prefix. Applied on any prefix match.
	PackageNamePrefixes []string `json:"packageNamePrefixes,omitempty" yaml:"packageNamePrefixes,omitempty"`
}
