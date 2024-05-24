package macie2

type OrganizationAdminAccount struct {
	// The AWS account ID for the account to designate as the delegated Amazon Macie administrator account for the organization.
	AdminAccountId string `json:"adminAccountId,omitempty" yaml:"adminAccountId,omitempty"`
}
