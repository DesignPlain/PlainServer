package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceRepositoryZypper struct {
	// The location of the repository directory.
	BaseUrl string `json:"baseUrl,omitempty" yaml:"baseUrl,omitempty"`

	// The display name of the repository.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// URIs of GPG keys.
	GpgKeys []string `json:"gpgKeys,omitempty" yaml:"gpgKeys,omitempty"`

	/*
	   A one word, unique name for this repository. This is the
	   `repo id` in the zypper config file and also the `display_name` if
	   `display_name` is omitted. This id is also used as the unique identifier
	   when checking for GuestPolicy conflicts.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
}
