package types

type Osconfig_PatchDeploymentPatchConfigApt struct {
	// List of packages to exclude from update. These packages will be excluded.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	/*
	   An exclusive list of packages to be updated. These are the only packages that will be updated.
	   If these packages are not installed, they will be ignored. This field cannot be specified with
	   any other patch configuration fields.
	*/
	ExclusivePackages []string `json:"exclusivePackages,omitempty" yaml:"exclusivePackages,omitempty"`

	/*
	   By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead.
	   Possible values are: `DIST`, `UPGRADE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
