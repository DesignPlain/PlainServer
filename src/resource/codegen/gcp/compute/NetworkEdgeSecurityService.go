package compute

type NetworkEdgeSecurityService struct {
	// Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is created.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the gateway security policy.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The resource URL for the network edge security service associated with this network edge security service.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`
}
