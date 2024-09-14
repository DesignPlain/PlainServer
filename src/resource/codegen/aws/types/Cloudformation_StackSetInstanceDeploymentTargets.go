package types

type Cloudformation_StackSetInstanceDeploymentTargets struct {
	// Limit deployment targets to individual accounts or include additional accounts with provided OUs. Valid values: `INTERSECTION`, `DIFFERENCE`, `UNION`, `NONE`.
	AccountFilterType string `json:"accountFilterType,omitempty" yaml:"accountFilterType,omitempty"`

	// List of accounts to deploy stack set updates.
	Accounts []string `json:"accounts,omitempty" yaml:"accounts,omitempty"`

	// S3 URL of the file containing the list of accounts.
	AccountsUrl string `json:"accountsUrl,omitempty" yaml:"accountsUrl,omitempty"`

	// Organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	OrganizationalUnitIds []string `json:"organizationalUnitIds,omitempty" yaml:"organizationalUnitIds,omitempty"`
}
