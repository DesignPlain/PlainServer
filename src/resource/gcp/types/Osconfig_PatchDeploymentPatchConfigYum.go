package types

type Osconfig_PatchDeploymentPatchConfigYum struct {
	/*
	   An exclusive list of packages to be updated. These are the only packages that will be updated.
	   If these packages are not installed, they will be ignored. This field cannot be specified with
	   any other patch configuration fields.
	*/
	ExclusivePackages []string `json:"exclusivePackages,omitempty" yaml:"exclusivePackages,omitempty"`

	// Will cause patch to run yum update-minimal instead.
	Minimal bool `json:"minimal,omitempty" yaml:"minimal,omitempty"`

	// Adds the --security flag to yum update. Not supported on all platforms.
	Security bool `json:"security,omitempty" yaml:"security,omitempty"`

	// List of packages to exclude from update. These packages will be excluded.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`
}
