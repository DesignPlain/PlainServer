package gamelift

import types "libds/aws/types"

type MatchmakingConfiguration struct {
	// Information to be added to all events related to this matchmaking configuration.
	CustomEventData string `json:"customEventData,omitempty" yaml:"customEventData,omitempty"`

	// Indicates whether this matchmaking configuration is being used with GameLift hosting or as a standalone matchmaking solution.
	FlexMatchMode string `json:"flexMatchMode,omitempty" yaml:"flexMatchMode,omitempty"`

	// Name of the matchmaking configuration
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An SNS topic ARN that is set up to receive matchmaking notifications.
	NotificationTarget string `json:"notificationTarget,omitempty" yaml:"notificationTarget,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies if the match that was created with this configuration must be accepted by matched players.
	AcceptanceRequired bool `json:"acceptanceRequired,omitempty" yaml:"acceptanceRequired,omitempty"`

	// One or more custom game properties. See below.
	GameProperties []types.Gamelift_MatchmakingConfigurationGameProperty `json:"gameProperties,omitempty" yaml:"gameProperties,omitempty"`

	// The ARNs of the GameLift game session queue resources.
	GameSessionQueueArns []string `json:"gameSessionQueueArns,omitempty" yaml:"gameSessionQueueArns,omitempty"`

	// The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.
	AcceptanceTimeoutSeconds int `json:"acceptanceTimeoutSeconds,omitempty" yaml:"acceptanceTimeoutSeconds,omitempty"`

	// A set of custom game session properties.
	GameSessionData string `json:"gameSessionData,omitempty" yaml:"gameSessionData,omitempty"`

	// A human-readable description of the matchmaking configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.
	RequestTimeoutSeconds int `json:"requestTimeoutSeconds,omitempty" yaml:"requestTimeoutSeconds,omitempty"`

	// A rule set names for the matchmaking rule set to use with this configuration.
	RuleSetName string `json:"ruleSetName,omitempty" yaml:"ruleSetName,omitempty"`

	// The number of player slots in a match to keep open for future players.
	AdditionalPlayerCount int `json:"additionalPlayerCount,omitempty" yaml:"additionalPlayerCount,omitempty"`

	// The method used to backfill game sessions that are created with this matchmaking configuration.
	BackfillMode string `json:"backfillMode,omitempty" yaml:"backfillMode,omitempty"`
}
