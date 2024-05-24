package ecr

import types "DesignSphere_Server/src/resource/aws/types"

type RegistryScanningConfiguration struct {
	// One or multiple blocks specifying scanning rules to determine which repository filters are used and at what frequency scanning will occur. See below for schema.
	Rules []types.Ecr_RegistryScanningConfigurationRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// the scanning type to set for the registry. Can be either `ENHANCED` or `BASIC`.
	ScanType string `json:"scanType,omitempty" yaml:"scanType,omitempty"`
}
