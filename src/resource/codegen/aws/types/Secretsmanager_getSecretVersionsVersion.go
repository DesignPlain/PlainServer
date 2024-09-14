package types

type Secretsmanager_getSecretVersionsVersion struct {
	//
	CreatedTime string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`

	// Date that this version of the secret was last accessed.
	LastAccessedDate string `json:"lastAccessedDate,omitempty" yaml:"lastAccessedDate,omitempty"`

	// Unique version identifier of this version of the secret.
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`

	//
	VersionStages []string `json:"versionStages,omitempty" yaml:"versionStages,omitempty"`
}
