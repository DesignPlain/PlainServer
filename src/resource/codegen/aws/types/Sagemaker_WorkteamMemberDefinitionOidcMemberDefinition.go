package types

type Sagemaker_WorkteamMemberDefinitionOidcMemberDefinition struct {
	// A list of comma separated strings that identifies user groups in your OIDC IdP. Each user group is made up of a group of private workers.
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`
}
