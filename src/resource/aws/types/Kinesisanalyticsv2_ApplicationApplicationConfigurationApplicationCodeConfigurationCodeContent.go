package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContent struct {
	// Information about the Amazon S3 bucket containing the application code.
	S3ContentLocation Kinesisanalyticsv2_ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation `json:"s3ContentLocation,omitempty" yaml:"s3ContentLocation,omitempty"`

	/*
	   The text-format code for the application.

	   The `s3_content_location` object supports the following:
	*/
	TextContent string `json:"textContent,omitempty" yaml:"textContent,omitempty"`
}
