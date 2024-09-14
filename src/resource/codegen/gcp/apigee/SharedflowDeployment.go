package apigee

type SharedflowDeployment struct {
	// The Apigee Organization associated with the Sharedflow
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   Revision of the Sharedflow to be deployed.


	   - - -
	*/
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`

	// The service account represents the identity of the deployed proxy, and determines what permissions it has. The format must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Id of the Sharedflow to be deployed.
	SharedflowId string `json:"sharedflowId,omitempty" yaml:"sharedflowId,omitempty"`

	// The resource ID of the environment.
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`
}
