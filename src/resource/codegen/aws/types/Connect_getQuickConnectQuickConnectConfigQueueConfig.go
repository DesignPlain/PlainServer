package types

type Connect_getQuickConnectQuickConnectConfigQueueConfig struct {
	// Identifier of the contact flow.
	ContactFlowId string `json:"contactFlowId,omitempty" yaml:"contactFlowId,omitempty"`

	// Identifier for the queue.
	QueueId string `json:"queueId,omitempty" yaml:"queueId,omitempty"`
}
