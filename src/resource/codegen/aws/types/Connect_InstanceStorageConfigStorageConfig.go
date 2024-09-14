package types

type Connect_InstanceStorageConfigStorageConfig struct {
	// A block that specifies the configuration of the Kinesis data stream. Documented below.
	KinesisStreamConfig Connect_InstanceStorageConfigStorageConfigKinesisStreamConfig `json:"kinesisStreamConfig,omitempty" yaml:"kinesisStreamConfig,omitempty"`

	// A block that specifies the configuration of the Kinesis video stream. Documented below.
	KinesisVideoStreamConfig Connect_InstanceStorageConfigStorageConfigKinesisVideoStreamConfig `json:"kinesisVideoStreamConfig,omitempty" yaml:"kinesisVideoStreamConfig,omitempty"`

	// A block that specifies the configuration of S3 Bucket. Documented below.
	S3Config Connect_InstanceStorageConfigStorageConfigS3Config `json:"s3Config,omitempty" yaml:"s3Config,omitempty"`

	// A valid storage type. Valid Values: `S3` | `KINESIS_VIDEO_STREAM` | `KINESIS_STREAM` | `KINESIS_FIREHOSE`.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// A block that specifies the configuration of the Kinesis Firehose delivery stream. Documented below.
	KinesisFirehoseConfig Connect_InstanceStorageConfigStorageConfigKinesisFirehoseConfig `json:"kinesisFirehoseConfig,omitempty" yaml:"kinesisFirehoseConfig,omitempty"`
}
