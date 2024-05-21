package apigee

type EnvGroupAttachment struct {
	/*
	   The Apigee environment group associated with the Apigee environment,
	   in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.


	   - - -
	*/
	EnvgroupId string `json:"envgroupId,omitempty" yaml:"envgroupId,omitempty"`

	// The resource ID of the environment.
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`
}
