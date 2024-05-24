package types

type Sesv2_getConfigurationSetVdmOption struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardOptions []Sesv2_getConfigurationSetVdmOptionDashboardOption `json:"dashboardOptions,omitempty" yaml:"dashboardOptions,omitempty"`

	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianOptions []Sesv2_getConfigurationSetVdmOptionGuardianOption `json:"guardianOptions,omitempty" yaml:"guardianOptions,omitempty"`
}
