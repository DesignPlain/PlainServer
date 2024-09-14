package route53recoverycontrol

import types "libds/aws/types"

type SafetyRule struct {
	// Gating controls for the new gating rule. That is, routing controls that are evaluated by the rule configuration that you specify.
	GatingControls []string `json:"gatingControls,omitempty" yaml:"gatingControls,omitempty"`

	// Name describing the safety rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for safety rule criteria. See below.
	RuleConfig types.Route53recoverycontrol_SafetyRuleRuleConfig `json:"ruleConfig,omitempty" yaml:"ruleConfig,omitempty"`

	// Routing controls that can only be set or unset if the specified `rule_config` evaluates to true for the specified `gating_controls`.
	TargetControls []string `json:"targetControls,omitempty" yaml:"targetControls,omitempty"`

	/*
	   Evaluation period, in milliseconds (ms), during which any request against the target routing controls will fail.

	   The following arguments are optional:
	*/
	WaitPeriodMs int `json:"waitPeriodMs,omitempty" yaml:"waitPeriodMs,omitempty"`

	// Routing controls that are part of transactions that are evaluated to determine if a request to change a routing control state is allowed.
	AssertedControls []string `json:"assertedControls,omitempty" yaml:"assertedControls,omitempty"`

	// ARN of the control panel in which this safety rule will reside.
	ControlPanelArn string `json:"controlPanelArn,omitempty" yaml:"controlPanelArn,omitempty"`
}
