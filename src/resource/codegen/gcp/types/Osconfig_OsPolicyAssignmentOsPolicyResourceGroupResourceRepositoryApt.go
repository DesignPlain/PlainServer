package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryApt struct {
	// URI for this repository.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	/*
	   Type of archive files in this repository.
	   Possible values are: `ARCHIVE_TYPE_UNSPECIFIED`, `DEB`, `DEB_SRC`.
	*/
	ArchiveType string `json:"archiveType,omitempty" yaml:"archiveType,omitempty"`

	/*
	   List of components for this repository. Must
	   contain at least one item.
	*/
	Components []string `json:"components,omitempty" yaml:"components,omitempty"`

	// Distribution of this repository.
	Distribution string `json:"distribution,omitempty" yaml:"distribution,omitempty"`

	/*
	   URI of the key file for this repository. The agent
	   maintains a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg`.
	*/
	GpgKey string `json:"gpgKey,omitempty" yaml:"gpgKey,omitempty"`
}
