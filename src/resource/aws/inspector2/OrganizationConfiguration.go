package inspector2

import types "DesignSphere_Server/src/resource/aws/types"

type OrganizationConfiguration struct {
	// Configuration block for auto enabling. See below.
	AutoEnable types.Inspector2_OrganizationConfigurationAutoEnable `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`
}
