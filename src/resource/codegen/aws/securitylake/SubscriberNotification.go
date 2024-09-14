package securitylake

import types "libds/aws/types"

type SubscriberNotification struct {
	// Specify the configuration using which you want to create the subscriber notification..
	Configuration types.Securitylake_SubscriberNotificationConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// The subscriber ID for the notification subscription.
	SubscriberId string `json:"subscriberId,omitempty" yaml:"subscriberId,omitempty"`
}
