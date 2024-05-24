package types

type Organizations_getDelegatedServicesDelegatedService struct {
	// The date that the account became a delegated administrator for this service.
	DelegationEnabledDate string `json:"delegationEnabledDate,omitempty" yaml:"delegationEnabledDate,omitempty"`

	// The name of an AWS service that can request an operation for the specified service.
	ServicePrincipal string `json:"servicePrincipal,omitempty" yaml:"servicePrincipal,omitempty"`
}
