package types

type Elastictranscoder_PipelineNotifications struct {
	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
	Completed string `json:"completed,omitempty" yaml:"completed,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
	Error string `json:"error,omitempty" yaml:"error,omitempty"`

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
	Progressing string `json:"progressing,omitempty" yaml:"progressing,omitempty"`

	/*
	   The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.

	   The `thumbnail_config` object specifies information about the Amazon S3 bucket in
	   which you want Elastic Transcoder to save thumbnail files: which bucket to use,
	   which users you want to have access to the files, the type of access you want
	   users to have, and the storage class that you want to assign to the files. If
	   you specify values for `content_config`, you must also specify values for
	   `thumbnail_config` even if you don't want to create thumbnails. (You control
	   whether to create thumbnails when you create a job. For more information, see
	   ThumbnailPattern in the topic Create Job.) If you specify values for
	   `content_config` and `thumbnail_config`, omit the OutputBucket object.
	*/
	Warning string `json:"warning,omitempty" yaml:"warning,omitempty"`
}
