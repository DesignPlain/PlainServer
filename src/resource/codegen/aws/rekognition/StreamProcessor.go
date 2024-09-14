package rekognition

import types "libds/aws/types"

type StreamProcessor struct {
	// Optional parameter for label detection stream processors.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The name of the Stream Processor.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Kinesis data stream stream or Amazon S3 bucket location to which Amazon Rekognition Video puts the analysis results. See `output`.
	Output types.Rekognition_StreamProcessorOutput `json:"output,omitempty" yaml:"output,omitempty"`

	// Specifies locations in the frames where Amazon Rekognition checks for objects or people. See `regions_of_interest`.
	RegionsOfInterests []types.Rekognition_StreamProcessorRegionsOfInterest `json:"regionsOfInterests,omitempty" yaml:"regionsOfInterests,omitempty"`

	/*
	   Input parameters used in a streaming video analyzed by a stream processor. See `settings`.

	   The following arguments are optional:
	*/
	Settings types.Rekognition_StreamProcessorSettings `json:"settings,omitempty" yaml:"settings,omitempty"`

	//
	Timeouts types.Rekognition_StreamProcessorTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// See `data_sharing_preference`.
	DataSharingPreference types.Rekognition_StreamProcessorDataSharingPreference `json:"dataSharingPreference,omitempty" yaml:"dataSharingPreference,omitempty"`

	// Input video stream. See `input`.
	Input types.Rekognition_StreamProcessorInput `json:"input,omitempty" yaml:"input,omitempty"`

	// The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the completion status. See `notification_channel`.
	NotificationChannel types.Rekognition_StreamProcessorNotificationChannel `json:"notificationChannel,omitempty" yaml:"notificationChannel,omitempty"`

	// The Amazon Resource Number (ARN) of the IAM role that allows access to the stream processor. The IAM role provides Rekognition read permissions for a Kinesis stream. It also provides write permissions to an Amazon S3 bucket and Amazon Simple Notification Service topic for a label detection stream processor. This is required for both face search and label detection stream processors.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
