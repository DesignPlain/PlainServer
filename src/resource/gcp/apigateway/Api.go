package apigateway

type Api struct {
	/*
	   Identifier to assign to the API. Must be unique within scope of the parent resource(project)


	   - - -
	*/
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// A user-visible name for the API.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	   If not specified, a new Service will automatically be created in the same project as this API.
	*/
	ManagedService string `json:"managedService,omitempty" yaml:"managedService,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
