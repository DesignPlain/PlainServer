package apigateway

type Gateway struct {
	/*
	   Resource name of the API Config for this Gateway. Format: projects/{project}/locations/global/apis/{api}/configs/{apiConfig}.
	   When changing api configs please ensure the new config is a new resource and the
	   lifecycle rule `create_before_destroy` is set.
	*/
	ApiConfig string `json:"apiConfig,omitempty" yaml:"apiConfig,omitempty"`

	// A user-visible name for the API.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Identifier to assign to the Gateway. Must be unique within scope of the parent resource(project).


	   - - -
	*/
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	/*
	   Resource labels to represent user-provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the gateway for the API.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
