package apigee

type EnvReferences struct {
	/*
	   The Apigee environment group associated with the Apigee environment,
	   in the format `organizations/{{org_name}}/environments/{{env_name}}`.


	   - - -
	*/
	EnvId string `json:"envId,omitempty" yaml:"envId,omitempty"`

	// Required. The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.
	Refers string `json:"refers,omitempty" yaml:"refers,omitempty"`

	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Optional. A human-readable description of this reference.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
