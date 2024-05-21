package networksecurity

type GatewaySecurityPolicy struct {
	/*
	   Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
	   gatewaySecurityPolicy should match the pattern:(^a-z?$).


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
	TlsInspectionPolicy string `json:"tlsInspectionPolicy,omitempty" yaml:"tlsInspectionPolicy,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The location of the gateway security policy.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
