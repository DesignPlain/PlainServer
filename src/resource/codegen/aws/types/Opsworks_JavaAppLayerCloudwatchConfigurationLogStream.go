package types

type Opsworks_JavaAppLayerCloudwatchConfigurationLogStream struct {
	//
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	//
	File string `json:"file,omitempty" yaml:"file,omitempty"`

	//
	FileFingerprintLines string `json:"fileFingerprintLines,omitempty" yaml:"fileFingerprintLines,omitempty"`

	//
	InitialPosition string `json:"initialPosition,omitempty" yaml:"initialPosition,omitempty"`

	//
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	//
	BatchCount int `json:"batchCount,omitempty" yaml:"batchCount,omitempty"`

	//
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	//
	BufferDuration int `json:"bufferDuration,omitempty" yaml:"bufferDuration,omitempty"`

	//
	MultilineStartPattern string `json:"multilineStartPattern,omitempty" yaml:"multilineStartPattern,omitempty"`

	//
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	//
	DatetimeFormat string `json:"datetimeFormat,omitempty" yaml:"datetimeFormat,omitempty"`
}
