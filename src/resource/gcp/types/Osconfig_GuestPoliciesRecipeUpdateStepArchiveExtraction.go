package types

type Osconfig_GuestPoliciesRecipeUpdateStepArchiveExtraction struct {
	// The id of the relevant artifact in the recipe.
	ArtifactId string `json:"artifactId,omitempty" yaml:"artifactId,omitempty"`

	// Directory to extract archive to. Defaults to / on Linux or C:\ on Windows.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	/*
	   The type of the archive to extract.
	   Possible values are: `TAR`, `TAR_GZIP`, `TAR_BZIP`, `TAR_LZMA`, `TAR_XZ`, `ZIP`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
