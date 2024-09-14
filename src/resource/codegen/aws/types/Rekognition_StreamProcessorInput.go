package types

type Rekognition_StreamProcessorInput struct {
	// Kinesis input stream. See `kinesis_video_stream`.
	KinesisVideoStream Rekognition_StreamProcessorInputKinesisVideoStream `json:"kinesisVideoStream,omitempty" yaml:"kinesisVideoStream,omitempty"`
}
