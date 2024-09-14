package types

type Rekognition_StreamProcessorOutput struct {
	// The Amazon Kinesis Data Streams stream to which the Amazon Rekognition stream processor streams the analysis results. See `kinesis_data_stream`.
	KinesisDataStream Rekognition_StreamProcessorOutputKinesisDataStream `json:"kinesisDataStream,omitempty" yaml:"kinesisDataStream,omitempty"`

	// The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation. See `s3_destination`.
	S3Destination Rekognition_StreamProcessorOutputS3Destination `json:"s3Destination,omitempty" yaml:"s3Destination,omitempty"`
}
