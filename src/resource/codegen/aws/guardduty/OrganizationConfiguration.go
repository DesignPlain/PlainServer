package guardduty

import types "libds/aws/types"

type OrganizationConfiguration struct {
	// The detector ID of the GuardDuty account.
	DetectorId string `json:"detectorId,omitempty" yaml:"detectorId,omitempty"`

	// -Deprecated:- Use `auto_enable_organization_members` instead. When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable bool `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// Indicates the auto-enablement configuration of GuardDuty for the member accounts in the organization. Valid values are `ALL`, `NEW`, `NONE`.
	AutoEnableOrganizationMembers string `json:"autoEnableOrganizationMembers,omitempty" yaml:"autoEnableOrganizationMembers,omitempty"`

	// Configuration for the collected datasources.
	Datasources types.Guardduty_OrganizationConfigurationDatasources `json:"datasources,omitempty" yaml:"datasources,omitempty"`
}
