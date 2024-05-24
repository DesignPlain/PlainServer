package types

type Pipes_PipeTargetParametersSqsQueueParameters struct {
	// This parameter applies only to FIFO (first-in-first-out) queues. The token used for deduplication of sent messages.
	MessageDeduplicationId string `json:"messageDeduplicationId,omitempty" yaml:"messageDeduplicationId,omitempty"`

	// The FIFO message group ID to use as the target.
	MessageGroupId string `json:"messageGroupId,omitempty" yaml:"messageGroupId,omitempty"`
}
