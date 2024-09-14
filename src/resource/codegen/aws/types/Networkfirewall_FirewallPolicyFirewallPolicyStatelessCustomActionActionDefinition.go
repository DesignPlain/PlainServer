package types

type Networkfirewall_FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinition struct {
	// A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See Publish Metric Action below for details.
	PublishMetricAction Networkfirewall_FirewallPolicyFirewallPolicyStatelessCustomActionActionDefinitionPublishMetricAction `json:"publishMetricAction,omitempty" yaml:"publishMetricAction,omitempty"`
}
