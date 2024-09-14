package compute

type ProjectMetadata struct {
	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A series of key value pairs.

	   - - -
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}
