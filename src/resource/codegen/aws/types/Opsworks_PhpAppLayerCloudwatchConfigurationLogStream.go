package types

type Opsworks_PhpAppLayerCloudwatchConfigurationLogStream struct {
	//
	File string `json:"file,omitempty" yaml:"file,omitempty"`

	//
	FileFingerprintLines string `json:"fileFingerprintLines,omitempty" yaml:"fileFingerprintLines,omitempty"`

	//
	InitialPosition string `json:"initialPosition,omitempty" yaml:"initialPosition,omitempty"`

	//
	MultilineStartPattern string `json:"multilineStartPattern,omitempty" yaml:"multilineStartPattern,omitempty"`

	//
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	//
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	//
	BatchCount int `json:"batchCount,omitempty" yaml:"batchCount,omitempty"`

	//
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	//
	BufferDuration int `json:"bufferDuration,omitempty" yaml:"bufferDuration,omitempty"`

	//
	DatetimeFormat string `json:"datetimeFormat,omitempty" yaml:"datetimeFormat,omitempty"`

	//
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}
