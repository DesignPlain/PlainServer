package types

type Connect_QuickConnectQuickConnectConfigQueueConfig struct {
	// Specifies the identifier of the contact flow.
	ContactFlowId string `json:"contactFlowId,omitempty" yaml:"contactFlowId,omitempty"`

	// Specifies the identifier for the queue.
	QueueId string `json:"queueId,omitempty" yaml:"queueId,omitempty"`
}
