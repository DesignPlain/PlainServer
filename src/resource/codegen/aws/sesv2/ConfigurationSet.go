package sesv2

import types "libds/aws/types"

type ConfigurationSet struct {
	// A map of tags to assign to the service. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An object that defines the open and click tracking options for emails that you send using the configuration set. See `tracking_options` Block for details.
	TrackingOptions types.Sesv2_ConfigurationSetTrackingOptions `json:"trackingOptions,omitempty" yaml:"trackingOptions,omitempty"`

	// An object that defines the VDM settings that apply to emails that you send using the configuration set. See `vdm_options` Block for details.
	VdmOptions types.Sesv2_ConfigurationSetVdmOptions `json:"vdmOptions,omitempty" yaml:"vdmOptions,omitempty"`

	// The name of the configuration set.
	ConfigurationSetName string `json:"configurationSetName,omitempty" yaml:"configurationSetName,omitempty"`

	// An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set. See `delivery_options` Block for details.
	DeliveryOptions types.Sesv2_ConfigurationSetDeliveryOptions `json:"deliveryOptions,omitempty" yaml:"deliveryOptions,omitempty"`

	// An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set. See `reputation_options` Block for details.
	ReputationOptions types.Sesv2_ConfigurationSetReputationOptions `json:"reputationOptions,omitempty" yaml:"reputationOptions,omitempty"`

	// An object that defines whether or not Amazon SES can send email that you send using the configuration set. See `sending_options` Block for details.
	SendingOptions types.Sesv2_ConfigurationSetSendingOptions `json:"sendingOptions,omitempty" yaml:"sendingOptions,omitempty"`

	// An object that contains information about the suppression list preferences for your account. See `suppression_options` Block for details.
	SuppressionOptions types.Sesv2_ConfigurationSetSuppressionOptions `json:"suppressionOptions,omitempty" yaml:"suppressionOptions,omitempty"`
}
