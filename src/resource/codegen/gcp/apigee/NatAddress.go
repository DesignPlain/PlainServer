package apigee

type NatAddress struct {
	/*
	   The Apigee instance associated with the Apigee environment,
	   in the format `organizations/{{org_name}}/instances/{{instance_name}}`.


	   - - -
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Resource ID of the NAT address.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
