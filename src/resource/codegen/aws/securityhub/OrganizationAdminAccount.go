package securityhub

type OrganizationAdminAccount struct {
	// The AWS account identifier of the account to designate as the Security Hub administrator account.
	AdminAccountId string `json:"adminAccountId,omitempty" yaml:"adminAccountId,omitempty"`
}
