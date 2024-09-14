package ses

import types "libds/aws/types"

type ConfigurationSet struct {
	// Whether messages that use the configuration set are required to use TLS. See below.
	DeliveryOptions types.Ses_ConfigurationSetDeliveryOptions `json:"deliveryOptions,omitempty" yaml:"deliveryOptions,omitempty"`

	/*
	   Name of the configuration set.

	   The following argument is optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is `false`.
	ReputationMetricsEnabled bool `json:"reputationMetricsEnabled,omitempty" yaml:"reputationMetricsEnabled,omitempty"`

	// Whether email sending is enabled or disabled for the configuration set. The default value is `true`.
	SendingEnabled bool `json:"sendingEnabled,omitempty" yaml:"sendingEnabled,omitempty"`

	// Domain that is used to redirect email recipients to an Amazon SES-operated domain. See below. --NOTE:-- This functionality is best effort.
	TrackingOptions types.Ses_ConfigurationSetTrackingOptions `json:"trackingOptions,omitempty" yaml:"trackingOptions,omitempty"`
}
