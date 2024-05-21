package types

type Monitoring_AlertPolicyAlertStrategy struct {
	// If an alert policy that was active has no data for this long, any open incidents will close.
	AutoClose string `json:"autoClose,omitempty" yaml:"autoClose,omitempty"`

	/*
	   Control over how the notification channels in `notification_channels`
	   are notified when this alert fires, on a per-channel basis.
	   Structure is documented below.
	*/
	NotificationChannelStrategies []Monitoring_AlertPolicyAlertStrategyNotificationChannelStrategy `json:"notificationChannelStrategies,omitempty" yaml:"notificationChannelStrategies,omitempty"`

	/*
	   Required for alert policies with a LogMatch condition.
	   This limit is not implemented for alert policies that are not log-based.
	   Structure is documented below.
	*/
	NotificationRateLimit Monitoring_AlertPolicyAlertStrategyNotificationRateLimit `json:"notificationRateLimit,omitempty" yaml:"notificationRateLimit,omitempty"`
}
