package networkservices

type ServiceBinding struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Set of label tags associated with the ServiceBinding resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Name of the ServiceBinding resource.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The full Service Directory Service name of the format
	   projects/-/locations/-/namespaces/-/services/-
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
