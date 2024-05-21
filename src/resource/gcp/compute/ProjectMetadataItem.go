package compute

type ProjectMetadataItem struct {
	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The value to set for the given metadata key.

	   - - -
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The metadata key to set.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
