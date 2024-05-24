package types

type Connect_getInstanceStorageConfigStorageConfig struct {
	// A block that specifies the configuration of the Kinesis Firehose delivery stream. Documented below.
	KinesisFirehoseConfigs []Connect_getInstanceStorageConfigStorageConfigKinesisFirehoseConfig `json:"kinesisFirehoseConfigs,omitempty" yaml:"kinesisFirehoseConfigs,omitempty"`

	// A block that specifies the configuration of the Kinesis data stream. Documented below.
	KinesisStreamConfigs []Connect_getInstanceStorageConfigStorageConfigKinesisStreamConfig `json:"kinesisStreamConfigs,omitempty" yaml:"kinesisStreamConfigs,omitempty"`

	// A block that specifies the configuration of the Kinesis video stream. Documented below.
	KinesisVideoStreamConfigs []Connect_getInstanceStorageConfigStorageConfigKinesisVideoStreamConfig `json:"kinesisVideoStreamConfigs,omitempty" yaml:"kinesisVideoStreamConfigs,omitempty"`

	// A block that specifies the configuration of S3 Bucket. Documented below.
	S3Configs []Connect_getInstanceStorageConfigStorageConfigS3Config `json:"s3Configs,omitempty" yaml:"s3Configs,omitempty"`

	// A valid storage type. Valid Values: `S3` | `KINESIS_VIDEO_STREAM` | `KINESIS_STREAM` | `KINESIS_FIREHOSE`.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`
}
