package types

type Ec2_FlowLogDestinationOptions struct {
	// The format for the flow log. Default value: `plain-text`. Valid values: `plain-text`, `parquet`.
	FileFormat string `json:"fileFormat,omitempty" yaml:"fileFormat,omitempty"`

	// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3. Default value: `false`.
	HiveCompatiblePartitions bool `json:"hiveCompatiblePartitions,omitempty" yaml:"hiveCompatiblePartitions,omitempty"`

	// Indicates whether to partition the flow log per hour. This reduces the cost and response time for queries. Default value: `false`.
	PerHourPartition bool `json:"perHourPartition,omitempty" yaml:"perHourPartition,omitempty"`
}
