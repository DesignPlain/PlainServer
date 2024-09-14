package cloudbuildv2

type Repository struct {
	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the repository.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The connection for the resource


	   - - -
	*/
	ParentConnection string `json:"parentConnection,omitempty" yaml:"parentConnection,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Required. Git Clone HTTPS URI.
	RemoteUri string `json:"remoteUri,omitempty" yaml:"remoteUri,omitempty"`

	/*
	   Allows clients to store small amounts of arbitrary data.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}
