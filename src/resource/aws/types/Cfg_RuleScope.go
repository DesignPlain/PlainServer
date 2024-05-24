package types

type Cfg_RuleScope struct {
	// The IDs of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for `compliance_resource_types`.
	ComplianceResourceId string `json:"complianceResourceId,omitempty" yaml:"complianceResourceId,omitempty"`

	// A list of resource types of only those AWS resources that you want to trigger an evaluation for the ruleE.g., `AWS::EC2::Instance`. You can only specify one type if you also specify a resource ID for `compliance_resource_id`. See [relevant part of AWS Docs](http://docs.aws.amazon.com/config/latest/APIReference/API_ResourceIdentifier.html#config-Type-ResourceIdentifier-resourceType) for available
	ComplianceResourceTypes []string `json:"complianceResourceTypes,omitempty" yaml:"complianceResourceTypes,omitempty"`

	// The tag key that is applied to only those AWS resources that you want you want to trigger an evaluation for the rule.
	TagKey string `json:"tagKey,omitempty" yaml:"tagKey,omitempty"`

	// The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule.
	TagValue string `json:"tagValue,omitempty" yaml:"tagValue,omitempty"`
}
