package apigee

type EnvKeystore struct {
	/*
	   The Apigee environment group associated with the Apigee environment,
	   in the format `organizations/{{org_name}}/environments/{{env_name}}`.


	   - - -
	*/
	EnvId string `json:"envId,omitempty" yaml:"envId,omitempty"`

	// The name of the newly created keystore.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
