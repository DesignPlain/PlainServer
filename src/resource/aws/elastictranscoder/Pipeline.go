package elastictranscoder

import types "DesignSphere_Server/src/resource/aws/types"

type Pipeline struct {
	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket string `json:"outputBucket,omitempty" yaml:"outputBucket,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKmsKeyArn string `json:"awsKmsKeyArn,omitempty" yaml:"awsKmsKeyArn,omitempty"`

	// The name of the pipeline. Maximum 40 characters
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications types.Elastictranscoder_PipelineNotifications `json:"notifications,omitempty" yaml:"notifications,omitempty"`

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig types.Elastictranscoder_PipelineThumbnailConfig `json:"thumbnailConfig,omitempty" yaml:"thumbnailConfig,omitempty"`

	/*
	   The permissions for the `thumbnail_config` object. (documented below)

	   The `content_config` object specifies information about the Amazon S3 bucket in
	   which you want Elastic Transcoder to save transcoded files and playlists: which
	   bucket to use, and the storage class that you want to assign to the files. If
	   you specify values for `content_config`, you must also specify values for
	   `thumbnail_config`. If you specify values for `content_config` and
	   `thumbnail_config`, omit the `output_bucket` object.
	*/
	ThumbnailConfigPermissions []types.Elastictranscoder_PipelineThumbnailConfigPermission `json:"thumbnailConfigPermissions,omitempty" yaml:"thumbnailConfigPermissions,omitempty"`

	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig types.Elastictranscoder_PipelineContentConfig `json:"contentConfig,omitempty" yaml:"contentConfig,omitempty"`

	// The permissions for the `content_config` object. (documented below)
	ContentConfigPermissions []types.Elastictranscoder_PipelineContentConfigPermission `json:"contentConfigPermissions,omitempty" yaml:"contentConfigPermissions,omitempty"`

	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket string `json:"inputBucket,omitempty" yaml:"inputBucket,omitempty"`
}
