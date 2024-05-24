package organizations

type DelegatedAdministrator struct {
	// The service principal of the AWS service for which you want to make the member account a delegated administrator.
	ServicePrincipal string `json:"servicePrincipal,omitempty" yaml:"servicePrincipal,omitempty"`

	// The account ID number of the member account in the organization to register as a delegated administrator.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
