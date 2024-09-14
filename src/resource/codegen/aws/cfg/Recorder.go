package cfg

import types "libds/aws/types"

type Recorder struct {
	// The name of the recorder. Defaults to `default`. Changing it recreates the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Recording group - see below.
	RecordingGroup types.Cfg_RecorderRecordingGroup `json:"recordingGroup,omitempty" yaml:"recordingGroup,omitempty"`

	// Recording mode - see below.
	RecordingMode types.Cfg_RecorderRecordingMode `json:"recordingMode,omitempty" yaml:"recordingMode,omitempty"`

	// Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
