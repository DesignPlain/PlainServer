package types

type Guardduty_OrganizationConfigurationDatasourcesS3Logs struct {
	// -Deprecated:- Use `auto_enable_organization_members` instead. When this setting is enabled, all new accounts that are created in, or added to, the organization are added as a member accounts of the organization’s GuardDuty delegated administrator and GuardDuty is enabled in that AWS Region.
	AutoEnable bool `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`
}
