package types

type Osconfig_GuestPoliciesRecipeInstallStepFileExec struct {
	// A list of possible return values that the program can return to indicate a success. Defaults to [0].
	AllowedExitCodes string `json:"allowedExitCodes,omitempty" yaml:"allowedExitCodes,omitempty"`

	// Arguments to be passed to the provided executable.
	Args []string `json:"args,omitempty" yaml:"args,omitempty"`

	// The id of the relevant artifact in the recipe.
	ArtifactId string `json:"artifactId,omitempty" yaml:"artifactId,omitempty"`

	// The absolute path of the file on the local filesystem.
	LocalPath string `json:"localPath,omitempty" yaml:"localPath,omitempty"`
}
