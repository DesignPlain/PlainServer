package types

type Cfg_RuleSourceCustomPolicyDetails struct {
	// The boolean expression for enabling debug logging for your Config Custom Policy rule. The default value is `false`.
	EnableDebugLogDelivery bool `json:"enableDebugLogDelivery,omitempty" yaml:"enableDebugLogDelivery,omitempty"`

	// The runtime system for your Config Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by Config Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://github.com/aws-cloudformation/cloudformation-guard).
	PolicyRuntime string `json:"policyRuntime,omitempty" yaml:"policyRuntime,omitempty"`

	// The policy definition containing the logic for your Config Custom Policy rule.
	PolicyText string `json:"policyText,omitempty" yaml:"policyText,omitempty"`
}
