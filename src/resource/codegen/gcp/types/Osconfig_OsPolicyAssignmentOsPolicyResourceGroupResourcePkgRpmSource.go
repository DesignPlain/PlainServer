package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpmSource struct {
	/*
	   Defaults to false. When false, files are
	   subject to validations based on the file type: Remote: A checksum must be
	   specified. Cloud Storage: An object generation number must be specified.
	*/
	AllowInsecure bool `json:"allowInsecure,omitempty" yaml:"allowInsecure,omitempty"`

	/*
	   A Cloud Storage object. Structure is
	   documented below.
	*/
	Gcs Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpmSourceGcs `json:"gcs,omitempty" yaml:"gcs,omitempty"`

	// A local path within the VM to use.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`

	/*
	   A generic remote file. Structure is
	   documented below.
	*/
	Remote Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourcePkgRpmSourceRemote `json:"remote,omitempty" yaml:"remote,omitempty"`
}
