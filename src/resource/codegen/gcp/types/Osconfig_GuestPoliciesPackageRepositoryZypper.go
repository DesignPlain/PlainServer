package types

type Osconfig_GuestPoliciesPackageRepositoryZypper struct {
	// URIs of GPG keys.
	GpgKeys []string `json:"gpgKeys,omitempty" yaml:"gpgKeys,omitempty"`

	/*
	   A one word, unique name for this repository. This is the repo id in the zypper config file and also the displayName
	   if displayName is omitted. This id is also used as the unique identifier when checking for guest policy conflicts.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The location of the repository directory.
	BaseUrl string `json:"baseUrl,omitempty" yaml:"baseUrl,omitempty"`

	// The display name of the repository.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
