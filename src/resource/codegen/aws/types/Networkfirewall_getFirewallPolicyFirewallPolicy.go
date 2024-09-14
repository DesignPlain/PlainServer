package types

type Networkfirewall_getFirewallPolicyFirewallPolicy struct {
	//
	StatefulRuleGroupReferences []Networkfirewall_getFirewallPolicyFirewallPolicyStatefulRuleGroupReference `json:"statefulRuleGroupReferences,omitempty" yaml:"statefulRuleGroupReferences,omitempty"`

	//
	StatelessCustomActions []Networkfirewall_getFirewallPolicyFirewallPolicyStatelessCustomAction `json:"statelessCustomActions,omitempty" yaml:"statelessCustomActions,omitempty"`

	//
	StatelessDefaultActions []string `json:"statelessDefaultActions,omitempty" yaml:"statelessDefaultActions,omitempty"`

	//
	StatelessFragmentDefaultActions []string `json:"statelessFragmentDefaultActions,omitempty" yaml:"statelessFragmentDefaultActions,omitempty"`

	//
	StatelessRuleGroupReferences []Networkfirewall_getFirewallPolicyFirewallPolicyStatelessRuleGroupReference `json:"statelessRuleGroupReferences,omitempty" yaml:"statelessRuleGroupReferences,omitempty"`

	//
	TlsInspectionConfigurationArn string `json:"tlsInspectionConfigurationArn,omitempty" yaml:"tlsInspectionConfigurationArn,omitempty"`

	//
	StatefulDefaultActions []string `json:"statefulDefaultActions,omitempty" yaml:"statefulDefaultActions,omitempty"`

	//
	StatefulEngineOptions []Networkfirewall_getFirewallPolicyFirewallPolicyStatefulEngineOption `json:"statefulEngineOptions,omitempty" yaml:"statefulEngineOptions,omitempty"`
}
