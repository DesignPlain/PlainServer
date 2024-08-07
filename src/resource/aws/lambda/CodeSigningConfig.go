package lambda

import types "DesignSphere_Server/src/resource/aws/types"

type CodeSigningConfig struct {
	// Descriptive name for this code signing configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A configuration block of code signing policies that define the actions to take if the validation checks fail. Detailed below.
	Policies types.Lambda_CodeSigningConfigPolicies `json:"policies,omitempty" yaml:"policies,omitempty"`

	// A configuration block of allowed publishers as signing profiles for this code signing configuration. Detailed below.
	AllowedPublishers types.Lambda_CodeSigningConfigAllowedPublishers `json:"allowedPublishers,omitempty" yaml:"allowedPublishers,omitempty"`
}
