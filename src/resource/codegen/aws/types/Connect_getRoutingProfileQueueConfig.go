package types

type Connect_getRoutingProfileQueueConfig struct {
	// Channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are `VOICE`, `CHAT`, `TASK`.
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`

	// Delay, in seconds, that a contact should be in the queue before they are routed to an available agent
	Delay int `json:"delay,omitempty" yaml:"delay,omitempty"`

	// Order in which contacts are to be handled for the queue.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// ARN for the queue.
	QueueArn string `json:"queueArn,omitempty" yaml:"queueArn,omitempty"`

	// Identifier for the queue.
	QueueId string `json:"queueId,omitempty" yaml:"queueId,omitempty"`

	// Name for the queue.
	QueueName string `json:"queueName,omitempty" yaml:"queueName,omitempty"`
}
