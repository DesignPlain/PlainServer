package gamelift

import types "libds/aws/types"

type GameSessionQueue struct {
	// Name of the session queue.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An SNS topic ARN that is set up to receive game session placement notifications.
	NotificationTarget string `json:"notificationTarget,omitempty" yaml:"notificationTarget,omitempty"`

	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies []types.Gamelift_GameSessionQueuePlayerLatencyPolicy `json:"playerLatencyPolicies,omitempty" yaml:"playerLatencyPolicies,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds int `json:"timeoutInSeconds,omitempty" yaml:"timeoutInSeconds,omitempty"`

	// Information to be added to all events that are related to this game session queue.
	CustomEventData string `json:"customEventData,omitempty" yaml:"customEventData,omitempty"`

	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []string `json:"destinations,omitempty" yaml:"destinations,omitempty"`
}
