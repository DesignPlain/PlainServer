package detective

type OrganizationAdminAccount struct {
	// AWS account identifier to designate as a delegated administrator for Detective.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
