package types

type Elastictranscoder_PipelineContentConfig struct {
	// The Amazon S3 storage class, `Standard` or `ReducedRedundancy`, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
