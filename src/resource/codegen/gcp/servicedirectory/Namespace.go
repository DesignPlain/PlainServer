package servicedirectory

type Namespace struct {
	/*
	   Resource labels associated with this Namespace. No more than 64 user
	   labels can be associated with a given resource. Label keys and values can
	   be no longer than 63 characters.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location for the Namespace.
	   A full list of valid locations can be found by running
	   `gcloud beta service-directory locations list`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The Resource ID must be 1-63 characters long, including digits,
	   lowercase letters or the hyphen character.


	   - - -
	*/
	NamespaceId string `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
