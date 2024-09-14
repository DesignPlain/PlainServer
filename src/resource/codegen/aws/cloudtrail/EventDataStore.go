package cloudtrail

import types "libds/aws/types"

type EventDataStore struct {
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
	OrganizationEnabled bool `json:"organizationEnabled,omitempty" yaml:"organizationEnabled,omitempty"`

	// The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The billing mode for the event data store. The valid values are `EXTENDABLE_RETENTION_PRICING` and `FIXED_RETENTION_PRICING`. Defaults to `EXTENDABLE_RETENTION_PRICING`.
	BillingMode string `json:"billingMode,omitempty" yaml:"billingMode,omitempty"`

	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The name of the event data store.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
	TerminationProtectionEnabled bool `json:"terminationProtectionEnabled,omitempty" yaml:"terminationProtectionEnabled,omitempty"`

	// The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	AdvancedEventSelectors []types.Cloudtrail_EventDataStoreAdvancedEventSelector `json:"advancedEventSelectors,omitempty" yaml:"advancedEventSelectors,omitempty"`

	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
	MultiRegionEnabled bool `json:"multiRegionEnabled,omitempty" yaml:"multiRegionEnabled,omitempty"`
}
