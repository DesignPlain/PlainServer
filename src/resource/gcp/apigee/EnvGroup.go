package apigee

type EnvGroup struct {
	// The resource ID of the environment group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The Apigee Organization associated with the Apigee environment group,
	   in the format `organizations/{{org_name}}`.


	   - - -
	*/
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Hostnames of the environment group.
	Hostnames []string `json:"hostnames,omitempty" yaml:"hostnames,omitempty"`
}
