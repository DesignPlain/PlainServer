package types

type Connect_InstanceStorageConfigStorageConfigKinesisVideoStreamConfig struct {
	// The encryption configuration. Documented below.
	EncryptionConfig Connect_InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	// The prefix of the video stream. Minimum length of `1`. Maximum length of `128`. When read from the state, the value returned is `<prefix>-connect-<connect_instance_alias>-contact-` since the API appends additional details to the `prefix`.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	// The number of hours data is retained in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. Minimum value of `0`. Maximum value of `87600`. A value of `0`, indicates that the stream does not persist data.
	RetentionPeriodHours int `json:"retentionPeriodHours,omitempty" yaml:"retentionPeriodHours,omitempty"`
}
