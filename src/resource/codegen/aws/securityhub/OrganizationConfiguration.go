package securityhub

import types "libds/aws/types"

type OrganizationConfiguration struct {
	// Provides information about the way an organization is configured in Security Hub.
	OrganizationConfiguration types.Securityhub_OrganizationConfigurationOrganizationConfiguration `json:"organizationConfiguration,omitempty" yaml:"organizationConfiguration,omitempty"`

	// Whether to automatically enable Security Hub for new accounts in the organization.
	AutoEnable bool `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// Whether to automatically enable Security Hub default standards for new member accounts in the organization. By default, this parameter is equal to `DEFAULT`, and new member accounts are automatically enabled with default Security Hub standards. To opt out of enabling default standards for new member accounts, set this parameter equal to `NONE`.
	AutoEnableStandards string `json:"autoEnableStandards,omitempty" yaml:"autoEnableStandards,omitempty"`
}
