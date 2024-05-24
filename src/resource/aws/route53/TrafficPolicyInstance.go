package route53

type TrafficPolicyInstance struct {
	// ID of the hosted zone that you want Amazon Route 53 to create resource record sets in by using the configuration in a traffic policy.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// Domain name for which Amazon Route 53 responds to DNS queries by using the resource record sets that Route 53 creates for this traffic policy instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ID of the traffic policy that you want to use to create resource record sets in the specified hosted zone.
	TrafficPolicyId string `json:"trafficPolicyId,omitempty" yaml:"trafficPolicyId,omitempty"`

	// Version of the traffic policy
	TrafficPolicyVersion int `json:"trafficPolicyVersion,omitempty" yaml:"trafficPolicyVersion,omitempty"`

	// TTL that you want Amazon Route 53 to assign to all the resource record sets that it creates in the specified hosted zone.
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
