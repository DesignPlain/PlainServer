package types

type Opsworks_RailsAppLayerCloudwatchConfigurationLogStream struct {
	//
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	//
	BufferDuration int `json:"bufferDuration,omitempty" yaml:"bufferDuration,omitempty"`

	//
	FileFingerprintLines string `json:"fileFingerprintLines,omitempty" yaml:"fileFingerprintLines,omitempty"`

	//
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	//
	MultilineStartPattern string `json:"multilineStartPattern,omitempty" yaml:"multilineStartPattern,omitempty"`

	//
	File string `json:"file,omitempty" yaml:"file,omitempty"`

	//
	InitialPosition string `json:"initialPosition,omitempty" yaml:"initialPosition,omitempty"`

	//
	BatchCount int `json:"batchCount,omitempty" yaml:"batchCount,omitempty"`

	//
	BatchSize int `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`

	//
	DatetimeFormat string `json:"datetimeFormat,omitempty" yaml:"datetimeFormat,omitempty"`

	//
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}
