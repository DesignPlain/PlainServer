package types

type Osconfig_GuestPoliciesRecipeUpdateStepMsiInstallation struct {
	// Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]
	AllowedExitCodes []int `json:"allowedExitCodes,omitempty" yaml:"allowedExitCodes,omitempty"`

	// The id of the relevant artifact in the recipe.
	ArtifactId string `json:"artifactId,omitempty" yaml:"artifactId,omitempty"`

	// The flags to use when installing the MSI. Defaults to the install flag.
	Flags []string `json:"flags,omitempty" yaml:"flags,omitempty"`
}
