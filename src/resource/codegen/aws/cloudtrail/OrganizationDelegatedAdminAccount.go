package cloudtrail

type OrganizationDelegatedAdminAccount struct {
	// An organization member account ID that you want to designate as a delegated administrator.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
