package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type RealtimeLogConfig struct {
	// The unique name to identify this real-time log configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
	SamplingRate int `json:"samplingRate,omitempty" yaml:"samplingRate,omitempty"`

	// The Amazon Kinesis data streams where real-time log data is sent.
	Endpoint types.Cloudfront_RealtimeLogConfigEndpoint `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
	Fields []string `json:"fields,omitempty" yaml:"fields,omitempty"`
}
