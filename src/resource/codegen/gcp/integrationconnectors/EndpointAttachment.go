package integrationconnectors

type EndpointAttachment struct {
	/*
	   Resource labels to represent user provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Location in which Endpoint Attachment needs to be created.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Name of Endpoint Attachment needs to be created.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The path of the service attachment.
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`

	// Description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
