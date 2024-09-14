package firebase

type StorageBucket struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Required. Immutable. The ID of the underlying Google Cloud Storage bucket
	BucketId string `json:"bucketId,omitempty" yaml:"bucketId,omitempty"`
}
