package sesv2

import types "libds/aws/types"

type AccountVdmAttributes struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes types.Sesv2_AccountVdmAttributesDashboardAttributes `json:"dashboardAttributes,omitempty" yaml:"dashboardAttributes,omitempty"`

	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes types.Sesv2_AccountVdmAttributesGuardianAttributes `json:"guardianAttributes,omitempty" yaml:"guardianAttributes,omitempty"`

	/*
	   Specifies the status of your VDM configuration. Valid values: `ENABLED`, `DISABLED`.

	   The following arguments are optional:
	*/
	VdmEnabled string `json:"vdmEnabled,omitempty" yaml:"vdmEnabled,omitempty"`
}
