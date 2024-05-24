package qldb

import types "DesignSphere_Server/src/resource/aws/types"

type Stream struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the [Amazon QLDB Developer Guide](https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming).
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.
	ExclusiveEndTime string `json:"exclusiveEndTime,omitempty" yaml:"exclusiveEndTime,omitempty"`

	// The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: `"2019-06-13T21:36:34Z"`.  This cannot be in the future and must be before `exclusive_end_time`.  If you provide a value that is before the ledger's `CreationDateTime`, QLDB effectively defaults it to the ledger's `CreationDateTime`.
	InclusiveStartTime string `json:"inclusiveStartTime,omitempty" yaml:"inclusiveStartTime,omitempty"`

	// The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
	KinesisConfiguration types.Qldb_StreamKinesisConfiguration `json:"kinesisConfiguration,omitempty" yaml:"kinesisConfiguration,omitempty"`

	// The name of the QLDB ledger.
	LedgerName string `json:"ledgerName,omitempty" yaml:"ledgerName,omitempty"`
}
