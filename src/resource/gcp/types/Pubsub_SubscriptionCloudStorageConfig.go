package types

type Pubsub_SubscriptionCloudStorageConfig struct {
	/*
	   If set, message data will be written to Cloud Storage in Avro format.
	   Structure is documented below.
	*/
	AvroConfig Pubsub_SubscriptionCloudStorageConfigAvroConfig `json:"avroConfig,omitempty" yaml:"avroConfig,omitempty"`

	// User-provided name for the Cloud Storage bucket. The bucket must be created by the user. The bucket name must be without any prefix like "gs://".
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// User-provided prefix for Cloud Storage filename.
	FilenamePrefix string `json:"filenamePrefix,omitempty" yaml:"filenamePrefix,omitempty"`

	// User-provided suffix for Cloud Storage filename. Must not end in "/".
	FilenameSuffix string `json:"filenameSuffix,omitempty" yaml:"filenameSuffix,omitempty"`

	/*
	   The maximum bytes that can be written to a Cloud Storage file before a new file is created. Min 1 KB, max 10 GiB.
	   The maxBytes limit may be exceeded in cases where messages are larger than the limit.
	*/
	MaxBytes int `json:"maxBytes,omitempty" yaml:"maxBytes,omitempty"`

	/*
	   The maximum duration that can elapse before a new Cloud Storage file is created. Min 1 minute, max 10 minutes, default 5 minutes.
	   May not exceed the subscription's acknowledgement deadline.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	MaxDuration string `json:"maxDuration,omitempty" yaml:"maxDuration,omitempty"`

	/*
	   (Output)
	   An output-only field that indicates whether or not the subscription can receive messages.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
