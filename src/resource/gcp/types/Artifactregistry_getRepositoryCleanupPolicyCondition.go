package types

type Artifactregistry_getRepositoryCleanupPolicyCondition struct {
	// Match versions older than a duration.
	OlderThan string `json:"olderThan,omitempty" yaml:"olderThan,omitempty"`

	// Match versions by package prefix. Applied on any prefix match.
	PackageNamePrefixes []string `json:"packageNamePrefixes,omitempty" yaml:"packageNamePrefixes,omitempty"`

	// Match versions by tag prefix. Applied on any prefix match.
	TagPrefixes []string `json:"tagPrefixes,omitempty" yaml:"tagPrefixes,omitempty"`

	// Match versions by tag status. Default value: "ANY" Possible values: ["TAGGED", "UNTAGGED", "ANY"]
	TagState string `json:"tagState,omitempty" yaml:"tagState,omitempty"`

	// Match versions by version name prefix. Applied on any prefix match.
	VersionNamePrefixes []string `json:"versionNamePrefixes,omitempty" yaml:"versionNamePrefixes,omitempty"`

	// Match versions newer than a duration.
	NewerThan string `json:"newerThan,omitempty" yaml:"newerThan,omitempty"`
}
