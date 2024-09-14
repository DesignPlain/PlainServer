package types

type Gamelift_GameSessionQueuePlayerLatencyPolicy struct {
	// Maximum latency value that is allowed for any player.
	MaximumIndividualPlayerLatencyMilliseconds int `json:"maximumIndividualPlayerLatencyMilliseconds,omitempty" yaml:"maximumIndividualPlayerLatencyMilliseconds,omitempty"`

	// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
	PolicyDurationSeconds int `json:"policyDurationSeconds,omitempty" yaml:"policyDurationSeconds,omitempty"`
}
