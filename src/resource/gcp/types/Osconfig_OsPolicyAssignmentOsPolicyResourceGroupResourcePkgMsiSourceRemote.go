package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgMsiSourceRemote struct {
	// SHA256 checksum of the remote file.
	Sha256Checksum string `json:"sha256Checksum,omitempty" yaml:"sha256Checksum,omitempty"`

	/*
	   URI from which to fetch the object. It should contain
	   both the protocol and path following the format `{protocol}://{location}`.
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
