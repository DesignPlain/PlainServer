package apigee

type InstanceAttachment struct {
	/*
	   The Apigee instance associated with the Apigee environment,
	   in the format `organizations/{{org_name}}/instances/{{instance_name}}`.


	   - - -
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// The resource ID of the environment.
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`
}
