package cfg

import types "libds/aws/types"

type RemediationConfiguration struct {
	// Version of the target. For example, version of the SSM document
	TargetVersion string `json:"targetVersion,omitempty" yaml:"targetVersion,omitempty"`

	// Type of resource.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Name of the AWS Config rule.
	ConfigRuleName string `json:"configRuleName,omitempty" yaml:"configRuleName,omitempty"`

	// Configuration block for execution controls. See below.
	ExecutionControls types.Cfg_RemediationConfigurationExecutionControls `json:"executionControls,omitempty" yaml:"executionControls,omitempty"`

	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts int `json:"maximumAutomaticAttempts,omitempty" yaml:"maximumAutomaticAttempts,omitempty"`

	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameters []types.Cfg_RemediationConfigurationParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds int `json:"retryAttemptSeconds,omitempty" yaml:"retryAttemptSeconds,omitempty"`

	// Target ID is the name of the public document.
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`

	/*
	   Type of the target. Target executes remediation. For example, SSM document.

	   The following arguments are optional:
	*/
	TargetType string `json:"targetType,omitempty" yaml:"targetType,omitempty"`

	// Remediation is triggered automatically if `true`.
	Automatic bool `json:"automatic,omitempty" yaml:"automatic,omitempty"`
}
