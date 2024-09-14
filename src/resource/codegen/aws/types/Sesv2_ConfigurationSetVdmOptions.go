package types

type Sesv2_ConfigurationSetVdmOptions struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard. See `dashboard_options` Block for details.
	DashboardOptions Sesv2_ConfigurationSetVdmOptionsDashboardOptions `json:"dashboardOptions,omitempty" yaml:"dashboardOptions,omitempty"`

	// Specifies additional settings for your VDM configuration as applicable to the Guardian. See `guardian_options` Block for details.
	GuardianOptions Sesv2_ConfigurationSetVdmOptionsGuardianOptions `json:"guardianOptions,omitempty" yaml:"guardianOptions,omitempty"`
}
