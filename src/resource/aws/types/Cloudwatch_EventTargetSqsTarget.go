package types

type Cloudwatch_EventTargetSqsTarget struct {
	// The FIFO message group ID to use as the target.
	MessageGroupId string `json:"messageGroupId,omitempty" yaml:"messageGroupId,omitempty"`
}
