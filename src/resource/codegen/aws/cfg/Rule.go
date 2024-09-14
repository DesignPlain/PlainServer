package cfg

import types "libds/aws/types"

type Rule struct {
	// The name of the rule
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Scope defines which resources can trigger an evaluation for the rule. See Scope Below.
	Scope types.Cfg_RuleScope `json:"scope,omitempty" yaml:"scope,omitempty"`

	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Source Below.
	Source types.Cfg_RuleSource `json:"source,omitempty" yaml:"source,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the rule
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The modes the Config rule can be evaluated in. See Evaluation Mode for more details.
	EvaluationModes []types.Cfg_RuleEvaluationMode `json:"evaluationModes,omitempty" yaml:"evaluationModes,omitempty"`

	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters string `json:"inputParameters,omitempty" yaml:"inputParameters,omitempty"`

	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency string `json:"maximumExecutionFrequency,omitempty" yaml:"maximumExecutionFrequency,omitempty"`
}
