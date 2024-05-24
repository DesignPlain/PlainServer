package types

type Globalaccelerator_getAcceleratorAttribute struct {
	//
	FlowLogsS3Prefix string `json:"flowLogsS3Prefix,omitempty" yaml:"flowLogsS3Prefix,omitempty"`

	//
	FlowLogsEnabled bool `json:"flowLogsEnabled,omitempty" yaml:"flowLogsEnabled,omitempty"`

	//
	FlowLogsS3Bucket string `json:"flowLogsS3Bucket,omitempty" yaml:"flowLogsS3Bucket,omitempty"`
}
