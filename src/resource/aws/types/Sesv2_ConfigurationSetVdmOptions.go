package types

type Sesv2_ConfigurationSetVdmOptions struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardOptions Sesv2_ConfigurationSetVdmOptionsDashboardOptions `json:"dashboardOptions,omitempty" yaml:"dashboardOptions,omitempty"`

	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianOptions Sesv2_ConfigurationSetVdmOptionsGuardianOptions `json:"guardianOptions,omitempty" yaml:"guardianOptions,omitempty"`
}
