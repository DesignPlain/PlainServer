package kms

type KeyRing struct {
	/*
	   The location for the KeyRing.
	   A full list of valid locations can be found by running `gcloud kms locations list`.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The resource name for the KeyRing.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
