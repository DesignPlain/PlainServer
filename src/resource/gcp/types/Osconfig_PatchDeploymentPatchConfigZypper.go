package types

type Osconfig_PatchDeploymentPatchConfigZypper struct {
	// Adds the --with-update flag, to zypper patch.
	WithUpdate bool `json:"withUpdate,omitempty" yaml:"withUpdate,omitempty"`

	// Install only patches with these categories. Common categories include security, recommended, and feature.
	Categories []string `json:"categories,omitempty" yaml:"categories,omitempty"`

	// List of packages to exclude from update.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty"`

	/*
	   An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command.
	   This field must not be used with any other patch configuration fields.
	*/
	ExclusivePatches []string `json:"exclusivePatches,omitempty" yaml:"exclusivePatches,omitempty"`

	// Install only patches with these severities. Common severities include critical, important, moderate, and low.
	Severities []string `json:"severities,omitempty" yaml:"severities,omitempty"`

	// Adds the --with-optional flag to zypper patch.
	WithOptional bool `json:"withOptional,omitempty" yaml:"withOptional,omitempty"`
}
