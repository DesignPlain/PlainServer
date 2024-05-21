package dns

import types "DesignSphere_Server/src/resource/gcp/types"

type ResponsePolicyRule struct {
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	/*
	   Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name;
	   in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	   Structure is documented below.
	*/
	LocalData types.Dns_ResponsePolicyRuleLocalData `json:"localData,omitempty" yaml:"localData,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Identifies the response policy addressed by this request.


	   - - -
	*/
	ResponsePolicy string `json:"responsePolicy,omitempty" yaml:"responsePolicy,omitempty"`

	// An identifier for this rule. Must be unique with the ResponsePolicy.
	RuleName string `json:"ruleName,omitempty" yaml:"ruleName,omitempty"`

	// Answer this query with a behavior rather than DNS data. Acceptable values are 'behaviorUnspecified', and 'bypassResponsePolicy'
	Behavior string `json:"behavior,omitempty" yaml:"behavior,omitempty"`
}
