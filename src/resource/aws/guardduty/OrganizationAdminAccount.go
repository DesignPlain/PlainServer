package guardduty

type OrganizationAdminAccount struct {
	// AWS account identifier to designate as a delegated administrator for GuardDuty.
	AdminAccountId string `json:"adminAccountId,omitempty" yaml:"adminAccountId,omitempty"`
}
