package types

type Datastream_StreamDestinationConfigGcsDestinationConfig struct {
	// AVRO file format configuration.
	AvroFileFormat Datastream_StreamDestinationConfigGcsDestinationConfigAvroFileFormat `json:"avroFileFormat,omitempty" yaml:"avroFileFormat,omitempty"`

	/*
	   The maximum duration for which new events are added before a file is closed and a new file is created.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	*/
	FileRotationInterval string `json:"fileRotationInterval,omitempty" yaml:"fileRotationInterval,omitempty"`

	// The maximum file size to be saved in the bucket.
	FileRotationMb int `json:"fileRotationMb,omitempty" yaml:"fileRotationMb,omitempty"`

	/*
	   JSON file format configuration.
	   Structure is documented below.
	*/
	JsonFileFormat Datastream_StreamDestinationConfigGcsDestinationConfigJsonFileFormat `json:"jsonFileFormat,omitempty" yaml:"jsonFileFormat,omitempty"`

	// Path inside the Cloud Storage bucket to write data to.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
