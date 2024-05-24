package types

type Cloudwatch_EventTargetKinesisTarget struct {
	// The JSON path to be extracted from the event and used as the partition key.
	PartitionKeyPath string `json:"partitionKeyPath,omitempty" yaml:"partitionKeyPath,omitempty"`
}
