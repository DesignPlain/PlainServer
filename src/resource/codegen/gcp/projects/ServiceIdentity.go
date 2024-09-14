package projects

type ServiceIdentity struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The service to generate identity for.

	   - - -
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
