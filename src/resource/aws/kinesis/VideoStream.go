package kinesis

type VideoStream struct {
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours int `json:"dataRetentionInHours,omitempty" yaml:"dataRetentionInHours,omitempty"`

	// The name of the device that is writing to the stream. --In the current implementation, Kinesis Video Streams does not use this name.--
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType string `json:"mediaType,omitempty" yaml:"mediaType,omitempty"`

	/*
	   A name to identify the stream. This is unique to the
	   AWS account and region the Stream is created in.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
