package types

type Connect_RoutingProfileQueueConfig struct {
	// Specifies the channels agents can handle in the Contact Control Panel (CCP) for this routing profile. Valid values are `VOICE`, `CHAT`, `TASK`.
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`

	// Specifies the delay, in seconds, that a contact should be in the queue before they are routed to an available agent
	Delay int `json:"delay,omitempty" yaml:"delay,omitempty"`

	// Specifies the order in which contacts are to be handled for the queue.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// ARN for the queue.
	QueueArn string `json:"queueArn,omitempty" yaml:"queueArn,omitempty"`

	// Specifies the identifier for the queue.
	QueueId string `json:"queueId,omitempty" yaml:"queueId,omitempty"`

	// Name for the queue.
	QueueName string `json:"queueName,omitempty" yaml:"queueName,omitempty"`
}
