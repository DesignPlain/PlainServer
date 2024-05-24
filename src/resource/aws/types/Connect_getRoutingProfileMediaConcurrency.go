package types

type Connect_getRoutingProfileMediaConcurrency struct {
	// Number of contacts an agent can have on a channel simultaneously. Valid Range for `VOICE`: Minimum value of 1. Maximum value of 1. Valid Range for `CHAT`: Minimum value of 1. Maximum value of 10. Valid Range for `TASK`: Minimum value of 1. Maximum value of 10.
	Concurrency int `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`

	// Channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are `VOICE`, `CHAT`, `TASK`.
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`
}
