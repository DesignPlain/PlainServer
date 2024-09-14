package securitylake

import types "libds/aws/types"

type Subscriber struct {
	//
	Timeouts types.Securitylake_SubscriberTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The Amazon S3 or Lake Formation access type.
	AccessType string `json:"accessType,omitempty" yaml:"accessType,omitempty"`

	// The supported AWS services from which logs and events are collected. Security Lake supports log and event collection for natively supported AWS services. See `source` Blocks below.
	Source types.Securitylake_SubscriberSource `json:"source,omitempty" yaml:"source,omitempty"`

	// The description for your subscriber account in Security Lake.
	SubscriberDescription string `json:"subscriberDescription,omitempty" yaml:"subscriberDescription,omitempty"`

	// The AWS identity used to access your data. See `subscriber_identity` Block below.
	SubscriberIdentity types.Securitylake_SubscriberSubscriberIdentity `json:"subscriberIdentity,omitempty" yaml:"subscriberIdentity,omitempty"`

	// The name of your Security Lake subscriber account.
	SubscriberName string `json:"subscriberName,omitempty" yaml:"subscriberName,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
