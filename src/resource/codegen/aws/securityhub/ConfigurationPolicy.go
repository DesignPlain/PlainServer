package securityhub

import types "libds/aws/types"

type ConfigurationPolicy struct {
	// The name of the configuration policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Defines how Security Hub is configured. See below.
	ConfigurationPolicy types.Securityhub_ConfigurationPolicyConfigurationPolicy `json:"configurationPolicy,omitempty" yaml:"configurationPolicy,omitempty"`

	// The description of the configuration policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
